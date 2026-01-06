package exporter

import (
	"context"
	"log/slog"
	"time"

	"github.com/j-dumbell/go-qbittorrent/pkg/transmission"
	"github.com/prometheus/client_golang/prometheus"
)

// Exporter implements prometheus.Collector
type Exporter struct {
	transmissionClient        TransmissionClient
	logger                    *slog.Logger
	exportTorrentLevelMetrics bool
}

type TransmissionClient interface {
	SessionStats(ctx context.Context) (*transmission.SessionStatsResult, error)
	TorrentGet(ctx context.Context, args transmission.TorrentGetArgs) (*transmission.TorrentGetResult, error)
}

func New(transmissionClient *transmission.Client, logger *slog.Logger, exportTorrentLevelMetrics bool) *Exporter {
	return &Exporter{transmissionClient: transmissionClient, logger: logger, exportTorrentLevelMetrics: exportTorrentLevelMetrics}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- downloadedBytes
	ch <- uploadedBytes
	ch <- torrentsAdded
	ch <- secondsActive
	ch <- sessions
	ch <- uploadSpeed
	ch <- downloadSpeed
	ch <- torrents

	if e.exportTorrentLevelMetrics {
		ch <- torrentDownloadSpeed
		ch <- torrentUploadSpeed
		ch <- torrentTotalSize
		ch <- torrentSizeWhenDone
		ch <- torrentLeftUntilDone
		ch <- torrentDownloadedEver
		ch <- torrentUploadedEver
		ch <- torrentCorruptedEver
		ch <- torrentPeersConnected
		ch <- torrentPeersSendingToUs
		ch <- torrentPeersGettingFromUs
		ch <- torrentWebseedsSendingToUs
		ch <- torrentSecondsDownloading
		ch <- torrentSecondsSeeding
		ch <- torrentInfo
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	statsResult, err := e.transmissionClient.SessionStats(ctx)
	if err != nil {
		e.logger.Error(
			"error getting session stats from Transmission API",
			"err", err,
		)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		downloadedBytes,
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.DownloadedBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		uploadedBytes,
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.UploadedBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		torrentsAdded,
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.FilesAdded),
	)

	ch <- prometheus.MustNewConstMetric(
		secondsActive,
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.SecondsActive),
	)

	ch <- prometheus.MustNewConstMetric(
		sessions,
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.SessionCount),
	)

	ch <- prometheus.MustNewConstMetric(
		uploadSpeed,
		prometheus.GaugeValue,
		float64(statsResult.UploadSpeed),
	)

	ch <- prometheus.MustNewConstMetric(
		downloadSpeed,
		prometheus.GaugeValue,
		float64(statsResult.DownloadSpeed),
	)

	torrentGetResult, err := e.transmissionClient.TorrentGet(ctx, transmission.TorrentGetArgs{
		IDs:    transmission.AllTorrents,
		Fields: transmission.AllTorrentFields,
	})
	if err != nil {
		e.logger.Error("error getting torrents from transmission API", "err", err)
		return
	}

	torrentCountByStatus := newTorrentCountByStatus()

	for _, torrent := range torrentGetResult.Torrents {
		torrentCountByStatus[torrent.Status.String()]++

		if e.exportTorrentLevelMetrics {
			ch <- prometheus.MustNewConstMetric(
				torrentDownloadSpeed,
				prometheus.GaugeValue,
				float64(torrent.RateDownload),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentUploadSpeed,
				prometheus.GaugeValue,
				float64(torrent.RateUpload),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentTotalSize,
				prometheus.GaugeValue,
				float64(torrent.TotalSize),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentSizeWhenDone,
				prometheus.GaugeValue,
				float64(torrent.SizeWhenDone),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLeftUntilDone,
				prometheus.GaugeValue,
				float64(torrent.LeftUntilDone),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentDownloadedEver,
				prometheus.CounterValue,
				float64(torrent.DownloadedEver),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentUploadedEver,
				prometheus.CounterValue,
				float64(torrent.UploadedEver),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentCorruptedEver,
				prometheus.CounterValue,
				float64(torrent.CorruptEver),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentPeersConnected,
				prometheus.GaugeValue,
				float64(torrent.PeersConnected),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentPeersSendingToUs,
				prometheus.GaugeValue,
				float64(torrent.PeersSendingToUs),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentPeersGettingFromUs,
				prometheus.GaugeValue,
				float64(torrent.PeersGettingFromUs),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentWebseedsSendingToUs,
				prometheus.GaugeValue,
				float64(torrent.WebseedsSendingToUs),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentSecondsDownloading,
				prometheus.CounterValue,
				float64(torrent.SecondsDownloading),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentSecondsSeeding,
				prometheus.CounterValue,
				float64(torrent.SecondsSeeding),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentInfo,
				prometheus.GaugeValue,
				1,
				torrent.HashString,
				torrent.Name,
			)
		}
	}

	for status, count := range torrentCountByStatus {
		ch <- prometheus.MustNewConstMetric(torrents, prometheus.GaugeValue, float64(count), status)
	}

}

func newTorrentCountByStatus() map[string]int {
	var m = make(map[string]int)
	for _, label := range transmission.TorrentStatusByLabel {
		m[label] = 0
	}
	m["unknown"] = 0
	return m
}
