package exporter

import "github.com/prometheus/client_golang/prometheus"

const (
	torrentHashLabel = "torrent_hash"
	torrentNameLabel = "torrent_name"
)

var (
	downloadedBytes = prometheus.NewDesc(
		"transmission_downloaded_bytes_total",
		"Total number of bytes downloaded since Transmission daemon started.",
		nil, nil,
	)

	uploadedBytes = prometheus.NewDesc(
		"transmission_uploaded_bytes_total",
		"Total number of bytes uploaded since Transmission daemon started.",
		nil, nil,
	)

	torrentsAdded = prometheus.NewDesc(
		"transmission_torrents_added_total",
		"Total number of torrents added since Transmission daemon started.",
		nil, nil,
	)

	secondsActive = prometheus.NewDesc(
		"transmission_seconds_active_total",
		"Total number of seconds the Transmission daemon has been active since it started.",
		nil, nil,
	)

	sessions = prometheus.NewDesc(
		"transmission_sessions_total",
		"Total number of sessions since Transmission daemon started.",
		nil, nil,
	)

	uploadSpeed = prometheus.NewDesc(
		"transmission_upload_bytes_per_second",
		"Current aggregated upload speed across all torrents in bytes per second.",
		nil, nil,
	)

	downloadSpeed = prometheus.NewDesc(
		"transmission_download_bytes_per_second",
		"Current aggregated download speed across all torrents in bytes per second.",
		nil, nil,
	)

	torrents = prometheus.NewDesc(
		"transmission_torrents",
		"Number of torrents grouped by status.",
		[]string{"status"}, nil,
	)
)

// torrent-level metrics
var (
	torrentDownloadSpeed = newTorrentLevelDesc(
		"transmission_torrent_download_bytes_per_second",
		"Current download speed for this torrent in bytes per second.",
	)

	torrentUploadSpeed = newTorrentLevelDesc(
		"transmission_torrent_upload_bytes_per_second",
		"Current upload speed for this torrent in bytes per second.",
	)

	torrentTotalSize = newTorrentLevelDesc(
		"transmission_torrent_total_size_bytes",
		"Total size of the torrent in bytes.",
	)

	torrentSizeWhenDone = newTorrentLevelDesc(
		"transmission_torrent_size_when_done_bytes",
		"Size of the torrent when download completes in bytes. May differ from total size if some files are not selected for download.",
	)

	torrentLeftUntilDone = newTorrentLevelDesc(
		"transmission_torrent_left_until_done_bytes",
		"Number of bytes remaining until the torrent download is complete. Only counts wanted data.",
	)

	torrentDownloadedEver = newTorrentLevelDesc(
		"transmission_torrent_downloaded_bytes_total",
		"Total number of bytes downloaded for this torrent since it was added.",
	)

	torrentUploadedEver = newTorrentLevelDesc(
		"transmission_torrent_uploaded_bytes_total",
		"Total number of bytes uploaded for this torrent since it was added.",
	)

	torrentCorruptedEver = newTorrentLevelDesc(
		"transmission_torrent_corrupt_bytes_total",
		"Total number of corrupt bytes recorded for this torrent since it was added.",
	)

	torrentPeersConnected = newTorrentLevelDesc(
		"transmission_torrent_peers_connected",
		"Current number of peers connected for this torrent.",
	)

	torrentPeersSendingToUs = newTorrentLevelDesc(
		"transmission_torrent_peers_sending_to_us",
		"Current number of connected peers sending data to us for this torrent.",
	)

	torrentPeersGettingFromUs = newTorrentLevelDesc(
		"transmission_torrent_peers_getting_from_us",
		"Current number of connected peers receiving data from us for this torrent.",
	)

	torrentWebseedsSendingToUs = newTorrentLevelDesc(
		"transmission_torrent_webseeds_sending_to_us",
		"Current number of webseeds sending data to us for this torrent.",
	)

	torrentSecondsDownloading = newTorrentLevelDesc(
		"transmission_torrent_seconds_downloading_total",
		"Total number of seconds this torrent has spent downloading since it was added.",
	)

	torrentSecondsSeeding = newTorrentLevelDesc(
		"transmission_torrent_seconds_seeding_total",
		"Total number of seconds this torrent has spent seeding since it was added.",
	)

	torrentInfo = prometheus.NewDesc(
		"transmission_torrent_info",
		"Static information about a Transmission torrent. Always has value 1. Use this metric to join with other torrent-level metrics using the torrent_hash and torrent_name labels.",
		[]string{torrentHashLabel, torrentNameLabel}, nil,
	)
)

func newTorrentLevelDesc(fqName string, help string) *prometheus.Desc {
	return prometheus.NewDesc(
		fqName,
		help,
		[]string{torrentHashLabel}, nil,
	)
}
