package videoparser

import "github.com/dlclark/regexp2"

var resolutionExp = regexp2.MustCompile(
	`(?<R2160P>2160p|4k[-_. ](?:UHD|HEVC|BD)|(?:UHD|HEVC|BD)[-_. ]4k|\b(4k)\b|COMPLETE.UHD|UHD.COMPLETE)`+
		`|(?<R1080P>1080(i|p)|1920x1080)(10bit)?`+
		`|(?<R720P>720(i|p)|1280x720|960p)(10bit)?`+
		`|(?<R576P>576(i|p))`+
		`|(?<R540P>540(i|p))`+
		`|(?<R480P>480(i|p)|640x480|848x480)`,
	regexp2.IgnoreCase,
)

// resolutionResult bundles a matched Resolution with the raw matched text.
type resolutionResult struct {
	Resolution Resolution
	Source     string
}

// parseResolution extracts video resolution from a title string.
// An optional pre-computed source list may be provided to skip an extra source parse.
func parseResolution(title string, precomputedSources ...[]Source) resolutionResult {
	m, err := resolutionExp.FindStringMatch(title)
	if err == nil && m != nil {
		for _, name := range []string{"R2160P", "R1080P", "R720P", "R576P", "R540P", "R480P"} {
			g := m.GroupByName(name)
			if g != nil && g.Length > 0 {
				res := map[string]Resolution{
					"R2160P": R2160P,
					"R1080P": R1080P,
					"R720P":  R720P,
					"R576P":  R576P,
					"R540P":  R540P,
					"R480P":  R480P,
				}[name]
				return resolutionResult{Resolution: res, Source: g.String()}
			}
		}
	}

	// Fallback: guess from source
	var sources []Source
	if len(precomputedSources) > 0 {
		sources = precomputedSources[0]
	} else {
		sources = parseSource(title)
	}
	for _, s := range sources {
		if s == SourceDVD {
			return resolutionResult{Resolution: R480P}
		}
	}
	return resolutionResult{}
}
