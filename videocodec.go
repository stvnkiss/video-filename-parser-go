package videoparser

import "github.com/dlclark/regexp2"

var codecExp = regexp2.MustCompile(
	`(?<x265>x265)`+
		`|(?<h265>h265)`+
		`|(?<x264>x264)`+
		`|(?<h264>h264)`+
		`|(?<wmv>WMV)`+
		`|(?<xvidhd>XvidHD)`+
		`|(?<xvid>X-?vid)`+
		`|(?<divx>divx)`+
		`|(?<hevc>HEVC)`+
		`|(?<dvdr>DVDR)\b`,
	regexp2.IgnoreCase,
)

type videoCodecResult struct {
	Codec  VideoCodec
	Source string
}

func parseVideoCodec(title string) videoCodecResult {
	m, err := codecExp.FindStringMatch(title)
	if err != nil || m == nil {
		return videoCodecResult{}
	}

	groupVal := func(name string) string {
		g := m.GroupByName(name)
		if g != nil && g.Length > 0 {
			return g.String()
		}
		return ""
	}

	if v := groupVal("h264"); v != "" {
		return videoCodecResult{Codec: VideoCodecH264, Source: v}
	}
	if v := groupVal("h265"); v != "" {
		return videoCodecResult{Codec: VideoCodecH265, Source: v}
	}
	if v := groupVal("x265"); v != "" {
		return videoCodecResult{Codec: VideoCodecX265, Source: v}
	}
	if v := groupVal("hevc"); v != "" {
		return videoCodecResult{Codec: VideoCodecX265, Source: v}
	}
	if v := groupVal("x264"); v != "" {
		return videoCodecResult{Codec: VideoCodecX264, Source: v}
	}
	if v := groupVal("xvidhd"); v != "" {
		return videoCodecResult{Codec: VideoCodecXVID, Source: v}
	}
	if v := groupVal("xvid"); v != "" {
		return videoCodecResult{Codec: VideoCodecXVID, Source: v}
	}
	if v := groupVal("divx"); v != "" {
		return videoCodecResult{Codec: VideoCodecXVID, Source: v}
	}
	if v := groupVal("wmv"); v != "" {
		return videoCodecResult{Codec: VideoCodecWMV, Source: v}
	}
	if v := groupVal("dvdr"); v != "" {
		return videoCodecResult{Codec: VideoCodecDVDR, Source: v}
	}
	return videoCodecResult{}
}
