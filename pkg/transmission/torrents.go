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
	ActivityDate                int64             `json:"activityDate,omitempty"`
	AddedDate                   int64             `json:"addedDate,omitempty"`
	Availability                []int64           `json:"availability,omitempty"`
	BandwidthPriority           int64             `json:"bandwidthPriority,omitempty"`
	BytesCompleted              []int64           `json:"bytesCompleted,omitempty"`
	Comment                     string            `json:"comment,omitempty"`
	CorruptEver                 int64             `json:"corruptEver,omitempty"`
	Creator                     string            `json:"creator,omitempty"`
	DateCreated                 int64             `json:"dateCreated,omitempty"`
	DesiredAvailable            int64             `json:"desiredAvailable,omitempty"`
	DoneDate                    int64             `json:"doneDate,omitempty"`
	DownloadDir                 string            `json:"downloadDir,omitempty"`
	DownloadedEver              int64             `json:"downloadedEver,omitempty"`
	DownloadLimit               int64             `json:"downloadLimit,omitempty"`
	DownloadLimited             bool              `json:"downloadLimited,omitempty"`
	EditDate                    int64             `json:"editDate,omitempty"`
	Error                       int64             `json:"error,omitempty"`
	ErrorString                 string            `json:"errorString,omitempty"`
	ETA                         int64             `json:"eta,omitempty"`
	ETAIdle                     int64             `json:"etaIdle,omitempty"`
	FileCount                   int64             `json:"fileCount,omitempty"`
	Files                       []TorrentFile     `json:"files,omitempty"`
	FileStats                   []TorrentFileStat `json:"fileStats,omitempty"`
	Group                       string            `json:"group,omitempty"`
	HashString                  string            `json:"hashString,omitempty"`
	HaveUnchecked               int64             `json:"haveUnchecked,omitempty"`
	HaveValid                   int64             `json:"haveValid,omitempty"`
	HonorsSessionLimits         bool              `json:"honorsSessionLimits,omitempty"`
	ID                          int64             `json:"id,omitempty"`
	IsFinished                  bool              `json:"isFinished,omitempty"`
	IsPrivate                   bool              `json:"isPrivate,omitempty"`
	IsStalled                   bool              `json:"isStalled,omitempty"`
	Labels                      []string          `json:"labels,omitempty"`
	LeftUntilDone               int64             `json:"leftUntilDone,omitempty"`
	MagnetLink                  string            `json:"magnetLink,omitempty"`
	ManualAnnounceTime          int64             `json:"manualAnnounceTime,omitempty"` // deprecated
	MaxConnectedPeers           int64             `json:"maxConnectedPeers,omitempty"`
	MetadataPercentComplete     float64           `json:"metadataPercentComplete,omitempty"`
	Name                        string            `json:"name,omitempty"`
	PeerLimit                   int64             `json:"peerLimit,omitempty"`
	Peers                       []Peer            `json:"peers,omitempty"`
	PeersConnected              int64             `json:"peersConnected,omitempty"`
	PeersFrom                   *PeersFrom        `json:"peersFrom,omitempty"`
	PeersGettingFromUs          int64             `json:"peersGettingFromUs,omitempty"`
	PeersSendingToUs            int64             `json:"peersSendingToUs,omitempty"`
	PercentComplete             float64           `json:"percentComplete,omitempty"`
	PercentDone                 float64           `json:"percentDone,omitempty"`
	Pieces                      string            `json:"pieces,omitempty"`
	PieceCount                  int64             `json:"pieceCount,omitempty"`
	PieceSize                   int64             `json:"pieceSize,omitempty"`
	Priorities                  []int64           `json:"priorities,omitempty"`
	PrimaryMIMEType             string            `json:"primaryMimeType,omitempty"`
	QueuePosition               int64             `json:"queuePosition,omitempty"`
	RateDownload                int64             `json:"rateDownload,omitempty"` // B/s
	RateUpload                  int64             `json:"rateUpload,omitempty"`   // B/s
	RecheckProgress             float64           `json:"recheckProgress,omitempty"`
	SecondsDownloading          int64             `json:"secondsDownloading,omitempty"`
	SecondsSeeding              int64             `json:"secondsSeeding,omitempty"`
	SeedIdleLimit               int64             `json:"seedIdleLimit,omitempty"`
	SeedIdleMode                int64             `json:"seedIdleMode,omitempty"`
	SeedRatioLimit              float64           `json:"seedRatioLimit,omitempty"`
	SeedRatioMode               int64             `json:"seedRatioMode,omitempty"`
	SequentialDownload          bool              `json:"sequentialDownload,omitempty"`
	SequentialDownloadFromPiece int64             `json:"sequentialDownloadFromPiece,omitempty"`
	SizeWhenDone                int64             `json:"sizeWhenDone,omitempty"`
	StartDate                   int64             `json:"startDate,omitempty"`
	Status                      TorrentStatus     `json:"status,omitempty"` // 0..6
	TorrentFile                 string            `json:"torrentFile,omitempty"`
	TotalSize                   int64             `json:"totalSize,omitempty"`
	Trackers                    []Tracker         `json:"trackers,omitempty"`
	TrackerList                 string            `json:"trackerList,omitempty"`
	TrackerStats                []TrackerStat     `json:"trackerStats,omitempty"`
	UploadedEver                int64             `json:"uploadedEver,omitempty"`
	UploadLimit                 int64             `json:"uploadLimit,omitempty"`
	UploadLimited               bool              `json:"uploadLimited,omitempty"`
	UploadRatio                 float64           `json:"uploadRatio,omitempty"`
	Wanted                      []int64           `json:"wanted,omitempty"` // 0/1 values in 4.x
	Webseeds                    []string          `json:"webseeds,omitempty"`
	WebseedsSendingToUs         int64             `json:"webseedsSendingToUs,omitempty"`
}

type TorrentStatus int

const (
	TorrentStatusStopped      TorrentStatus = 0
	TorrentStatusCheckWait    TorrentStatus = 1
	TorrentStatusCheck        TorrentStatus = 2
	TorrentStatusDownloadWait TorrentStatus = 3
	TorrentStatusDownload     TorrentStatus = 4
	TorrentStatusSeedWait     TorrentStatus = 5
	TorrentStatusSeed         TorrentStatus = 6
)

var TorrentStatusByLabel = map[TorrentStatus]string{
	TorrentStatusStopped:      "stopped",
	TorrentStatusCheckWait:    "check_wait",
	TorrentStatusCheck:        "check",
	TorrentStatusDownloadWait: "download_wait",
	TorrentStatusDownload:     "download",
	TorrentStatusSeedWait:     "seed_wait",
	TorrentStatusSeed:         "seed",
}

func (ts TorrentStatus) String() string {
	label, exists := TorrentStatusByLabel[ts]
	if !exists {
		return "unknown"
	}
	return label
}

type TorrentFile struct {
	BytesCompleted int64  `json:"bytesCompleted,omitempty"`
	Length         int64  `json:"length,omitempty"`
	Name           string `json:"name,omitempty"`
	BeginPiece     int64  `json:"beginPiece,omitempty"`
	EndPiece       int64  `json:"endPiece,omitempty"`
}

type TorrentFileStat struct {
	BytesCompleted int64 `json:"bytesCompleted,omitempty"`
	Wanted         bool  `json:"wanted,omitempty"`   // NOTE: different from Torrent.Wanted (0/1 array)
	Priority       int64 `json:"priority,omitempty"` // tr_priority_t
}

type Peer struct {
	Address            string  `json:"address,omitempty"`
	BytesToClient      int64   `json:"bytesToClient,omitempty"`
	BytesToPeer        int64   `json:"bytesToPeer,omitempty"`
	ClientIsChoked     bool    `json:"clientIsChoked,omitempty"`
	ClientIsInterested bool    `json:"clientIsInterested,omitempty"`
	ClientName         string  `json:"clientName,omitempty"`
	FlagStr            string  `json:"flagStr,omitempty"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom,omitempty"`
	IsEncrypted        bool    `json:"isEncrypted,omitempty"`
	IsIncoming         bool    `json:"isIncoming,omitempty"`
	IsUploadingTo      bool    `json:"isUploadingTo,omitempty"`
	IsUTP              bool    `json:"isUTP,omitempty"`
	PeerID             string  `json:"peerId,omitempty"`
	PeerIsChoked       bool    `json:"peerIsChoked,omitempty"`
	PeerIsInterested   bool    `json:"peerIsInterested,omitempty"`
	Port               int64   `json:"port,omitempty"`
	Progress           float64 `json:"progress,omitempty"`
	RateToClient       int64   `json:"rateToClient,omitempty"`
	RateToPeer         int64   `json:"rateToPeer,omitempty"`
}

type PeersFrom struct {
	FromCache    int64 `json:"fromCache,omitempty"`
	FromDHT      int64 `json:"fromDHT,omitempty"`
	FromIncoming int64 `json:"fromIncoming,omitempty"`
	FromLPD      int64 `json:"fromLPD,omitempty"`
	FromLTEP     int64 `json:"fromLTEP,omitempty"`
	FromPEX      int64 `json:"fromPEX,omitempty"`
	FromTracker  int64 `json:"fromTracker,omitempty"`
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
	AnnounceState         int64  `json:"announceState,omitempty"`
	DownloadCount         int64  `json:"downloadCount,omitempty"`
	DownloaderCount       int64  `json:"downloaderCount,omitempty"`
	HasAnnounced          bool   `json:"hasAnnounced,omitempty"`
	HasScraped            bool   `json:"hasScraped,omitempty"`
	Host                  string `json:"host,omitempty"`
	ID                    int64  `json:"id,omitempty"`
	IsBackup              bool   `json:"isBackup,omitempty"`
	LastAnnouncePeerCount int64  `json:"lastAnnouncePeerCount,omitempty"`
	LastAnnounceResult    string `json:"lastAnnounceResult,omitempty"`
	LastAnnounceStartTime int64  `json:"lastAnnounceStartTime,omitempty"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded,omitempty"`
	LastAnnounceTime      int64  `json:"lastAnnounceTime,omitempty"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut,omitempty"`
	LastScrapeResult      string `json:"lastScrapeResult,omitempty"`
	LastScrapeStartTime   int64  `json:"lastScrapeStartTime,omitempty"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded,omitempty"`
	LastScrapeTime        int64  `json:"lastScrapeTime,omitempty"`
	LastScrapeTimedOut    bool   `json:"lastScrapeTimedOut,omitempty"`
	LeecherCount          int64  `json:"leecherCount,omitempty"`
	NextAnnounceTime      int64  `json:"nextAnnounceTime,omitempty"`
	NextScrapeTime        int64  `json:"nextScrapeTime,omitempty"`
	Scrape                string `json:"scrape,omitempty"`
	ScrapeState           int64  `json:"scrapeState,omitempty"`
	SeederCount           int64  `json:"seederCount,omitempty"`
	Sitename              string `json:"sitename,omitempty"`
	Tier                  int64  `json:"tier,omitempty"`
}

var AllTorrentFields = structJSONFields[Torrent]()

type TorrentGetArgs struct {
	IDs    *TorrentIDs
	Fields []string
}

func (c *Client) TorrentGet(ctx context.Context, args TorrentGetArgs) (*TorrentGetResult, error) {
	params := torrentGetArgs{
		IDs:    args.IDs,
		Fields: args.Fields,
		Format: "object",
	}
	return postWithArgs[torrentGetArgs, TorrentGetResult](ctx, c, "torrent-get", params)
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
