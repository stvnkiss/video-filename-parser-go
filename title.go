package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var movieTitleYearRegexes = []*regexp2.Regexp{
	// Special, Despecialized, etc. Edition Movies
	regexp2.MustCompile(
		`^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*\(?\b(?<edition>(((Extended.|Ultimate.)?(Director.?s|Collector.?s|Theatrical|Anniversary|The.Uncut|Ultimate|Final(?=(.(Cut|Edition|Version)))|Extended|Rogue|Special|Despecialized|\d{2,3}(th)?.Anniversary)(.(Cut|Edition|Version))?(.(Extended|Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit))?|((Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit|Edition|Restored|((2|3|4)in1))))))\b\)?.{1,3}(?<year>(1(8|9)|20)\d{2}(?!p|i|\d+|\]|\W\d+)))+(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase,
	),
	// Folder movie format: Blade Runner 2049 (2017)
	regexp2.MustCompile(
		`^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*\((?<year>(1(8|9)|20)\d{2}(?!p|i|(1(8|9)|20)\d{2}|\]|\W(1(8|9)|20)\d{2})))+`,
		regexp2.IgnoreCase,
	),
	// Normal movie format: Mission.Impossible.3.2011
	regexp2.MustCompile(
		`^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*(?<year>(1(8|9)|20)\d{2}(?!p|i|(1(8|9)|20)\d{2}|\]|\W(1(8|9)|20)\d{2})))+(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase,
	),
	// PassThePopcorn: Star.Wars[PassThePopcorn]
	regexp2.MustCompile(
		`^(?<title>.+?)?(?:(?:[-_\W](?<![()[!]))*(?<year>(\[\w *\])))+(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase,
	),
	// [] for years fallback
	regexp2.MustCompile(
		`^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)!]))*(?<year>(1(8|9)|20)\d{2}(?!p|i|\d+|\W\d+)))+(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase,
	),
	// Last resort: title has ( or [
	regexp2.MustCompile(
		`^(?<title>.+?)?(?:(?:[-_\W](?<![)[!]))*(?<year>(1(8|9)|20)\d{2}(?!p|i|\d+|\]|\W\d+)))+(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase,
	),
}

// TitleAndYear holds the parsed title and optional year.
type TitleAndYear struct {
	Title string
	Year  string
}

// ParseTitleAndYear extracts the movie/show title and publication year from a filename.
func ParseTitleAndYear(title string) TitleAndYear {
	simpleTitle := SimplifyTitle(title)
	grouplessTitle, _ := regexp2.MustCompile(`-([a-z0-9]+)$`, regexp2.IgnoreCase).Replace(simpleTitle, "", -1, -1)

	for _, exp := range movieTitleYearRegexes {
		m, err := exp.FindStringMatch(grouplessTitle)
		if err != nil || m == nil {
			continue
		}
		tg := m.GroupByName("title")
		if tg == nil || tg.Length == 0 {
			continue
		}
		cleaned := ReleaseTitleCleaner(tg.String())
		if cleaned == "" {
			continue
		}
		var year string
		yg := m.GroupByName("year")
		if yg != nil && yg.Length > 0 {
			year = yg.String()
		}
		return TitleAndYear{Title: cleaned, Year: year}
	}

	// Fallback: find earliest codec/resolution/channel artifact position
	resResult := parseResolution(title)
	resPos := -1
	if resResult.Source != "" {
		resPos = strings.Index(title, resResult.Source)
	}

	vcResult := parseVideoCodec(title)
	vcPos := -1
	if vcResult.Source != "" {
		vcPos = strings.Index(title, vcResult.Source)
	}

	chResult := parseAudioChannels(title)
	chPos := -1
	if chResult.Source != "" {
		chPos = strings.Index(title, chResult.Source)
	}

	acResult := parseAudioCodec(title)
	acPos := -1
	if acResult.Source != "" {
		acPos = strings.Index(title, acResult.Source)
	}

	minPos := -1
	for _, pos := range []int{resPos, vcPos, chPos, acPos} {
		if pos > 0 && (minPos < 0 || pos < minPos) {
			minPos = pos
		}
	}

	if minPos > 0 {
		cleaned := ReleaseTitleCleaner(title[:minPos])
		return TitleAndYear{Title: cleaned}
	}

	return TitleAndYear{Title: strings.TrimSpace(title)}
}
