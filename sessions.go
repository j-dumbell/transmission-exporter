package transmission

import "context"

type Session struct {
	AltSpeedDown                     int     `json:"alt-speed-down"`
	AltSpeedEnabled                  bool    `json:"alt-speed-enabled"`
	AltSpeedTimeBegin                int     `json:"alt-speed-time-begin"`
	AltSpeedTimeDay                  int     `json:"alt-speed-time-day"`
	AltSpeedTimeEnabled              bool    `json:"alt-speed-time-enabled"`
	AltSpeedTimeEnd                  int     `json:"alt-speed-time-end"`
	AltSpeedUp                       int     `json:"alt-speed-up"`
	AntiBruteForceEnabled            bool    `json:"anti-brute-force-enabled"`
	AntiBruteForceThreshold          int     `json:"anti-brute-force-threshold"`
	BlocklistEnabled                 bool    `json:"blocklist-enabled"`
	BlocklistSize                    int     `json:"blocklist-size"`
	BlocklistURL                     string  `json:"blocklist-url"`
	CacheSizeMB                      int     `json:"cache-size-mb"`
	ConfigDir                        string  `json:"config-dir"`
	DefaultTrackers                  string  `json:"default-trackers"`
	DHTEnabled                       bool    `json:"dht-enabled"`
	DownloadDir                      string  `json:"download-dir"`
	DownloadDirFreeSpace             int64   `json:"download-dir-free-space"`
	DownloadQueueEnabled             bool    `json:"download-queue-enabled"`
	DownloadQueueSize                int     `json:"download-queue-size"`
	Encryption                       string  `json:"encryption"`
	IdleSeedingLimit                 int     `json:"idle-seeding-limit"`
	IdleSeedingLimitEnabled          bool    `json:"idle-seeding-limit-enabled"`
	IncompleteDir                    string  `json:"incomplete-dir"`
	IncompleteDirEnabled             bool    `json:"incomplete-dir-enabled"`
	LPDEnabled                       bool    `json:"lpd-enabled"`
	PeerLimitGlobal                  int     `json:"peer-limit-global"`
	PeerLimitPerTorrent              int     `json:"peer-limit-per-torrent"`
	PeerPort                         int     `json:"peer-port"`
	PeerPortRandomOnStart            bool    `json:"peer-port-random-on-start"`
	PEXEnabled                       bool    `json:"pex-enabled"`
	PortForwardingEnabled            bool    `json:"port-forwarding-enabled"`
	QueueStalledEnabled              bool    `json:"queue-stalled-enabled"`
	QueueStalledMinutes              int     `json:"queue-stalled-minutes"`
	RenamePartialFiles               bool    `json:"rename-partial-files"`
	RPCVersion                       int     `json:"rpc-version"`
	RPCVersionMinimum                int     `json:"rpc-version-minimum"`
	RPCVersionSemver                 string  `json:"rpc-version-semver"`
	ScriptTorrentAddedEnabled        bool    `json:"script-torrent-added-enabled"`
	ScriptTorrentAddedFilename       string  `json:"script-torrent-added-filename"`
	ScriptTorrentDoneEnabled         bool    `json:"script-torrent-done-enabled"`
	ScriptTorrentDoneFilename        string  `json:"script-torrent-done-filename"`
	ScriptTorrentDoneSeedingEnabled  bool    `json:"script-torrent-done-seeding-enabled"`
	ScriptTorrentDoneSeedingFilename string  `json:"script-torrent-done-seeding-filename"`
	SeedQueueEnabled                 bool    `json:"seed-queue-enabled"`
	SeedQueueSize                    int     `json:"seed-queue-size"`
	SeedRatioLimit                   float64 `json:"seedRatioLimit"`
	SeedRatioLimited                 bool    `json:"seedRatioLimited"`
	SessionID                        string  `json:"session-id"`
	SpeedLimitDown                   int     `json:"speed-limit-down"`
	SpeedLimitDownEnabled            bool    `json:"speed-limit-down-enabled"`
	SpeedLimitUp                     int     `json:"speed-limit-up"`
	SpeedLimitUpEnabled              bool    `json:"speed-limit-up-enabled"`
	StartAddedTorrents               bool    `json:"start-added-torrents"`
	TCPEnabled                       bool    `json:"tcp-enabled"`
	TrashOriginalTorrentFiles        bool    `json:"trash-original-torrent-files"`
	Units                            Units   `json:"units"`
	UTPEnabled                       bool    `json:"utp-enabled"`
	Version                          string  `json:"version"`
}

type Units struct {
	MemoryBytes int      `json:"memory-bytes"`
	MemoryUnits []string `json:"memory-units"`
	SizeBytes   int      `json:"size-bytes"`
	SizeUnits   []string `json:"size-units"`
	SpeedBytes  int      `json:"speed-bytes"`
	SpeedUnits  []string `json:"speed-units"`
}

func (c *Client) SessionGet(ctx context.Context) (*Session, error) {
	var session Session
	if err := c.post(ctx, Request{Method: "session-get"}, &session); err != nil {
		return nil, err
	}
	return &session, nil
}
