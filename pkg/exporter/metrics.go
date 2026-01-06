package exporter

import "github.com/prometheus/client_golang/prometheus"

const (
	torrentHashLabel = "torrent_hash"
	torrentNameLabel = "torrent_name"
)

var (
	downloadedBytes = prometheus.NewDesc(
		"transmission_downloaded_bytes_total",
		"Total number of bytes downloaded since Transmission started (from cumulative-stats).",
		nil, nil,
	)

	uploadedBytes = prometheus.NewDesc(
		"transmission_uploaded_bytes_total",
		"Total number of bytes uploaded since Transmission started (from cumulative-stats).",
		nil, nil,
	)

	torrentsAdded = prometheus.NewDesc(
		"transmission_torrents_added_total",
		"Total number of torrents added since Transmission started (from cumulative-stats).",
		nil, nil,
	)

	secondsActive = prometheus.NewDesc(
		"transmission_seconds_active_total",
		"Total number of seconds Transmission has been active since it started (from cumulative-stats).",
		nil, nil,
	)

	sessions = prometheus.NewDesc(
		"transmission_sessions_total",
		"Total number of sessions since Transmission started (from cumulative-stats).",
		nil, nil,
	)

	uploadSpeed = prometheus.NewDesc(
		"transmission_upload_bytes_per_second",
		"Aggregated upload speed in bytes/s across all torrents.",
		nil, nil,
	)

	downloadSpeed = prometheus.NewDesc(
		"transmission_download_bytes_per_second",
		"Aggregated download speed in bytes/s across all torrents.",
		nil, nil,
	)

	torrents = prometheus.NewDesc(
		"transmission_torrents",
		"Number of torrents by status",
		[]string{"status"}, nil,
	)
)

// torrent-level metrics
var (
	torrentDownloadSpeed = newTorrentLevelDesc(
		"transmission_torrent_download_bytes_per_second",
		"Torrent download speed in bytes/s.",
	)

	torrentUploadSpeed = newTorrentLevelDesc(
		"transmission_torrent_upload_bytes_per_second",
		"Torrent upload speed in bytes/s.",
	)

	torrentTotalSize = newTorrentLevelDesc(
		"transmission_torrent_total_size_bytes",
		"Total size of the torrent in bytes.",
	)

	torrentSizeWhenDone = newTorrentLevelDesc(
		"transmission_torrent_size_when_done_bytes",
		"Size of the torrent when done (bytes); differs from transmission_torrent_total_size_bytes if files are not selected.",
	)

	torrentLeftUntilDone = newTorrentLevelDesc(
		"transmission_torrent_left_until_done_bytes",
		"Bytes left until the torrent is considered done (only counts wanted data).",
	)

	torrentDownloadedEver = newTorrentLevelDesc(
		"transmission_torrent_downloaded_bytes_total",
		"Total bytes downloaded for this torrent since it was added.",
	)

	torrentUploadedEver = newTorrentLevelDesc(
		"transmission_torrent_uploaded_bytes_total",
		"Total bytes uploaded for this torrent since it was added.",
	)

	torrentCorruptedEver = newTorrentLevelDesc(
		"transmission_torrent_corrupt_bytes_total",
		"Total corrupt bytes recorded for this torrent since it was added.",
	)

	torrentPeersConnected = newTorrentLevelDesc(
		"transmission_torrent_peers_connected",
		"Number of peers currently connected for this torrent.",
	)

	torrentPeersSendingToUs = newTorrentLevelDesc(
		"transmission_torrent_peers_sending_to_us",
		"Number of connected peers currently sending data to us for this torrent.",
	)

	torrentPeersGettingFromUs = newTorrentLevelDesc(
		"transmission_torrent_peers_getting_from_us",
		"Number of connected peers currently receiving data from us for this torrent.",
	)

	torrentWebseedsSendingToUs = newTorrentLevelDesc(
		"transmission_torrent_webseeds_sending_to_us",
		"Number of webseeds currently sending data to us for this torrent.",
	)

	torrentSecondsDownloading = newTorrentLevelDesc(
		"transmission_torrent_seconds_downloading",
		"Total seconds this torrent has spent downloading since it was added.",
	)

	torrentSecondsSeeding = newTorrentLevelDesc(
		"transmission_torrent_seconds_seeding",
		"Total seconds this torrent has spent seeding since it was added.",
	)

	torrentInfo = prometheus.NewDesc(
		"transmission_torrent_info",
		"Static information about a Transmission torrent, exposed as labels and intended for joining with other torrent metrics.",
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
