package exporter

import "github.com/prometheus/client_golang/prometheus"

const (
	hashLabel   = "hash"
	nameLabel   = "name"
	statusLabel = "status"
)

type metricName string

const (
	// global
	metricNameDownloadedBytesTotal   metricName = "transmission_downloaded_bytes_total"
	metricNameUploadedBytesTotal     metricName = "transmission_uploaded_bytes_total"
	metricNameTorrentsAddedTotal     metricName = "transmission_torrents_added_total"
	metricNameSecondsActiveTotal     metricName = "transmission_seconds_active_total"
	metricNameSessionsTotal          metricName = "transmission_sessions_total"
	metricNameUploadBytesPerSecond   metricName = "transmission_upload_bytes_per_second"
	metricNameDownloadBytesPerSecond metricName = "transmission_download_bytes_per_second"
	metricNameTorrents               metricName = "transmission_torrents"

	// torrent-level
	metricNameTorrentDownloadBytesPerSecond  metricName = "transmission_torrent_download_bytes_per_second"
	metricNameTorrentUploadBytesPerSecond    metricName = "transmission_torrent_upload_bytes_per_second"
	metricNameTorrentTotalSizeBytes          metricName = "transmission_torrent_total_size_bytes"
	metricNameTorrentSizeWhenDoneBytes       metricName = "transmission_torrent_size_when_done_bytes"
	metricNameTorrentLeftUntilDoneBytes      metricName = "transmission_torrent_left_until_done_bytes"
	metricNameTorrentDownloadBytesTotal      metricName = "transmission_torrent_downloaded_bytes_total"
	metricNameTorrentUploadBytesTotal        metricName = "transmission_torrent_uploaded_bytes_total"
	metricNameTorrentCorruptBytesTotal       metricName = "transmission_torrent_corrupt_bytes_total"
	metricNameTorrentPeersConnected          metricName = "transmission_torrent_peers_connected"
	metricNameTorrentPeersSendingToUs        metricName = "transmission_torrent_peers_sending_to_us"
	metricNameTorrentPeersGettingFromUs      metricName = "transmission_torrent_peers_getting_from_us"
	metricNameTorrentWebseedsSendingToUs     metricName = "transmission_torrent_webseeds_sending_to_us"
	metricNameTorrentSecondsDownloadingTotal metricName = "transmission_torrent_seconds_downloading_total"
	metricNameTorrentSecondsSeedingTotal     metricName = "transmission_torrent_seconds_seeding_total"
	metricNameTorrentInfo                    metricName = "transmission_torrent_info"
)

type descConfig struct {
	Metric         metricName
	Help           string
	VariableLabels []string
}

var globalDescConfigs = []descConfig{
	{
		Metric: metricNameDownloadedBytesTotal,
		Help:   "Total number of bytes downloaded since Transmission daemon started.",
	},
	{
		Metric: metricNameUploadedBytesTotal,
		Help:   "Total number of bytes uploaded since Transmission daemon started.",
	},
	{
		Metric: metricNameTorrentsAddedTotal,
		Help:   "Total number of torrents added since Transmission daemon started.",
	},
	{
		Metric: metricNameSecondsActiveTotal,
		Help:   "Total number of seconds the Transmission daemon has been active since it started.",
	},
	{
		Metric: metricNameSessionsTotal,
		Help:   "Total number of sessions since Transmission daemon started.",
	},
	{
		Metric: metricNameUploadBytesPerSecond,
		Help:   "Current aggregated upload speed across all torrents in bytes per second.",
	},
	{
		Metric: metricNameDownloadBytesPerSecond,
		Help:   "Current aggregated download speed across all torrents in bytes per second.",
	},
	{
		Metric:         metricNameTorrents,
		Help:           "Number of torrents grouped by status.",
		VariableLabels: []string{statusLabel},
	},
}

var globalDescs = descByMetricName(globalDescConfigs)

var torrentLevelDescConfigs = []descConfig{
	{
		Metric:         metricNameTorrentDownloadBytesPerSecond,
		Help:           "Current download speed for this torrent in bytes per second.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentUploadBytesPerSecond,
		Help:           "Current upload speed for this torrent in bytes per second.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentTotalSizeBytes,
		Help:           "Total size of the torrent in bytes.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentSizeWhenDoneBytes,
		Help:           "Size of the torrent when download completes in bytes. May differ from total size if some files are not selected for download.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentLeftUntilDoneBytes,
		Help:           "Number of bytes remaining until the torrent download is complete. Only counts wanted data.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentDownloadBytesTotal,
		Help:           "Total number of bytes downloaded for this torrent since it was added.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentUploadBytesTotal,
		Help:           "Total number of bytes uploaded for this torrent since it was added.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentCorruptBytesTotal,
		Help:           "Total number of corrupt bytes recorded for this torrent since it was added.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentPeersConnected,
		Help:           "Current number of peers connected for this torrent.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentPeersSendingToUs,
		Help:           "Current number of connected peers sending data to us for this torrent.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentPeersGettingFromUs,
		Help:           "Current number of connected peers receiving data from us for this torrent.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentWebseedsSendingToUs,
		Help:           "Current number of webseeds sending data to us for this torrent.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentSecondsDownloadingTotal,
		Help:           "Total number of seconds this torrent has spent downloading since it was added.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentSecondsSeedingTotal,
		Help:           "Total number of seconds this torrent has spent seeding since it was added.",
		VariableLabels: []string{hashLabel},
	},
	{
		Metric:         metricNameTorrentInfo,
		Help:           "Static information about a Transmission torrent. Always has value 1. Use this metric to join with other torrent-level metrics using the torrent_hash and torrent_name labels.",
		VariableLabels: []string{hashLabel, nameLabel},
	},
}

var torrentLevelDescs = descByMetricName(torrentLevelDescConfigs)

func descByMetricName(descConfigs []descConfig) map[metricName]*prometheus.Desc {
	var result = make(map[metricName]*prometheus.Desc)
	for _, config := range descConfigs {
		result[config.Metric] = prometheus.NewDesc(
			string(config.Metric),
			config.Help,
			config.VariableLabels,
			nil,
		)
	}
	return result
}
