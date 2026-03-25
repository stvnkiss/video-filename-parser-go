package videoparser

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

// reportTitleExp holds all the season/episode detection patterns, in priority order.
var reportTitleExp = []*regexp2.Regexp{
	// Daily episodes without title (2018-10-12, 20181012)
	regexp2.MustCompile(`^(?<airyear>19[6-9]\d|20\d\d)(?<sep>[-_]?)(?<airmonth>0\d|1[0-2])\k<sep>(?<airday>[0-2]\d|3[01])(?!\d)`, regexp2.IgnoreCase),
	// Multi-Part episodes without a title (S01E05.S01E06)
	regexp2.MustCompile(`^(?:\W*S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:(?:[ex]){1,2}(?<episode>\d{1,3}(?!\d+)))+){2,}`, regexp2.IgnoreCase),
	// Multi-episode with single episode numbers (S6.E1-E2, S6.E1E2, S6E1E2)
	regexp2.MustCompile(`^(?<title>.+?)[-_. ]S(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:[E-_. ]?[ex]?(?<episode>(?<!\d+)\d{1,2}(?!\d+)))+(?:[-_. ]?[ex]?(?<episode1>(?<!\d+)\d{1,2}(?!\d+)))+`, regexp2.IgnoreCase),
	// Multi-Episode with a title (S01E05E06, S01E05-06) with trailing info in slashes
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+S?(?<season>(?<!\d+)(?:\d{1,2})(?!\d+))(?:[ex]|\W[ex]|_){1,2}(?<episode>\d{2,3}(?!\d+))(?:(?:-|[ex]|\W[ex]|_){1,2}(?<episode1>\d{2,3}(?!\d+)))+).+?(?:\[.+?\])(?!\\)`, regexp2.IgnoreCase),
	// Episodes without a title, Multi (S01E04E05, 1x04x05)
	regexp2.MustCompile(`(?:S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:(?:[-_]|[ex]){1,2}(?<episode>\d{2,3}(?!\d+))){2,})`, regexp2.IgnoreCase),
	// Episodes without a title, Single (S01E05, 1x05)
	regexp2.MustCompile(`^(?:S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:(?:[-_ ]?[ex])(?<episode>\d{2,3}(?!\d+))))`, regexp2.IgnoreCase),
	// Anime - [SubGroup] Title Episode Absolute Episode Number
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)(?<title>.+?)[-_. ](?:Episode)(?:[-_. ]+(?<absoluteepisode>(?<!\d+)\d{2,3}(\.\d{1,2})?(?!\d+)))+(?:_|-|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`, regexp2.IgnoreCase),
	// Anime - [SubGroup] Title Absolute + Season+Episode
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\](?:_|-|\s|\.)?)(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+(?<absoluteepisode>\d{2,3}(\.\d{1,2})?))+(?:_|-|\s|\.)+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:-|[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+).*?(?<hash>[([]\w{8}[)\]])?(?:$|\.)`, regexp2.IgnoreCase),
	// Anime - [SubGroup] Title Season+Episode + Absolute Episode Number
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\](?:_|-|\s|\.)?)(?<title>.+?)(?:[-_\W](?<![()[!]))+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:-|[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+)(?:(?:_|-|\s|\.)+(?<absoluteepisode>(?<!\d+)\d{2,3}(\.\d{1,2})?(?!\d+)))+.*?(?<hash>\[\w{8}\])?(?:$|\.)`, regexp2.IgnoreCase),
	// Anime - [SubGroup] Title Season+Episode
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\](?:_|-|\s|\.)?)(?<title>.+?)(?:[-_\W](?<![()[!]))+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+)(?:\s|\.).*?(?<hash>\[\w{8}\])?(?:$|\.)`, regexp2.IgnoreCase),
	// Anime - [SubGroup] Title with trailing number Absolute Episode Number
	regexp2.MustCompile(`^\[(?<subgroup>.+?)\][-_. ]?(?<title>[^-]+?\d+?)[-_. ]+(?:[-_. ]?(?<absoluteepisode>\d{3}(\.\d{1,2})?(?!\d+)))+(?:[-_. ]+(?<special>special|ova|ovd))?.*?(?<hash>\[\w{8}\])?(?:$|\.mkv)`, regexp2.IgnoreCase),
	// Anime - [SubGroup] Title - Absolute Episode Number
	regexp2.MustCompile(`^\[(?<subgroup>.+?)\][-_. ]?(?<title>.+?)(?:[. ]-[. ](?<absoluteepisode>\d{2,3}(\.\d{1,2})?(?!\d+|[-])))+(?:[-_. ]+(?<special>special|ova|ovd))?.*?(?<hash>\[\w{8}\])?(?:$|\.mkv)`, regexp2.IgnoreCase),
	// Anime - [SubGroup] Title Absolute Episode Number
	regexp2.MustCompile(`^\[(?<subgroup>.+?)\][-_. ]?(?<title>.+?)[-_. ]+\(?(?:[-_. ]?#?(?<absoluteepisode>\d{2,3}(\.\d{1,2})?(?!\d+)))+\)?(?:[-_. ]+(?<special>special|ova|ovd))?.*?(?<hash>\[\w{8}\])?(?:$|\.mkv)`, regexp2.IgnoreCase),
	// Multi-episode Repeated (S01E05 - S01E06, 1x05 - 1x06)
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:(?:[ex]|[-_. ]e){1,2}(?<episode>\d{1,3}(?!\d+)))+){2,}`, regexp2.IgnoreCase),
	// Single episodes with a title (S01E05, 1x05)
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+S?(?<season>(?<!\d+)(?:\d{1,2})(?!\d+))(?:[ex]|\W[ex]|_){1,2}(?<episode>(?!265|264)\d{2,3}(?!\d+|(?:[ex]|\W[ex]|_|-){1,2})))`, regexp2.IgnoreCase),
	// Anime - Title Season EpisodeNumber + Absolute Episode Number [SubGroup]
	regexp2.MustCompile(`^(?<title>.+?)(?:[-_\W](?<![()[!]))+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:[ex]|\W[ex]){1,2}(?<episode>(?<!\d+)\d{2}(?!\d+)))).+?(?:[-_. ]?(?<absoluteepisode>(?<!\d+)\d{3}(\.\d{1,2})?(?!\d+)))+.+?\[(?<subgroup>.+?)\](?:$|\.mkv)`, regexp2.IgnoreCase),
	// Anime - Title Absolute Episode Number [SubGroup] [Hash]?
	regexp2.MustCompile(`^(?<title>.+?)[-_. ]Episode(?:[-_. ]+(?<absoluteepisode>\d{2,3}(\.\d{1,2})?(?!\d+)))+(?:.+?)\[(?<subgroup>.+?)\].*?(?<hash>\[\w{8}\])?(?:$|\.)`, regexp2.IgnoreCase),
	// Anime - Title Absolute Episode Number [SubGroup] [Hash]
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:_|-|\s|\.)+(?<absoluteepisode>\d{3}(\.\d{1,2})(?!\d+)))+(?:.+?)\[(?<subgroup>.+?)\].*?(?<hash>\[\w{8}\])?(?:$|\.)`, regexp2.IgnoreCase),
	// Anime - Title Absolute Episode Number [Hash]
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:_|-|\s|\.)+(?<absoluteepisode>\d{2,3}(\.\d{1,2})?(?!\d+)))+(?:[-_. ]+(?<special>special|ova|ovd))?[-_. ]+.*?(?<hash>\[\w{8}\])(?:$|\.)`, regexp2.IgnoreCase),
	// Episodes with airdate AND season/episode number, capture season/episode only
	regexp2.MustCompile(`^(?<title>.+?)?\W*(?<airdate>\d{4}\W+[0-1][0-9]\W+[0-3][0-9])(?!\W+[0-3][0-9])[-_. ](?:s?(?<season>(?<!\d+)(?:\d{1,2})(?!\d+)))(?:[ex](?<episode>(?<!\d+)(?:\d{1,3})(?!\d+)))`, regexp2.IgnoreCase),
	// Episodes with airdate AND season/episode number
	regexp2.MustCompile(`^(?<title>.+?)?\W*(?<airyear>\d{4})\W+(?<airmonth>[0-1][0-9])\W+(?<airday>[0-3][0-9])(?!\W+[0-3][0-9]).+?(?:s?(?<season>(?<!\d+)(?:\d{1,2})(?!\d+)))(?:[ex](?<episode>(?<!\d+)(?:\d{1,3})(?!\d+)))`, regexp2.IgnoreCase),
	// Episodes 4 digit season, S2016E05
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+S(?<season>(?<!\d+)(?:\d{4})(?!\d+))(?:e|\We|_){1,2}(?<episode>\d{2,3}(?!\d+))(?:(?:-|e|\We|_){1,2}(?<episode1>\d{2,3}(?!\d+)))*)\W?(?!\\)`, regexp2.IgnoreCase),
	// Episodes 4 digit season, 2016x05
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+(?<season>(?<!\d+)(?:\d{4})(?!\d+))(?:x|\Wx){1,2}(?<episode>\d{2,3}(?!\d+))(?:(?:-|x|\Wx|_){1,2}(?<episode1>\d{2,3}(?!\d+)))*)\W?(?!\\)`, regexp2.IgnoreCase),
	// Multi-season pack
	regexp2.MustCompile(`^(?<title>.+?)[-_. ]+S(?<season>(?<!\d+)(?:\d{1,2})(?!\d+))\W?-\W?S?(?<season1>(?<!\d+)(?:\d{1,2})(?!\d+))`, regexp2.IgnoreCase),
	// Partial season pack
	regexp2.MustCompile(`^(?<title>.+?)(?:\W+S(?<season>(?<!\d+)(?:\d{1,2})(?!\d+))\W+(?:(?:Part\W?|(?<!\d+\W+)e)(?<seasonpart>\d{1,2}(?!\d+)))+)`, regexp2.IgnoreCase),
	// Mini-Series with year in title, Part01, Part 01, Part.1
	regexp2.MustCompile(`^(?<title>.+?\d{4})(?:\W+(?:(?:Part\W?|e)(?<episode>\d{1,2}(?!\d+)))+)`, regexp2.IgnoreCase),
	// Mini-Series multi episodes E1-E2
	regexp2.MustCompile(`^(?<title>.+?)(?:[-._ ][e])(?<episode>\d{2,3}(?!\d+))(?:(?:-?[e])(?<episode1>\d{2,3}(?!\d+)))+`, regexp2.IgnoreCase),
	// Mini-Series Part01, Part 01, Part.1
	regexp2.MustCompile(`^(?<title>.+?)(?:\W+(?:(?:Part\W?|(?<!\d+\W+)e)(?<episode>\d{1,2}(?!\d+)))+)`, regexp2.IgnoreCase),
	// Mini-Series Part One/Two/...Nine
	regexp2.MustCompile(`^(?<title>.+?)(?:\W+(?:Part[-._ ](?<episode>One|Two|Three|Four|Five|Six|Seven|Eight|Nine)(>[-._ ])))`, regexp2.IgnoreCase),
	// Mini-Series XofY
	regexp2.MustCompile(`^(?<title>.+?)(?:\W+(?:(?<episode>(?<!\d+)\d{1,2}(?!\d+))of\d+)+)`, regexp2.IgnoreCase),
	// Supports Season 01 Episode 03
	regexp2.MustCompile(`(?:.*(?:""|^))(?<title>.*?)(?:[-_\W](?<![()[]))+(?:\W?Season\W?)(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:\W|_)+(?:Episode\W)(?:[-_. ]?(?<episode>(?<!\d+)\d{1,2}(?!\d+)))+`, regexp2.IgnoreCase),
	// Multi-episode in square brackets [S01E11E12] or [S01E11-12]
	regexp2.MustCompile(`(?:.*(?:^))(?<title>.*?)[-._ ]+\[S(?<season>(?<!\d+)\d{2}(?!\d+))(?:[E-]{1,2}(?<episode>(?<!\d+)\d{2}(?!\d+)))+\]`, regexp2.IgnoreCase),
	// Multi-episode no space (S01E11E12)
	regexp2.MustCompile(`(?:.*(?:^))(?<title>.*?)S(?<season>(?<!\d+)\d{2}(?!\d+))(?:E(?<episode>(?<!\d+)\d{2}(?!\d+)))+`, regexp2.IgnoreCase),
	// Single episode S1E1 or S1-E1 or S1.Ep1 or S01.Ep.01
	regexp2.MustCompile(`(?:.*(?:""|^))(?<title>.*?)(?:\W?|_)S(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:\W|_)?Ep?[ ._]?(?<episode>(?<!\d+)\d{1,2}(?!\d+))`, regexp2.IgnoreCase),
	// 3 digit season S010E05
	regexp2.MustCompile(`(?:.*(?:""|^))(?<title>.*?)(?:\W?|_)S(?<season>(?<!\d+)\d{3}(?!\d+))(?:\W|_)?E(?<episode>(?<!\d+)\d{1,2}(?!\d+))`, regexp2.IgnoreCase),
	// 5 digit episode with title
	regexp2.MustCompile(`^(?:(?<title>.+?)(?:_|-|\s|\.)+)(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+)))(?:(?:-|[ex]|\W[ex]|_){1,2}(?<episode>(?<!\d+)\d{5}(?!\d+)))`, regexp2.IgnoreCase),
	// 5 digit multi-episode with title
	regexp2.MustCompile(`^(?:(?<title>.+?)(?:_|-|\s|\.)+)(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+)))(?:(?:[-_. ]{1,3}ep){1,2}(?<episode>(?<!\d+)\d{5}(?!\d+)))+`, regexp2.IgnoreCase),
	// Separated S01 - E01
	regexp2.MustCompile(`^(?<title>.+?)(?:_|-|\s|\.)+S(?<season>\d{2}(?!\d+))(\W-\W)E(?<episode>(?<!\d+)\d{2}(?!\d+))(?!\\)`, regexp2.IgnoreCase),
	// Anime - Title with season number - Absolute Episode Number (Title S01 - EP14)
	regexp2.MustCompile(`^(?<title>.+?S\d{1,2})[-_. ]{3,}(?:EP)?(?<absoluteepisode>\d{2,3}(\.\d{1,2})?(?!\d+|[-]))`, regexp2.IgnoreCase),
	// Anime - French titles with single episode numbers
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)[-_. ]+?(?:Episode[-_. ]+?)(?<absoluteepisode>\d{1}(\.\d{1,2})?(?!\d+))`, regexp2.IgnoreCase),
	// Season only releases
	regexp2.MustCompile(`^(?<title>.+?)\W(?:S|Season)\W?(?<season>\d{1,2}(?!\d+))(\W+|_|$)(?<extras>EXTRAS|SUBPACK)?(?!\\)`, regexp2.IgnoreCase),
	// 4 digit season only releases
	regexp2.MustCompile(`^(?<title>.+?)\W(?:S|Season)\W?(?<season>\d{4}(?!\d+))(\W+|_|$)(?<extras>EXTRAS|SUBPACK)?(?!\\)`, regexp2.IgnoreCase),
	// Episodes in square brackets [S01E05-06]
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+\[S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:-|[ex]|\W[ex]|_){1,2}(?<episode>(?<!\d+)\d{2}(?!\d+|i|p)))+\])\W?(?!\\)`, regexp2.IgnoreCase),
	// Supports 103/113 naming
	regexp2.MustCompile(`^(?<title>.+?)?(?:(?:[_.](?<![()[!]))+(?<season>(?<!\d+)[1-9])(?<episode>[1-9][0-9]|[0][1-9])(?![a-z]|\d+))+(?:[_.]|$)`, regexp2.IgnoreCase),
	// 4 digit episode, no title (S01E0500)
	regexp2.MustCompile(`^(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:-|[ex]|\W[ex]|_){1,2}(?<episode>\d{4}(?!\d+|i|p)))+)(\W+|_|$)(?!\\)`, regexp2.IgnoreCase),
	// 4 digit episode, with title
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:-|[ex]|\W[ex]|_){1,2}(?<episode>\d{4}(?!\d+|i|p)))+)\W?(?!\\)`, regexp2.IgnoreCase),
	// Episodes with airdate (2018.04.28)
	regexp2.MustCompile(`^(?<title>.+?)?\W*(?<airyear>\d{4})[-_. ]+(?<airmonth>[0-1][0-9])[-_. ]+(?<airday>[0-3][0-9])(?![-_. ]+[0-3][0-9])`, regexp2.IgnoreCase),
	// Episodes with airdate (04.28.2018)
	regexp2.MustCompile(`^(?<title>.+?)?\W*(?<airmonth>[0-1][0-9])[-_. ]+(?<airday>[0-3][0-9])[-_. ]+(?<airyear>\d{4})(?!\d+)`, regexp2.IgnoreCase),
	// Supports 1103/1113 naming
	regexp2.MustCompile(`^(?<title>.+?)?(?:(?:[-_\W](?<![()[!]))*(?<season>(?<!\d+|\(|\[|e|x)\d{2})(?<episode>(?<!e|x)\d{2}(?!p|i|\d+|\)|\]|\W\d+|\W(?:e|ep|x)\d+)))+(\W+|_|$)(?!\\)`, regexp2.IgnoreCase),
	// Episodes with single digit episode number (S01E1, S01E5E6)
	regexp2.MustCompile(`^(?<title>.*?)(?:(?:[-_\W](?<![()[!]))+S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:-|[ex]){1,2}(?<episode>\d{1}))+)+(\W+|_|$)(?!\\)`, regexp2.IgnoreCase),
	// iTunes Season 1\05 Title
	regexp2.MustCompile(`^(?:Season(?:_|-|\s|\.)(?<season>(?<!\d+)\d{1,2}(?!\d+)))(?:_|-|\s|\.)(?<episode>(?<!\d+)\d{1,2}(?!\d+))`, regexp2.IgnoreCase),
	// iTunes 1-05 Title
	regexp2.MustCompile(`^(?:(?<season>(?<!\d+)(?:\d{1,2})(?!\d+))(?:-(?<episode>\d{2,3}(?!\d+))))`, regexp2.IgnoreCase),
	// Anime Range - Title Absolute Episode Number (ep01-12)
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)(?:_|\s|\.)+(?:e|ep)(?<absoluteepisode>\d{2,3}(\.\d{1,2})?)-(?<absoluteepisode1>(?<!\d+)\d{1,2}(\.\d{1,2})?(?!\d+|-)).*?(?<hash>\[\w{8}\])?(?:$|\.)`, regexp2.IgnoreCase),
	// Anime - Title Absolute Episode Number (e66)
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)(?:(?:_|-|\s|\.)+(?:e|ep)(?<absoluteepisode>\d{2,4}(\.\d{1,2})?))+.*?(?<hash>\[\w{8}\])?(?:$|\.)`, regexp2.IgnoreCase),
	// Anime - Title Episode Absolute Episode Number
	regexp2.MustCompile(`^(?<title>.+?)[-_. ](?:Episode)(?:[-_. ]+(?<absoluteepisode>(?<!\d+)\d{2,3}(\.\d{1,2})?(?!\d+)))+(?:_|-|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`, regexp2.IgnoreCase),
	// Anime Range 1 or 2 digit (1-10)
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)[_. ]+(?<absoluteepisode>(?<!\d+)\d{1,2}(\.\d{1,2})?(?!\d+))-(?<absoluteepisode1>(?<!\d+)\d{1,2}(\.\d{1,2})?(?!\d+|-))(?:_|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`, regexp2.IgnoreCase),
	// Anime - Title Absolute Episode Number
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)(?:[-_. ]+(?<absoluteepisode>(?<!\d+)\d{2,3}(\.\d{1,2})?(?!\d+)))+(?:_|-|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`, regexp2.IgnoreCase),
	// Anime - Title {Absolute Episode Number}
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)(?:(?:[-_\W](?<![()[!]))+(?<absoluteepisode>(?<!\d+)\d{2,3}(\.\d{1,2})?(?!\d+)))+(?:_|-|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`, regexp2.IgnoreCase),
	// Extant, terrible multi-episode naming (extant.10708.hdtv-lol.mp4)
	regexp2.MustCompile(`^(?<title>.+?)[-_. ](?<season>[0]?\d?)(?:(?<episode>\d{2}){2}(?!\d+))[-_. ]`, regexp2.IgnoreCase),
}

// rejectedRegexes pre-validate the title to filter out hash/garbage strings.
// NOTE: Several of these have stray quote characters that mirror the original source.
var rejectedRegexes = []*regexp2.Regexp{
	regexp2.MustCompile(`^[0-9a-zA-Z]{32}`, regexp2.IgnoreCase),
	regexp2.MustCompile(`^[a-z0-9]{24}$`, regexp2.IgnoreCase),
	// The two patterns below have a literal " as in the original TS source (they rarely match)
	regexp2.MustCompile(`"^[A-Z]{11}\d{3}$`, regexp2.IgnoreCase),
	regexp2.MustCompile(`"^[a-z]{12}\d{3}$`, regexp2.IgnoreCase),
	regexp2.MustCompile(`^Backup_\d{5,}S\d{2}-\d{2}$`, regexp2.IgnoreCase),
	regexp2.MustCompile(`^123$"`, regexp2.None),
	regexp2.MustCompile(`^abc$"`, regexp2.IgnoreCase),
	regexp2.MustCompile(`^b00bs$"`, regexp2.IgnoreCase),
	regexp2.MustCompile(`^\d{6}_\d{2}$"`, regexp2.None),
}

var (
	requestInfoSeasonExp    = regexp2.MustCompile(`^\[.+?\]+`, regexp2.None)
	sixDigitAirDateMatchExp = regexp2.MustCompile(
		`"(?<=[ _.-])(?<airdate>(?<!\d)(?<airyear>[1-9]\d{1})(?<airmonth>[0-1][0-9])(?<airday>[0-3][0-9]))(?=[ _.-])`,
		regexp2.IgnoreCase,
	)
)

// ParseSeason parses TV season/episode information from a release title.
// Returns nil if the title is rejected or no pattern matches.
func ParseSeason(title string) (*Season, error) {
	if !preValidation(title) {
		return nil, nil
	}

	simpleTitle := SimplifyTitle(title)

	// Handle 6-digit air dates like 140722 → 2014.07.22
	m, _ := sixDigitAirDateMatchExp.FindStringMatch(title)
	if m != nil {
		ayg := m.GroupByName("airyear")
		amg := m.GroupByName("airmonth")
		adg := m.GroupByName("airday")
		adtg := m.GroupByName("airdate")
		if ayg != nil && ayg.Length > 0 && amg != nil && amg.Length > 0 && adg != nil && adg.Length > 0 {
			airMonth := amg.String()
			airDay := adg.String()
			if airMonth != "00" || airDay != "00" {
				fixedDate := fmt.Sprintf("20%s.%s.%s", ayg.String(), airMonth, airDay)
				if adtg != nil && adtg.Length > 0 {
					simpleTitle = strings.ReplaceAll(simpleTitle, adtg.String(), fixedDate)
				}
			}
		}
	}

	for _, exp := range reportTitleExp {
		em, err := exp.FindStringMatch(simpleTitle)
		if err != nil || em == nil {
			continue
		}

		result, err := parseMatchCollection(em, simpleTitle)
		if err != nil || result == nil {
			continue
		}

		if result.fullSeason && result.releaseTokens != "" {
			ok, _ := regexp2.MustCompile(`Special`, regexp2.IgnoreCase).MatchString(result.releaseTokens)
			if ok {
				result.fullSeason = false
				result.isSpecial = true
			}
		}

		return &Season{
			ReleaseTitle:    title,
			SeriesTitle:     result.seriesName,
			Seasons:         result.seasonNumbers,
			EpisodeNumbers:  result.episodeNumbers,
			AirDate:         result.airDate,
			FullSeason:      result.fullSeason,
			IsPartialSeason: result.isPartialSeason,
			IsMultiSeason:   result.isMultiSeason,
			IsSeasonExtra:   result.isSeasonExtra,
			IsSpecial:       result.isSpecial,
			SeasonPart:      result.seasonPart,
		}, nil
	}

	return nil, nil
}

func preValidation(title string) bool {
	for _, re := range rejectedRegexes {
		ok, _ := re.MatchString(title)
		if ok {
			return false
		}
	}
	return true
}

// completeRange fills in integers between the first and last values.
func completeRange(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	seen := make(map[int]bool)
	var uniq []int
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			uniq = append(uniq, v)
		}
	}
	first := uniq[0]
	last := uniq[len(uniq)-1]
	if first > last {
		return arr
	}
	result := make([]int, last-first+1)
	for i := range result {
		result[i] = first + i
	}
	return result
}

func indexOfEnd(s, sub string) int {
	idx := strings.Index(s, sub)
	if idx == -1 {
		return -1
	}
	return idx + len(sub)
}

type parsedMatchCollection struct {
	seriesName      string
	seasonNumbers   []int
	episodeNumbers  []int
	isSpecial       bool
	isSeasonExtra   bool
	seasonPart      int
	isPartialSeason bool
	fullSeason      bool
	isMultiSeason   bool
	airDate         *time.Time
	releaseTokens   string
}

func groupStr(m *regexp2.Match, name string) string {
	g := m.GroupByName(name)
	if g == nil || g.Length == 0 {
		return ""
	}
	return g.String()
}

func parseMatchCollection(m *regexp2.Match, simpleTitle string) (*parsedMatchCollection, error) {
	titleStr := groupStr(m, "title")
	cleanedTitle := strings.ReplaceAll(strings.ReplaceAll(titleStr, ".", " "), "_", " ")
	cleanedTitle, _ = requestInfoSeasonExp.Replace(cleanedTitle, "", -1, -1)
	seriesName := strings.TrimSpace(cleanedTitle)

	result := &parsedMatchCollection{seriesName: seriesName}

	lastIdx := indexOfEnd(simpleTitle, titleStr)

	airYearStr := groupStr(m, "airyear")
	airYear, _ := strconv.Atoi(airYearStr)

	if airYear < 1900 || airYear == 0 {
		// Season/episode branch
		seasonStr := groupStr(m, "season")
		season1Str := groupStr(m, "season1")

		var seasons []int
		if seasonStr != "" {
			if end := indexOfEnd(simpleTitle, seasonStr); end > lastIdx {
				lastIdx = end
			}
			if v, err := strconv.Atoi(seasonStr); err == nil {
				seasons = append(seasons, v)
			}
		}
		if season1Str != "" {
			if end := indexOfEnd(simpleTitle, season1Str); end > lastIdx {
				lastIdx = end
			}
			if v, err := strconv.Atoi(season1Str); err == nil {
				seasons = append(seasons, v)
			}
		}

		if len(seasons) > 1 {
			seasons = completeRange(seasons)
			result.isMultiSeason = true
		}
		result.seasonNumbers = seasons

		// Episode captures
		episodeStr := groupStr(m, "episode")
		episode1Str := groupStr(m, "episode1")
		absoluteStr := groupStr(m, "absoluteepisode")
		absolute1Str := groupStr(m, "absoluteepisode1")

		var episodeCaptures []string
		for _, s := range []string{episodeStr, episode1Str} {
			if s != "" {
				episodeCaptures = append(episodeCaptures, s)
			}
		}
		var absoluteCaptures []string
		for _, s := range []string{absoluteStr, absolute1Str} {
			if s != "" {
				absoluteCaptures = append(absoluteCaptures, s)
			}
		}

		if len(episodeCaptures) > 0 {
			first, _ := strconv.ParseFloat(episodeCaptures[0], 64)
			last, _ := strconv.ParseFloat(episodeCaptures[len(episodeCaptures)-1], 64)
			if first > last {
				return nil, nil
			}
			count := int(last-first) + 1
			eps := make([]int, count)
			for i := range eps {
				eps[i] = int(first) + i
			}
			result.episodeNumbers = eps
		}

		if len(absoluteCaptures) > 0 {
			firstF, _ := strconv.ParseFloat(absoluteCaptures[0], 64)
			var lastF float64
			if len(episodeCaptures) > 0 && len(absoluteCaptures) > len(episodeCaptures)-1 {
				lastF = firstF
			} else {
				lastF, _ = strconv.ParseFloat(absoluteCaptures[len(absoluteCaptures)-1], 64)
			}

			if math.Mod(firstF, 1) != 0 || math.Mod(lastF, 1) != 0 {
				if len(absoluteCaptures) != 1 {
					return nil, nil
				}
				result.episodeNumbers = []int{int(firstF)}
				result.isSpecial = true
				if end := indexOfEnd(simpleTitle, absoluteCaptures[0]); end > lastIdx {
					lastIdx = end
				}
			} else {
				count := int(lastF-firstF) + 1
				eps := make([]int, count)
				for i := range eps {
					eps[i] = int(firstF) + i
				}
				result.episodeNumbers = eps
				if groupStr(m, "special") != "" {
					result.isSpecial = true
				}
			}
		}

		if len(episodeCaptures) == 0 && len(absoluteCaptures) == 0 {
			if groupStr(m, "extras") != "" {
				result.isSeasonExtra = true
			}
			if sp := groupStr(m, "seasonpart"); sp != "" {
				v, _ := strconv.Atoi(sp)
				result.seasonPart = v
				result.isPartialSeason = true
			} else {
				result.fullSeason = true
			}
		}

		if len(absoluteCaptures) > 0 && result.episodeNumbers == nil {
			result.seasonNumbers = []int{0}
		}
	} else {
		// Air-date branch
		airMonthStr := groupStr(m, "airmonth")
		airDayStr := groupStr(m, "airday")

		airMonth, _ := strconv.Atoi(airMonthStr)
		airDay, _ := strconv.Atoi(airDayStr)

		if airMonth > 12 {
			airMonth, airDay = airDay, airMonth
		}

		t := time.Date(airYear, time.Month(airMonth), airDay, 0, 0, 0, 0, time.UTC)
		if t.After(time.Now()) {
			return nil, fmt.Errorf("parsed date is in the future")
		}
		epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		if t.Before(epoch) {
			return nil, fmt.Errorf("parsed date error")
		}

		for _, part := range []string{airYearStr, airMonthStr, airDayStr} {
			if end := indexOfEnd(simpleTitle, part); end > lastIdx {
				lastIdx = end
			}
		}
		result.airDate = &t
	}

	if lastIdx == len(simpleTitle) || lastIdx == -1 {
		result.releaseTokens = simpleTitle
	} else {
		result.releaseTokens = simpleTitle[lastIdx:]
	}

	return result, nil
}

// requestInfoSeasonExpHelper is used inside parseMatchCollection to strip [REQ] etc.
func init() {
	// Ensure the regexp is compiled at startup (panics if invalid).
	_ = requestInfoSeasonExp
}
