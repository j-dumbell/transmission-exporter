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
	return post[Session](ctx, c, "session-get")
}

type SessionSetArgs struct {
	AltSpeedDown                     *int     `json:"alt-speed-down,omitempty"`
	AltSpeedEnabled                  *bool    `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeBegin                *int     `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeDay                  *int     `json:"alt-speed-time-day,omitempty"`
	AltSpeedTimeEnabled              *bool    `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd                  *int     `json:"alt-speed-time-end,omitempty"`
	AltSpeedUp                       *int     `json:"alt-speed-up,omitempty"`
	AntiBruteForceEnabled            *bool    `json:"anti-brute-force-enabled,omitempty"`
	AntiBruteForceThreshold          *int     `json:"anti-brute-force-threshold,omitempty"`
	BlocklistEnabled                 *bool    `json:"blocklist-enabled,omitempty"`
	BlocklistURL                     *string  `json:"blocklist-url,omitempty"`
	CacheSizeMB                      *int     `json:"cache-size-mb,omitempty"`
	DefaultTrackers                  *string  `json:"default-trackers,omitempty"`
	DHTEnabled                       *bool    `json:"dht-enabled,omitempty"`
	DownloadDir                      *string  `json:"download-dir,omitempty"`
	DownloadQueueEnabled             *bool    `json:"download-queue-enabled,omitempty"`
	DownloadQueueSize                *int     `json:"download-queue-size,omitempty"`
	Encryption                       *string  `json:"encryption,omitempty"`
	IdleSeedingLimit                 *int     `json:"idle-seeding-limit,omitempty"`
	IdleSeedingLimitEnabled          *bool    `json:"idle-seeding-limit-enabled,omitempty"`
	IncompleteDir                    *string  `json:"incomplete-dir,omitempty"`
	IncompleteDirEnabled             *bool    `json:"incomplete-dir-enabled,omitempty"`
	LPDEnabled                       *bool    `json:"lpd-enabled,omitempty"`
	PeerLimitGlobal                  *int     `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent              *int     `json:"peer-limit-per-torrent,omitempty"`
	PeerPort                         *int     `json:"peer-port,omitempty"`
	PeerPortRandomOnStart            *bool    `json:"peer-port-random-on-start,omitempty"`
	PEXEnabled                       *bool    `json:"pex-enabled,omitempty"`
	PortForwardingEnabled            *bool    `json:"port-forwarding-enabled,omitempty"`
	QueueStalledEnabled              *bool    `json:"queue-stalled-enabled,omitempty"`
	QueueStalledMinutes              *int     `json:"queue-stalled-minutes,omitempty"`
	RenamePartialFiles               *bool    `json:"rename-partial-files,omitempty"`
	ScriptTorrentAddedEnabled        *bool    `json:"script-torrent-added-enabled,omitempty"`
	ScriptTorrentAddedFilename       *string  `json:"script-torrent-added-filename,omitempty"`
	ScriptTorrentDoneEnabled         *bool    `json:"script-torrent-done-enabled,omitempty"`
	ScriptTorrentDoneFilename        *string  `json:"script-torrent-done-filename,omitempty"`
	ScriptTorrentDoneSeedingEnabled  *bool    `json:"script-torrent-done-seeding-enabled,omitempty"`
	ScriptTorrentDoneSeedingFilename *string  `json:"script-torrent-done-seeding-filename,omitempty"`
	SeedQueueEnabled                 *bool    `json:"seed-queue-enabled,omitempty"`
	SeedQueueSize                    *int     `json:"seed-queue-size,omitempty"`
	SeedRatioLimit                   *float64 `json:"seedRatioLimit,omitempty"`
	SeedRatioLimited                 *bool    `json:"seedRatioLimited,omitempty"`
	SpeedLimitDown                   *int     `json:"speed-limit-down,omitempty"`
	SpeedLimitDownEnabled            *bool    `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUp                     *int     `json:"speed-limit-up,omitempty"`
	SpeedLimitUpEnabled              *bool    `json:"speed-limit-up-enabled,omitempty"`
	StartAddedTorrents               *bool    `json:"start-added-torrents,omitempty"`
	TrashOriginalTorrentFiles        *bool    `json:"trash-original-torrent-files,omitempty"`
	UTPEnabled                       *bool    `json:"utp-enabled,omitempty"`
}

func (c *Client) SessionSet(ctx context.Context, args SessionSetArgs) error {
	_, err := postWithArgs[SessionSetArgs, any](ctx, c, "session-set", args)
	return err
}

type SessionStatsResult struct {
	ActiveTorrentCount int     `json:"activeTorrentCount"`
	DownloadSpeed      float64 `json:"downloadSpeed"`
	PausedTorrentCount int     `json:"pausedTorrentCount"`
	TorrentCount       int     `json:"torrentCount"`
	UploadSpeed        int     `json:"uploadSpeed"`
	CumulativeStats    Stats   `json:"cumulative-stats"`
	CurrentStats       Stats   `json:"current-stats"`
}

type Stats struct {
	UploadedBytes   int `json:"uploadedBytes"`
	DownloadedBytes int `json:"downloadedBytes"`
	FilesAdded      int `json:"filesAdded"`
	SecondsActive   int `json:"secondsActive"`
	SessionCount    int `json:"sessionCount"`
}

func (c *Client) SessionStats(ctx context.Context) (*SessionStatsResult, error) {
	return post[SessionStatsResult](ctx, c, "session-stats")
}

func (c *Client) BlocklistUpdate(ctx context.Context) error {
	_, err := post[any](ctx, c, "blocklist-update")
	return err
}

type PortTestResult struct {
	PortIsOpen bool   `json:"port_is_open"`
	IPProtocol string `json:"ip_protocol,omitempty"`
}

func (c *Client) PortTest(ctx context.Context) (*PortTestResult, error) {
	return post[PortTestResult](ctx, c, "port-test")
}

func (c *Client) SessionClose(ctx context.Context) error {
	_, err := post[any](ctx, c, "session-close")
	return err
}

type FreeSpaceResult struct {
	Path      string `json:"path"`
	SizeBytes int    `json:"size-bytes"`
	TotalSize int    `json:"total_size"`
}

type FreeSpaceArgs struct {
	Path string `json:"path"`
}

func (c *Client) FreeSpace(ctx context.Context, args FreeSpaceArgs) (*FreeSpaceResult, error) {
	return postWithArgs[FreeSpaceArgs, FreeSpaceResult](ctx, c, "free-space", args)
}

type GroupSetArgs struct {
	Name                  string `json:"name"`
	HonorsSessionLimits   *bool  `json:"honors_session_limits,omitempty"`
	SpeedLimitDown        *int   `json:"speed_limit_down"`
	SpeedLimitDownEnabled *bool  `json:"speed_limit_down_enabled"`
	SpeedLimitUp          *int   `json:"speed_limit_up"`
	SpeedLimitUpEnabled   *bool  `json:"speed_limit_up_enabled"`
}

func (c *Client) GroupSet(ctx context.Context, args GroupSetArgs) error {
	_, err := postWithArgs[GroupSetArgs, any](ctx, c, "group-set", args)
	return err
}

type GroupGetArgs struct {
	Group []string
}

type groupGetArgs struct {
	Group []string `json:"group,omitempty"`
}

type GroupGetResult struct {
	Group []Group `json:"group"`
}

type Group struct {
	Name                  string `json:"name"`
	HonorsSessionLimits   bool   `json:"honorsSessionLimits"`
	SpeedLimitDown        int64  `json:"speed-limit-down"`
	SpeedLimitDownEnabled bool   `json:"speed-limit-down-enabled"`
	SpeedLimitUp          int64  `json:"speed-limit-up"`
	SpeedLimitUpEnabled   bool   `json:"speed-limit-up-enabled"`
}

// pass nil to get all groups
func (c *Client) GroupGet(ctx context.Context, args *GroupGetArgs) (*GroupGetResult, error) {
	gArgs := groupGetArgs{Group: nil}
	if args != nil {
		gArgs = groupGetArgs{Group: args.Group}
	}
	return postWithArgs[groupGetArgs, GroupGetResult](ctx, c, "group-get", gArgs)
}
