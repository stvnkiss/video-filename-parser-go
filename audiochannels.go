package videoparser

import "github.com/dlclark/regexp2"

var channelExp = regexp2.MustCompile(
	`\b(?<eight>7.?[01])\b`+
		`|\b(?<six>(6[\W]0(?:ch)?)(?=[^\d]|$)|(5[\W][01](?:ch)?)(?=[^\d]|$)|5ch|6ch)\b`+
		`|(?<stereo>((2[\W]0(?:ch)?)(?=[^\d]|$))|(stereo))`+
		`|(?<mono>(1[\W]0(?:ch)?)(?=[^\\d]|$)|(mono)|(1ch))`,
	regexp2.IgnoreCase,
)

type audioChannelsResult struct {
	Channels Channels
	Source   string
}

func parseAudioChannels(title string) audioChannelsResult {
	m, err := channelExp.FindStringMatch(title)
	if err != nil || m == nil {
		return audioChannelsResult{}
	}

	groupVal := func(name string) string {
		g := m.GroupByName(name)
		if g != nil && g.Length > 0 {
			return g.String()
		}
		return ""
	}

	if v := groupVal("eight"); v != "" {
		return audioChannelsResult{Channels: ChannelsSEVEN, Source: v}
	}
	if v := groupVal("six"); v != "" {
		return audioChannelsResult{Channels: ChannelsSIX, Source: v}
	}
	if v := groupVal("stereo"); v != "" {
		return audioChannelsResult{Channels: ChannelsSTEREO, Source: v}
	}
	if v := groupVal("mono"); v != "" {
		return audioChannelsResult{Channels: ChannelsMONO, Source: v}
	}
	return audioChannelsResult{}
}
