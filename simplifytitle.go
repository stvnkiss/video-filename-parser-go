package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var (
	simpleTitleRegex = regexp2.MustCompile(
		`\s*(?:480[ip]|576[ip]|720[ip]|1080[ip]|2160[ip]|HVEC|[xh][\W_]?26[45]|DD\W?5\W1|[<>?*:|]|848x480|1280x720|1920x1080)((8|10)b(it))?`,
		regexp2.IgnoreCase,
	)
	websitePrefixRegex = regexp2.MustCompile(
		`^\[\s*[a-z]+(\.[a-z]+)+\s*\][- ]*|^www\.[a-z]+\.(?:com|net)[ -]*`,
		regexp2.IgnoreCase,
	)
	cleanTorrentPrefixRegex = regexp2.MustCompile(`^\[(?:REQ)\]`, regexp2.IgnoreCase)
	cleanTorrentSuffixRegex = regexp2.MustCompile(`\[(?:ettv|rartv|rarbg|cttv)\]$`, regexp2.IgnoreCase)
	commonSourcesRegex      = regexp2.MustCompile(
		`\b(Bluray|(dvdr?|BD)rip|HDTV|HDRip|TS|R5|CAM|SCR|(WEB|DVD)?.?SCREENER|DiVX|xvid|web-?dl)\b`,
		regexp2.IgnoreCase,
	)
	codecGlobalExp         = regexp2.MustCompile(codecExp.String(), regexp2.IgnoreCase)
	commonSourcesGlobalExp = regexp2.MustCompile(commonSourcesRegex.String(), regexp2.IgnoreCase)

	// releaseTitleCleaner regexes
	requestInfoRegex    = regexp2.MustCompile(`\[.+?\]`, regexp2.IgnoreCase)
	editionExp          = regexp2.MustCompile(`\b((Extended.|Ultimate.)?(Director.?s|Collector.?s|Theatrical|Anniversary|The.Uncut|DC|Ultimate|Final(?=(.(Cut|Edition|Version)))|Extended|Special|Despecialized|unrated|\d{2,3}(th)?.Anniversary)(.(Cut|Edition|Version))?(.(Extended|Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit))?|((Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit|Edition|Restored|((2|3|4)in1)))){1,3}`, regexp2.IgnoreCase)
	languageSimplifyExp = regexp2.MustCompile(`\b(TRUE.?FRENCH|videomann|SUBFRENCH|PLDUB|MULTI)`, regexp2.IgnoreCase)
	sceneGarbageExp     = regexp2.MustCompile(`\b(PROPER|REAL|READ.NFO)`, regexp2.None)

	allLanguagesExp = regexp2.MustCompile(
		`\b(`+strings.Join(allLanguageNames(), `|`)+`)`,
		regexp2.None,
	)
)

func allLanguageNames() []string {
	langs := []Language{
		LanguageEnglish, LanguageFrench, LanguageSpanish, LanguageGerman,
		LanguageItalian, LanguageDanish, LanguageDutch, LanguageJapanese,
		LanguageCantonese, LanguageMandarin, LanguageRussian, LanguagePolish,
		LanguageVietnamese, LanguageNordic, LanguageSwedish, LanguageNorwegian,
		LanguageFinnish, LanguageTurkish, LanguagePortuguese, LanguageFlemish,
		LanguageGreek, LanguageKorean, LanguageHungarian, LanguagePersian,
		LanguageBengali, LanguageBulgarian, LanguageBrazilian, LanguageHebrew,
		LanguageCzech, LanguageUkrainian, LanguageCatalan, LanguageChinese,
		LanguageThai, LanguageHindi, LanguageTamil, LanguageArabic,
		LanguageEstonian, LanguageIcelandic, LanguageLatvian, LanguageLithuanian,
		LanguageRomanian, LanguageSlovak, LanguageSerbian,
	}
	names := make([]string, len(langs))
	for i, l := range langs {
		names[i] = strings.ToUpper(string(l))
	}
	return names
}

// replaceAll replaces all occurrences of a pattern in s with repl.
func replaceAll(re *regexp2.Regexp, s, repl string) string {
	result, _ := re.Replace(s, repl, -1, -1)
	return result
}

// SimplifyTitle strips resolution, codec, and common source tags from a title.
func SimplifyTitle(title string) string {
	s := replaceAll(simpleTitleRegex, title, "")
	s = replaceAll(websitePrefixRegex, s, "")
	s = replaceAll(cleanTorrentPrefixRegex, s, "")
	s = replaceAll(cleanTorrentSuffixRegex, s, "")
	s = replaceAll(commonSourcesGlobalExp, s, "")
	s = replaceAll(webdlExp, s, "")
	s = replaceAll(codecGlobalExp, s, "")
	return strings.TrimSpace(s)
}

// ReleaseTitleCleaner strips edition, language, and scene-garbage tokens from a raw title segment.
// Returns empty string if the title is effectively empty.
func ReleaseTitleCleaner(title string) string {
	if title == "" || title == "(" {
		return ""
	}

	t := strings.ReplaceAll(title, "_", " ")
	t = replaceAll(requestInfoRegex, t, "")
	t = strings.TrimSpace(t)
	t = replaceAll(commonSourcesGlobalExp, t, "")
	t = strings.TrimSpace(t)
	t = replaceAll(webdlExp, t, "")
	t = strings.TrimSpace(t)
	t = replaceAll(editionExp, t, "")
	t = strings.TrimSpace(t)
	t = replaceAll(languageSimplifyExp, t, "")
	t = strings.TrimSpace(t)
	t = replaceAll(sceneGarbageExp, t, "")
	t = strings.TrimSpace(t)
	t = replaceAll(allLanguagesExp, t, "")
	t = strings.TrimSpace(t)

	// Drop everything after a double-space or double-dot gap
	if idx := strings.Index(t, "  "); idx >= 0 {
		t = t[:idx]
	}
	if idx := strings.Index(t, ".."); idx >= 0 {
		t = t[:idx]
	}

	// Reconstruct title from dot-separated parts respecting acronyms
	parts := strings.Split(t, ".")
	var result strings.Builder
	previousAcronym := false
	for n, part := range parts {
		var nextPart string
		if n+1 < len(parts) {
			nextPart = parts[n+1]
		}

		lower := strings.ToLower(part)
		_, notNum := func() (int, bool) {
			var v int
			_, err := func() (int, error) {
				for _, c := range part {
					if c < '0' || c > '9' {
						return 0, nil
					}
				}
				return 0, nil
			}()
			_ = err
			return v, false
		}()
		_ = notNum

		isNum := true
		for _, c := range part {
			if c < '0' || c > '9' {
				isNum = false
				break
			}
		}

		if len(part) == 1 && lower != "a" && !isNum {
			result.WriteString(part)
			result.WriteByte('.')
			previousAcronym = true
		} else if lower == "a" && (previousAcronym || len(nextPart) == 1) {
			result.WriteString(part)
			result.WriteByte('.')
			previousAcronym = true
		} else {
			if previousAcronym {
				result.WriteByte(' ')
				previousAcronym = false
			}
			result.WriteString(part)
			result.WriteByte(' ')
		}
	}

	return strings.TrimSpace(result.String())
}
