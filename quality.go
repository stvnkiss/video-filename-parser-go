package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var (
	properRegex    = regexp2.MustCompile(`\b(?<proper>proper|repack|rerip)\b`, regexp2.IgnoreCase)
	realRegex      = regexp2.MustCompile(`\b(?<real>REAL)\b`, regexp2.None) // case-sensitive
	versionExp     = regexp2.MustCompile(`(?<version>v\d\b|\[v\d\])`, regexp2.IgnoreCase)
	remuxExp       = regexp2.MustCompile(`\b(?<remux>(BD|UHD)?Remux)\b`, regexp2.IgnoreCase)
	bdiskExp       = regexp2.MustCompile(`\b(COMPLETE|ISO|BDISO|BDMux|BD25|BD50|BR.?DISK)\b`, regexp2.IgnoreCase)
	rawHdExp       = regexp2.MustCompile(`\b(?<rawhd>RawHD|1080i[-_. ]HDTV|Raw[-_. ]HD|MPEG[-_. ]?2)\b`, regexp2.IgnoreCase)
	highDefPdtvExp = regexp2.MustCompile(`hr[-_. ]ws`, regexp2.IgnoreCase)
)

// ParseQualityModifiers extracts proper/real/version revision information.
func ParseQualityModifiers(title string) Revision {
	normalized := strings.TrimSpace(strings.ReplaceAll(strings.TrimSpace(title), "_", " "))
	lower := strings.ToLower(normalized)

	result := Revision{Version: 1, Real: 0}

	if mustMatch(properRegex, lower) {
		result.Version = 2
	}

	vm, _ := versionExp.FindStringMatch(lower)
	if vm != nil {
		g := vm.GroupByName("version")
		if g != nil && g.Length > 0 {
			s := g.String()
			// Extract the digit
			for _, c := range s {
				if c >= '0' && c <= '9' {
					result.Version = int(c - '0')
					break
				}
			}
		}
	}

	// Count REAL occurrences (case-sensitive, not in REALLY etc.)
	realCount := 0
	rm, _ := realRegex.FindStringMatch(title)
	for rm != nil {
		realCount++
		rm, _ = realRegex.FindNextMatch(rm)
	}
	result.Real = realCount

	return result
}

// ParseQuality returns a QualityModel for the given title.
// An optional pre-parsed VideoCodec may be supplied to avoid redundant parsing.
func ParseQuality(title string, codec ...VideoCodec) QualityModel {
	normalized := strings.ToLower(strings.TrimSpace(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(strings.TrimSpace(title), "_", " "),
				"[", " "),
			"]", " "),
	))

	revision := ParseQualityModifiers(title)
	sg := parseSourceGroups(normalized)
	sources := parseSource(normalized, sg)
	resResult := parseResolution(normalized, sources)
	resolution := resResult.Resolution

	var vc VideoCodec
	if len(codec) > 0 {
		vc = codec[0]
	} else {
		vc = parseVideoCodec(title).Codec
	}

	result := QualityModel{
		Sources:    sources,
		Resolution: resolution,
		Revision:   revision,
	}

	if mustMatch(bdiskExp, normalized) && sg.bluray {
		result.Modifier = QualityModifierBRDISK
		result.Sources = []Source{SourceBLURAY}
	}

	if mustMatch(remuxExp, normalized) && !sg.webdl && !sg.hdtv {
		result.Modifier = QualityModifierREMUX
		result.Sources = []Source{SourceBLURAY}
	}

	if mustMatch(rawHdExp, normalized) &&
		result.Modifier != QualityModifierBRDISK &&
		result.Modifier != QualityModifierREMUX {
		result.Modifier = QualityModifierRAWHD
		result.Sources = []Source{SourceTV}
	}

	if len(sources) > 0 {
		if sg.bluray {
			result.Sources = []Source{SourceBLURAY}
			if vc == VideoCodecXVID {
				result.Resolution = R480P
				result.Sources = []Source{SourceDVD}
				return result
			}
			if resolution == "" {
				result.Resolution = R720P
			}
			if resolution == "" && result.Modifier == QualityModifierBRDISK {
				result.Resolution = R1080P
			}
			if resolution == "" && result.Modifier == QualityModifierREMUX {
				result.Resolution = R2160P
			}
			return result
		}

		if sg.webdl || sg.webrip {
			result.Sources = sources
			if resolution == "" {
				result.Resolution = R480P
			}
			if resolution == "" && strings.Contains(title, "[WEBDL]") {
				result.Resolution = R720P
			}
			return result
		}

		if sg.hdtv {
			result.Sources = []Source{SourceTV}
			if resolution == "" {
				result.Resolution = R480P
			}
			if resolution == "" && strings.Contains(title, "[HDTV]") {
				result.Resolution = R720P
			}
			return result
		}

		if sg.pdtv || sg.sdtv || sg.dsr || sg.tvrip {
			result.Sources = []Source{SourceTV}
			if mustMatch(highDefPdtvExp, normalized) {
				result.Resolution = R720P
				return result
			}
			result.Resolution = R480P
			return result
		}

		if sg.bdrip || sg.brrip {
			if vc == VideoCodecXVID {
				result.Resolution = R480P
				result.Sources = []Source{SourceDVD}
				return result
			}
			if resolution == "" {
				result.Resolution = R480P
			}
			result.Sources = []Source{SourceBLURAY}
			return result
		}

		if sg.workprint {
			result.Sources = []Source{SourceWORKPRINT}
			return result
		}
		if sg.cam {
			result.Sources = []Source{SourceCAM}
			return result
		}
		if sg.ts {
			result.Sources = []Source{SourceTELESYNC}
			return result
		}
		if sg.tc {
			result.Sources = []Source{SourceTELECINE}
			return result
		}
	}

	if result.Modifier == "" &&
		(resolution == R2160P || resolution == R1080P || resolution == R720P) {
		result.Sources = []Source{SourceWEBDL}
		return result
	}

	return result
}
