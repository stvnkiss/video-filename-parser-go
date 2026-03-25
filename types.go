package videoparser

import "time"

// Resolution represents video resolution.
type Resolution string

const (
	R2160P Resolution = "2160P"
	R1080P Resolution = "1080P"
	R720P  Resolution = "720P"
	R576P  Resolution = "576P"
	R540P  Resolution = "540P"
	R480P  Resolution = "480P"
)

// Source represents media source.
type Source string

const (
	SourceBLURAY    Source = "BLURAY"
	SourceWEBDL     Source = "WEBDL"
	SourceWEBRIP    Source = "WEBRIP"
	SourceDVD       Source = "DVD"
	SourceCAM       Source = "CAM"
	SourceSCREENER  Source = "SCREENER"
	SourcePPV       Source = "PPV"
	SourceTELESYNC  Source = "TELESYNC"
	SourceTELECINE  Source = "TELECINE"
	SourceWORKPRINT Source = "WORKPRINT"
	SourceTV        Source = "TV"
)

// VideoCodec represents video codec.
type VideoCodec string

const (
	VideoCodecX265 VideoCodec = "x265"
	VideoCodecX264 VideoCodec = "x264"
	VideoCodecH264 VideoCodec = "h264"
	VideoCodecH265 VideoCodec = "h265"
	VideoCodecWMV  VideoCodec = "WMV"
	VideoCodecXVID VideoCodec = "xvid"
	VideoCodecDVDR VideoCodec = "dvdr"
)

// AudioCodec represents audio codec.
type AudioCodec string

const (
	AudioCodecMP3    AudioCodec = "MP3"
	AudioCodecMP2    AudioCodec = "MP2"
	AudioCodecDOLBY  AudioCodec = "Dolby Digital"
	AudioCodecEAC3   AudioCodec = "Dolby Digital Plus"
	AudioCodecAAC    AudioCodec = "AAC"
	AudioCodecFLAC   AudioCodec = "FLAC"
	AudioCodecDTS    AudioCodec = "DTS"
	AudioCodecDTSHD  AudioCodec = "DTS-HD"
	AudioCodecTRUEHD AudioCodec = "Dolby TrueHD"
	AudioCodecOPUS   AudioCodec = "Opus"
	AudioCodecVORBIS AudioCodec = "Vorbis"
	AudioCodecPCM    AudioCodec = "PCM"
	AudioCodecLPCM   AudioCodec = "LPCM"
)

// Channels represents audio channel layout.
type Channels string

const (
	ChannelsSEVEN  Channels = "7.1"
	ChannelsSIX    Channels = "5.1"
	ChannelsSTEREO Channels = "stereo"
	ChannelsMONO   Channels = "mono"
)

// Language represents a language.
type Language string

const (
	LanguageEnglish    Language = "English"
	LanguageFrench     Language = "French"
	LanguageSpanish    Language = "Spanish"
	LanguageGerman     Language = "German"
	LanguageItalian    Language = "Italian"
	LanguageDanish     Language = "Danish"
	LanguageDutch      Language = "Dutch"
	LanguageJapanese   Language = "Japanese"
	LanguageCantonese  Language = "Cantonese"
	LanguageMandarin   Language = "Mandarin"
	LanguageRussian    Language = "Russian"
	LanguagePolish     Language = "Polish"
	LanguageVietnamese Language = "Vietnamese"
	LanguageNordic     Language = "Nordic"
	LanguageSwedish    Language = "Swedish"
	LanguageNorwegian  Language = "Norwegian"
	LanguageFinnish    Language = "Finnish"
	LanguageTurkish    Language = "Turkish"
	LanguagePortuguese Language = "Portuguese"
	LanguageFlemish    Language = "Flemish"
	LanguageGreek      Language = "Greek"
	LanguageKorean     Language = "Korean"
	LanguageHungarian  Language = "Hungarian"
	LanguagePersian    Language = "Persian"
	LanguageBengali    Language = "Bengali"
	LanguageBulgarian  Language = "Bulgarian"
	LanguageBrazilian  Language = "Brazilian"
	LanguageHebrew     Language = "Hebrew"
	LanguageCzech      Language = "Czech"
	LanguageUkrainian  Language = "Ukrainian"
	LanguageCatalan    Language = "Catalan"
	LanguageChinese    Language = "Chinese"
	LanguageThai       Language = "Thai"
	LanguageHindi      Language = "Hindi"
	LanguageTamil      Language = "Tamil"
	LanguageArabic     Language = "Arabic"
	LanguageEstonian   Language = "Estonian"
	LanguageIcelandic  Language = "Icelandic"
	LanguageLatvian    Language = "Latvian"
	LanguageLithuanian Language = "Lithuanian"
	LanguageRomanian   Language = "Romanian"
	LanguageSlovak     Language = "Slovak"
	LanguageSerbian    Language = "Serbian"
)

// QualityModifier represents a quality modifier tag.
type QualityModifier string

const (
	QualityModifierREMUX  QualityModifier = "REMUX"
	QualityModifierBRDISK QualityModifier = "BRDISK"
	QualityModifierRAWHD  QualityModifier = "RAWHD"
)

// Edition represents edition flags detected in a release name.
type Edition struct {
	Internal      bool
	Limited       bool
	Remastered    bool
	Extended      bool
	Theatrical    bool
	Directors     bool
	Unrated       bool
	IMAX          bool
	FanEdit       bool
	HDR           bool
	BW            bool
	ThreeD        bool
	HSBS          bool
	SBS           bool
	HOU           bool
	UHD           bool
	OAR           bool
	DolbyVision   bool
	HardcodedSubs bool
	DeletedScenes bool
	BonusContent  bool
}

// Revision holds proper/real version counters.
type Revision struct {
	Version int
	Real    int
}

// QualityModel holds parsed quality information.
type QualityModel struct {
	Sources    []Source
	Modifier   QualityModifier
	Resolution Resolution
	Revision   Revision
}

// Season holds parsed TV show season/episode information.
type Season struct {
	ReleaseTitle    string
	SeriesTitle     string
	Seasons         []int
	EpisodeNumbers  []int
	AirDate         *time.Time
	FullSeason      bool
	IsPartialSeason bool
	IsMultiSeason   bool
	IsSeasonExtra   bool
	IsSpecial       bool
	SeasonPart      int
}

// ParsedFilename holds all parsed metadata for a movie or TV show filename.
// Check IsTv to know whether TV-specific fields are meaningful.
type ParsedFilename struct {
	// Shared fields
	Title         string
	Year          string
	Edition       Edition
	Resolution    Resolution
	Sources       []Source
	VideoCodec    VideoCodec
	AudioCodec    AudioCodec
	AudioChannels Channels
	Group         string
	Revision      Revision
	Languages     []Language
	Multi         bool
	Complete      bool
	// TV-only fields (populated when IsTv == true)
	IsTv            bool
	Seasons         []int
	EpisodeNumbers  []int
	AirDate         *time.Time
	FullSeason      bool
	IsPartialSeason bool
	IsMultiSeason   bool
	IsSeasonExtra   bool
	IsSpecial       bool
	SeasonPart      int
}
