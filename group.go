package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var (
	websitePrefixGroupExp = regexp2.MustCompile(
		`^\[\s*[a-z]+(\.[a-z]+)+\s*\][- ]*|^www\.[a-z]+\.(?:com|net)[ -]*`,
		regexp2.IgnoreCase,
	)
	cleanReleaseGroupExp = regexp2.MustCompile(
		`(-(RP|1|NZBGeek|Obfuscated|Obfuscation|Scrambled|sample|Pre|postbot|xpost|Rakuv[a-z0-9]*|WhiteRev|BUYMORE|AsRequested|AlternativeToRequested|GEROV|Z0iDS3N|Chamele0n|4P|4Planet|AlteZachen|RePACKPOST))+$`,
		regexp2.IgnoreCase,
	)
	releaseGroupRegexExp = regexp2.MustCompile(
		`-(?<releasegroup>[a-z0-9]+)(?<!WEB-DL|WEB-RIP|480p|720p|1080p|2160p|DTS-(HD|X|MA|ES)|([a-zA-Z]{3}-ENG))(?:\b|[-._ ])`,
		regexp2.IgnoreCase,
	)
	animeReleaseGroupExp = regexp2.MustCompile(
		`^(?:\[(?<subgroup>(?!\s).+?(?<!\s))\](?:_|-|\s|\.)?)`,
		regexp2.IgnoreCase,
	)
	exceptionReleaseGroupRegex = regexp2.MustCompile(
		`(\[)?(?<releasegroup>(Joy|YIFY|YTS.(MX|LT|AG)|FreetheFish|VH-PROD|FTW-HS|DX-TV|Blu-bits|afm72|Anna|Bandi|Ghost|Kappa|MONOLITH|Qman|RZeroX|SAMPA|Silence|theincognito|D-Z0N3|t3nzin|Vyndros|HDO|DusIctv|DHD|SEV|CtrlHD|-ZR-|ADC|XZVN|RH|Kametsu|r00t|HONE))(\])?$`,
		regexp2.IgnoreCase,
	)
)

// ParseGroup extracts the release group from a title string.
func ParseGroup(title string, parsedTitle ...string) string {
	nowebsite, _ := websitePrefixGroupExp.Replace(title, "", -1, -1)

	var releaseTitle string
	if len(parsedTitle) > 0 && parsedTitle[0] != "" {
		releaseTitle = parsedTitle[0]
	} else {
		releaseTitle = ParseTitleAndYear(nowebsite).Title
	}
	releaseTitle = strings.ReplaceAll(releaseTitle, " ", ".")

	trimmed := strings.ReplaceAll(nowebsite, " ", ".")
	if releaseTitle != nowebsite {
		trimmed = strings.ReplaceAll(trimmed, releaseTitle, "")
	}
	trimmed, _ = regexp2.MustCompile(`\.-\.`, regexp2.None).Replace(trimmed, ".", -1, -1)
	trimmed = SimplifyTitle(RemoveFileExtension(strings.TrimSpace(trimmed)))

	if trimmed == "" {
		return ""
	}

	// Check known exception groups
	em, _ := exceptionReleaseGroupRegex.FindStringMatch(trimmed)
	if em != nil {
		g := em.GroupByName("releasegroup")
		if g != nil && g.Length > 0 {
			return g.String()
		}
	}

	// Check for anime subgroup prefix
	am, _ := animeReleaseGroupExp.FindStringMatch(trimmed)
	if am != nil {
		g := am.GroupByName("subgroup")
		if g != nil && g.Length > 0 {
			return g.String()
		}
	}

	trimmed, _ = cleanReleaseGroupExp.Replace(trimmed, "", -1, -1)

	// Walk all matches of releaseGroupRegexExp and return first hit
	m, _ := releaseGroupRegexExp.FindStringMatch(trimmed)
	for m != nil {
		g := m.GroupByName("releasegroup")
		if g != nil && g.Length > 0 {
			return g.String()
		}
		m, _ = releaseGroupRegexExp.FindNextMatch(m)
	}

	return ""
}
