package transmission

import (
	"context"
	"encoding/json"
)

// TorrentIDs identifies torrents in many methods.
// It should not be instantiated directly, instead use one of:
//   - NewTorrentIDs
//   - AllTorrents
//   - RecentlyActiveTorrents
type TorrentIDs struct {
	ids any
}

func (t TorrentIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ids)
}

// NewTorrentIDs accepts string torrent hashes or integer torrent IDs.
func NewTorrentIDs(ids ...any) *TorrentIDs {
	return &TorrentIDs{ids: ids}
}

var (
	// AllTorrents returns all torrents.
	AllTorrents *TorrentIDs = nil

	// RecentlyActiveTorrents returns all recently active torrents.
	RecentlyActiveTorrents = &TorrentIDs{ids: "recently-active"}
)

type torrentIDsArgs struct {
	IDs *TorrentIDs `json:"ids,omitempty"`
}

func (c *Client) TorrentStart(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "torrent-start", torrentIDsArgs{ids})
	return err
}

func (c *Client) TorrentStartNow(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "torrent-start-now", torrentIDsArgs{ids})
	return err
}

func (c *Client) TorrentStop(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "torrent-stop", torrentIDsArgs{ids})
	return err
}

func (c *Client) TorrentVerify(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "torrent-verify", torrentIDsArgs{ids})
	return err
}

func (c *Client) TorrentReannounce(ctx context.Context, ids *TorrentIDs) error {
	_, err := postWithArgs[torrentIDsArgs, any](ctx, c, "torrent-reannounce", torrentIDsArgs{ids})
	return err
}

// TorrentSetArgs represents the "arguments" object for the "torrent-set" RPC
// method in the Transmission API.
type TorrentSetArgs struct {
	// General per-torrent limits / properties
	BandwidthPriority   *int  `json:"bandwidthPriority,omitempty"`   // tr_priority_t
	DownloadLimit       *int  `json:"downloadLimit,omitempty"`       // KB/s
	DownloadLimited     *bool `json:"downloadLimited,omitempty"`     // honor DownloadLimit
	HonorsSessionLimits *bool `json:"honorsSessionLimits,omitempty"` // honor session limits
	PeerLimit           *int  `json:"peer-limit,omitempty"`          // max peers
	QueuePosition       *int  `json:"queuePosition,omitempty"`       // 0..n-1

	// Per-torrent seeding rules
	SeedIdleLimit  *int     `json:"seedIdleLimit,omitempty"`  // minutes
	SeedIdleMode   *int     `json:"seedIdleMode,omitempty"`   // see tr_idlelimit
	SeedRatioLimit *float64 `json:"seedRatioLimit,omitempty"` // ratio
	SeedRatioMode  *int     `json:"seedRatioMode,omitempty"`  // see tr_ratiolimit

	// File selection and priorities
	FilesWanted    []int `json:"files-wanted,omitempty"`    // file indices
	FilesUnwanted  []int `json:"files-unwanted,omitempty"`  // file indices
	PriorityHigh   []int `json:"priority-high,omitempty"`   // file indices
	PriorityLow    []int `json:"priority-low,omitempty"`    // file indices
	PriorityNormal []int `json:"priority-normal,omitempty"` // file indices

	// Location
	Location *string `json:"location,omitempty"` // new content location

	// Trackers
	TrackerAdd     []string        `json:"trackerAdd,omitempty"`     // announce URLs
	TrackerRemove  []int           `json:"trackerRemove,omitempty"`  // tracker IDs
	TrackerReplace [][]interface{} `json:"trackerReplace,omitempty"` // [ [id, url], ... ]

	// Target torrents (ids follows the usual Transmission rules:
	// single id, list of ids/hashStrings, or "recently-active")
	// To support mixed ints/strings, we use []interface{}.
	Ids []interface{} `json:"ids,omitempty"`

	// Per-torrent speed limits
	UploadLimit   *int  `json:"uploadLimit,omitempty"`   // KB/s
	UploadLimited *bool `json:"uploadLimited,omitempty"` // honor UploadLimit
}

func (c *Client) TorrentSet(ctx context.Context, args TorrentSetArgs) error {
	_, err := postWithArgs[TorrentSetArgs, any](ctx, c, "torrent-set", args)
	return err
}

type torrentGetArgs struct {
	IDs    *TorrentIDs `json:"ids,omitempty"`
	Fields []string    `json:"fields"`
	Format string      `json:"format"`
}

type TorrentGetResult struct {
	Torrents []Torrent `json:"torrents"`

	// Removed only present only when ids == "recently_active"
	Removed []int64 `json:"removed,omitempty"`
}

type Torrent struct {
	ActivityDate                int64             `json:"activity_date,omitempty"`
	AddedDate                   int64             `json:"added_date,omitempty"`
	Availability                []int64           `json:"availability,omitempty"`
	BandwidthPriority           int64             `json:"bandwidth_priority,omitempty"`
	BytesCompleted              []int64           `json:"bytes_completed,omitempty"`
	Comment                     string            `json:"comment,omitempty"`
	CorruptEver                 int64             `json:"corrupt_ever,omitempty"`
	Creator                     string            `json:"creator,omitempty"`
	DateCreated                 int64             `json:"date_created,omitempty"`
	DesiredAvailable            int64             `json:"desired_available,omitempty"`
	DoneDate                    int64             `json:"done_date,omitempty"`
	DownloadDir                 string            `json:"download_dir,omitempty"`
	DownloadedEver              int64             `json:"downloaded_ever,omitempty"`
	DownloadLimit               int64             `json:"download_limit,omitempty"`
	DownloadLimited             bool              `json:"download_limited,omitempty"`
	EditDate                    int64             `json:"edit_date,omitempty"`
	Error                       int64             `json:"error,omitempty"`
	ErrorString                 string            `json:"error_string,omitempty"`
	ETA                         int64             `json:"eta,omitempty"`
	ETAIdle                     int64             `json:"eta_idle,omitempty"`
	FileCount                   int64             `json:"file_count,omitempty"`
	Files                       []TorrentFile     `json:"files,omitempty"`
	FileStats                   []TorrentFileStat `json:"file_stats,omitempty"`
	Group                       string            `json:"group,omitempty"`
	HashString                  string            `json:"hash_string,omitempty"`
	HaveUnchecked               int64             `json:"have_unchecked,omitempty"`
	HaveValid                   int64             `json:"have_valid,omitempty"`
	HonorsSessionLimits         bool              `json:"honors_session_limits,omitempty"`
	ID                          int64             `json:"id,omitempty"`
	IsFinished                  bool              `json:"is_finished,omitempty"`
	IsPrivate                   bool              `json:"is_private,omitempty"`
	IsStalled                   bool              `json:"is_stalled,omitempty"`
	Labels                      []string          `json:"labels,omitempty"`
	LeftUntilDone               int64             `json:"left_until_done,omitempty"`
	MagnetLink                  string            `json:"magnet_link,omitempty"`
	ManualAnnounceTime          int64             `json:"manual_announce_time,omitempty"` // deprecated
	MaxConnectedPeers           int64             `json:"max_connected_peers,omitempty"`
	MetadataPercentComplete     float64           `json:"metadata_percent_complete,omitempty"`
	Name                        string            `json:"name,omitempty"`
	PeerLimit                   int64             `json:"peer_limit,omitempty"`
	Peers                       []Peer            `json:"peers,omitempty"`
	PeersConnected              int64             `json:"peers_connected,omitempty"`
	PeersFrom                   *PeersFrom        `json:"peers_from,omitempty"`
	PeersGettingFromUs          int64             `json:"peers_getting_from_us,omitempty"`
	PeersSendingToUs            int64             `json:"peers_sending_to_us,omitempty"`
	PercentComplete             float64           `json:"percent_complete,omitempty"`
	PercentDone                 float64           `json:"percent_done,omitempty"`
	Pieces                      string            `json:"pieces,omitempty"`
	PieceCount                  int64             `json:"piece_count,omitempty"`
	PieceSize                   int64             `json:"piece_size,omitempty"`
	Priorities                  []int64           `json:"priorities,omitempty"`
	PrimaryMIMEType             string            `json:"primary_mime_type,omitempty"`
	QueuePosition               int64             `json:"queue_position,omitempty"`
	RateDownload                int64             `json:"rate_download,omitempty"` // B/s
	RateUpload                  int64             `json:"rate_upload,omitempty"`   // B/s
	RecheckProgress             float64           `json:"recheck_progress,omitempty"`
	SecondsDownloading          int64             `json:"seconds_downloading,omitempty"`
	SecondsSeeding              int64             `json:"seconds_seeding,omitempty"`
	SeedIdleLimit               int64             `json:"seed_idle_limit,omitempty"`
	SeedIdleMode                int64             `json:"seed_idle_mode,omitempty"`
	SeedRatioLimit              float64           `json:"seed_ratio_limit,omitempty"`
	SeedRatioMode               int64             `json:"seed_ratio_mode,omitempty"`
	SequentialDownload          bool              `json:"sequential_download,omitempty"`
	SequentialDownloadFromPiece int64             `json:"sequential_download_from_piece,omitempty"`
	SizeWhenDone                int64             `json:"size_when_done,omitempty"`
	StartDate                   int64             `json:"start_date,omitempty"`
	Status                      int64             `json:"status,omitempty"` // 0..6
	TorrentFile                 string            `json:"torrent_file,omitempty"`
	TotalSize                   int64             `json:"total_size,omitempty"`
	Trackers                    []Tracker         `json:"trackers,omitempty"`
	TrackerList                 string            `json:"tracker_list,omitempty"`
	TrackerStats                []TrackerStat     `json:"tracker_stats,omitempty"`
	UploadedEver                int64             `json:"uploaded_ever,omitempty"`
	UploadLimit                 int64             `json:"upload_limit,omitempty"`
	UploadLimited               bool              `json:"upload_limited,omitempty"`
	UploadRatio                 float64           `json:"upload_ratio,omitempty"`
	Wanted                      []int64           `json:"wanted,omitempty"` // 0/1 values in 4.x
	Webseeds                    []string          `json:"webseeds,omitempty"`
	WebseedsSendingToUs         int64             `json:"webseeds_sending_to_us,omitempty"`
}

type TorrentFile struct {
	BytesCompleted int64  `json:"bytes_completed,omitempty"`
	Length         int64  `json:"length,omitempty"`
	Name           string `json:"name,omitempty"`
	BeginPiece     int64  `json:"begin_piece,omitempty"`
	EndPiece       int64  `json:"end_piece,omitempty"`
}

type TorrentFileStat struct {
	BytesCompleted int64 `json:"bytes_completed,omitempty"`
	Wanted         bool  `json:"wanted,omitempty"`   // NOTE: different from Torrent.Wanted (0/1 array)
	Priority       int64 `json:"priority,omitempty"` // tr_priority_t
}

type Peer struct {
	Address            string  `json:"address,omitempty"`
	BytesToClient      int64   `json:"bytes_to_client,omitempty"`
	BytesToPeer        int64   `json:"bytes_to_peer,omitempty"`
	ClientIsChoked     bool    `json:"client_is_choked,omitempty"`
	ClientIsInterested bool    `json:"client_is_interested,omitempty"`
	ClientName         string  `json:"client_name,omitempty"`
	FlagStr            string  `json:"flag_str,omitempty"`
	IsDownloadingFrom  bool    `json:"is_downloading_from,omitempty"`
	IsEncrypted        bool    `json:"is_encrypted,omitempty"`
	IsIncoming         bool    `json:"is_incoming,omitempty"`
	IsUploadingTo      bool    `json:"is_uploading_to,omitempty"`
	IsUTP              bool    `json:"is_utp,omitempty"`
	PeerID             string  `json:"peer_id,omitempty"`
	PeerIsChoked       bool    `json:"peer_is_choked,omitempty"`
	PeerIsInterested   bool    `json:"peer_is_interested,omitempty"`
	Port               int64   `json:"port,omitempty"`
	Progress           float64 `json:"progress,omitempty"`
	RateToClient       int64   `json:"rate_to_client,omitempty"`
	RateToPeer         int64   `json:"rate_to_peer,omitempty"`
}

type PeersFrom struct {
	FromCache    int64 `json:"from_cache,omitempty"`
	FromDHT      int64 `json:"from_dht,omitempty"`
	FromIncoming int64 `json:"from_incoming,omitempty"`
	FromLPD      int64 `json:"from_lpd,omitempty"`
	FromLTEP     int64 `json:"from_ltep,omitempty"`
	FromPEX      int64 `json:"from_pex,omitempty"`
	FromTracker  int64 `json:"from_tracker,omitempty"`
}

type Tracker struct {
	Announce string `json:"announce,omitempty"`
	ID       int64  `json:"id,omitempty"`
	Scrape   string `json:"scrape,omitempty"`
	Sitename string `json:"sitename,omitempty"`
	Tier     int64  `json:"tier,omitempty"`
}

type TrackerStat struct {
	Announce              string `json:"announce,omitempty"`
	AnnounceState         int64  `json:"announce_state,omitempty"`
	DownloadCount         int64  `json:"download_count,omitempty"`
	DownloaderCount       int64  `json:"downloader_count,omitempty"`
	HasAnnounced          bool   `json:"has_announced,omitempty"`
	HasScraped            bool   `json:"has_scraped,omitempty"`
	Host                  string `json:"host,omitempty"`
	ID                    int64  `json:"id,omitempty"`
	IsBackup              bool   `json:"is_backup,omitempty"`
	LastAnnouncePeerCount int64  `json:"last_announce_peer_count,omitempty"`
	LastAnnounceResult    string `json:"last_announce_result,omitempty"`
	LastAnnounceStartTime int64  `json:"last_announce_start_time,omitempty"`
	LastAnnounceSucceeded bool   `json:"last_announce_succeeded,omitempty"`
	LastAnnounceTime      int64  `json:"last_announce_time,omitempty"`
	LastAnnounceTimedOut  bool   `json:"last_announce_timed_out,omitempty"`
	LastScrapeResult      string `json:"last_scrape_result,omitempty"`
	LastScrapeStartTime   int64  `json:"last_scrape_start_time,omitempty"`
	LastScrapeSucceeded   bool   `json:"last_scrape_succeeded,omitempty"`
	LastScrapeTime        int64  `json:"last_scrape_time,omitempty"`
	LastScrapeTimedOut    bool   `json:"last_scrape_timed_out,omitempty"`
	LeecherCount          int64  `json:"leecher_count,omitempty"`
	NextAnnounceTime      int64  `json:"next_announce_time,omitempty"`
	NextScrapeTime        int64  `json:"next_scrape_time,omitempty"`
	Scrape                string `json:"scrape,omitempty"`
	ScrapeState           int64  `json:"scrape_state,omitempty"`
	SeederCount           int64  `json:"seeder_count,omitempty"`
	Sitename              string `json:"sitename,omitempty"`
	Tier                  int64  `json:"tier,omitempty"`
}

var allTorrentFields = structJSONFields[Torrent]()

func (c *Client) TorrentGet(ctx context.Context, ids *TorrentIDs) (*TorrentGetResult, error) {
	args := torrentGetArgs{
		IDs:    ids,
		Fields: allTorrentFields,
		Format: "object",
	}
	return postWithArgs[torrentGetArgs, TorrentGetResult](ctx, c, "torrent-get", args)
}

type TorrentAddArgs struct {
	// Filename or MetaInfo must be set
	Filename *string `json:"filename,omitempty"`
	MetaInfo *string `json:"metainfo,omitempty"`

	Cookies           *string `json:"cookies,omitempty"`
	DownloadDir       *string `json:"download_dir,omitempty"`
	Paused            *bool   `json:"paused,omitempty"`
	PeerLimit         *int64  `json:"peer_limit,omitempty"`
	BandwidthPriority *int64  `json:"bandwidth_priority,omitempty"`
	FilesWanted       []int64 `json:"files_wanted,omitempty"`
	FilesUnwanted     []int64 `json:"files_unwanted,omitempty"`
	PriorityHigh      []int64 `json:"priority_high,omitempty"`
	PriorityLow       []int64 `json:"priority_low,omitempty"`
	PriorityNormal    []int64 `json:"priority_normal,omitempty"`
}

type TorrentAddResult struct {
	TorrentAdded      *TorrentInfo `json:"torrent-added,omitempty"`
	TorrentDuplicated *TorrentInfo `json:"torrent-duplicate,omitempty"`
}

type TorrentInfo struct {
	ID         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	HashString string `json:"hashString,omitempty"`
}

func (c *Client) TorrentAdd(ctx context.Context, args TorrentAddArgs) (*TorrentAddResult, error) {
	return postWithArgs[TorrentAddArgs, TorrentAddResult](ctx, c, "torrent-add", args)
}

type TorrentRemoveArgs struct {
	IDs             *TorrentIDs `json:"ids,omitempty"`
	DeleteLocalData bool        `json:"delete_local_data"`
}

func (c *Client) TorrentRemove(ctx context.Context, args TorrentRemoveArgs) error {
	_, err := postWithArgs[TorrentRemoveArgs, any](ctx, c, "torrent-remove", args)
	return err
}

type TorrentSetLocationArgs struct {
	IDs      *TorrentIDs `json:"ids,omitempty"`
	Location string      `json:"location"`
	Move     bool        `json:"move"`
}

func (c *Client) TorrentSetLocation(ctx context.Context, args TorrentSetLocationArgs) error {
	_, err := postWithArgs[TorrentSetLocationArgs, any](ctx, c, "torrent-set-location", args)
	return err
}

type TorrentRenamePathArgs struct {
	// IDs must be exactly 1 torrent
	IDs  *TorrentIDs `json:"ids,omitempty"`
	Path string      `json:"path"`
	Name bool        `json:"name"`
}

type TorrentRenamePathResult struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`
	Name bool   `json:"name"`
}

func (c *Client) TorrentRenamePath(ctx context.Context, args TorrentRenamePathArgs) (*TorrentRenamePathResult, error) {
	return postWithArgs[TorrentRenamePathArgs, TorrentRenamePathResult](ctx, c, "torrent-rename-path", args)
}
