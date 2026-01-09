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
	SessionGet(ctx context.Context) (*transmission.Session, error)
	TorrentGet(ctx context.Context, args transmission.TorrentGetArgs) (*transmission.TorrentGetResult, error)
}

func New(transmissionClient TransmissionClient, logger *slog.Logger, exportTorrentLevelMetrics bool) *Exporter {
	return &Exporter{transmissionClient: transmissionClient, logger: logger, exportTorrentLevelMetrics: exportTorrentLevelMetrics}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, desc := range globalDescs {
		ch <- desc
	}

	if e.exportTorrentLevelMetrics {
		for _, desc := range torrentLevelDescs {
			ch <- desc
		}
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
		globalDescs[metricNameDownloadedBytesTotal],
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.DownloadedBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		globalDescs[metricNameUploadedBytesTotal],
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.UploadedBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		globalDescs[metricNameTorrentsAddedTotal],
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.FilesAdded),
	)

	ch <- prometheus.MustNewConstMetric(
		globalDescs[metricNameSecondsActiveTotal],
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.SecondsActive),
	)

	ch <- prometheus.MustNewConstMetric(
		globalDescs[metricNameSessionsTotal],
		prometheus.CounterValue,
		float64(statsResult.CumulativeStats.SessionCount),
	)

	ch <- prometheus.MustNewConstMetric(
		globalDescs[metricNameUploadBytesPerSecond],
		prometheus.GaugeValue,
		float64(statsResult.UploadSpeed),
	)

	ch <- prometheus.MustNewConstMetric(
		globalDescs[metricNameDownloadBytesPerSecond],
		prometheus.GaugeValue,
		float64(statsResult.DownloadSpeed),
	)

	session, err := e.transmissionClient.SessionGet(ctx)
	if err != nil {
		e.logger.Error("error getting session from Transmission API", "err", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		globalDescs[metricNameVersion],
		prometheus.GaugeValue,
		1,
		session.Version.Sem(),
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
				torrentLevelDescs[metricNameTorrentDownloadBytesPerSecond],
				prometheus.GaugeValue,
				float64(torrent.RateDownload),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentUploadBytesPerSecond],
				prometheus.GaugeValue,
				float64(torrent.RateUpload),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentTotalSizeBytes],
				prometheus.GaugeValue,
				float64(torrent.TotalSize),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentSizeWhenDoneBytes],
				prometheus.GaugeValue,
				float64(torrent.SizeWhenDone),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentLeftUntilDoneBytes],
				prometheus.GaugeValue,
				float64(torrent.LeftUntilDone),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentDownloadBytesTotal],
				prometheus.CounterValue,
				float64(torrent.DownloadedEver),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentUploadBytesTotal],
				prometheus.CounterValue,
				float64(torrent.UploadedEver),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentCorruptBytesTotal],
				prometheus.CounterValue,
				float64(torrent.CorruptEver),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentPeersConnected],
				prometheus.GaugeValue,
				float64(torrent.PeersConnected),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentPeersSendingToUs],
				prometheus.GaugeValue,
				float64(torrent.PeersSendingToUs),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentPeersGettingFromUs],
				prometheus.GaugeValue,
				float64(torrent.PeersGettingFromUs),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentWebseedsSendingToUs],
				prometheus.GaugeValue,
				float64(torrent.WebseedsSendingToUs),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentSecondsDownloadingTotal],
				prometheus.CounterValue,
				float64(torrent.SecondsDownloading),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentSecondsSeedingTotal],
				prometheus.CounterValue,
				float64(torrent.SecondsSeeding),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentStatus],
				prometheus.GaugeValue,
				float64(torrent.Status),
				torrent.HashString,
			)

			ch <- prometheus.MustNewConstMetric(
				torrentLevelDescs[metricNameTorrentInfo],
				prometheus.GaugeValue,
				1,
				torrent.HashString,
				torrent.Name,
			)
		}
	}

	for status, count := range torrentCountByStatus {
		ch <- prometheus.MustNewConstMetric(globalDescs[metricNameTorrents], prometheus.GaugeValue, float64(count), status)
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
