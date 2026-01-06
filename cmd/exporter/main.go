package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/j-dumbell/go-qbittorrent/pkg/exporter"
	"github.com/j-dumbell/go-qbittorrent/pkg/transmission"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: getLogLevel(),
	})
	logger := slog.New(logHandler)

	if err := run(logger); err != nil {
		logger.Error("fatal error", "err", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	host, errHost := getEnv("TRANSMISSION_HOST")
	user, errUser := getEnv("TRANSMISSION_USER")
	password, errPass := getEnv("TRANSMISSION_PASSWORD")
	port, errPort := getEnv("PORT")
	if err := errors.Join(errHost, errUser, errPass, errPort); err != nil {
		return err
	}

	exportTorrentLevelMetrics := false
	exportTorrentLevelMetricsEnvVarValue, ok := os.LookupEnv("EXPORT_TORRENT_LEVEL_METRICS")
	if ok && exportTorrentLevelMetricsEnvVarValue == "true" {
		exportTorrentLevelMetrics = true
		logger.Info("EXPORT_TORRENT_LEVEL_METRICS set to true, so will export torrent-level metrics")
	}

	client, err := transmission.New(transmission.ClientParams{
		Host:     host,
		User:     user,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("error instantiating transmission client: %w", err)
	}

	transmissionExporter := exporter.New(client, logger, exportTorrentLevelMetrics)

	reg := prometheus.NewRegistry()
	if err := reg.Register(transmissionExporter); err != nil {
		return fmt.Errorf("error registering collector: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	metricsServer := http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(stop)

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("starting metrics server", "port", port)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
			return
		}
		serverErrors <- nil
	}()

	select {
	case stopSignal := <-stop:
		timeout := 10 * time.Second
		logger.Info(
			"signal received, shutting down server",
			"signal", stopSignal,
			"timeoutSeconds", timeout.Seconds(),
		)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if errShutdown := metricsServer.Shutdown(ctx); errShutdown != nil {
			logger.Error("server failed to shutdown gracefully, closing instead", "err", errShutdown)
			errClose := metricsServer.Close()
			return errors.Join(errShutdown, errClose)
		}
		if err := <-serverErrors; err != nil {
			return err
		}
		logger.Info("server shutdown gracefully")

	case err := <-serverErrors:
		return err
	}

	return nil
}

func getEnv(envName string) (string, error) {
	envValue, ok := os.LookupEnv(envName)
	if !ok {
		return "", fmt.Errorf("missing required environment variable '%s'", envName)
	}
	return envValue, nil
}

func getLogLevel() slog.Level {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
