package videoparser

import "testing"

func TestParseAudioChannels(t *testing.T) {
	cases := []struct {
		input   string
		wantCh  Channels
		wantSrc string
	}{
		{
			"Hannibal 2001 4K UHD Dolby Vision MP4 DD+5 1 H265-d3g",
			ChannelsSIX, "5 1",
		},
		{
			"Aladdin 2019 720p BluRay x264 AC3 5 1-OMEGA",
			ChannelsSIX, "5 1",
		},
		{
			"Trespass Against Us (2017) 1080p BluRay x265 6ch -Dtech mkv",
			ChannelsSIX, "6ch",
		},
		{
			"Abbot and Costello Meet Frankenstein 1948 BluRay 1080p HEVC Dts Stereo-D3FiL3R",
			ChannelsSTEREO, "Stereo",
		},
		{
			"The.Daily.Show.2015.07.01.Kirsten.Gillibrand.Extended.720p.Comedy.Central.WEBRip.AAC2.0.x264-BTW.mkv",
			ChannelsSTEREO, "2.0",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := parseAudioChannels(c.input)
			if got.Channels != c.wantCh {
				t.Errorf("Channels = %q, want %q", got.Channels, c.wantCh)
			}
			if got.Source != c.wantSrc {
				t.Errorf("Source = %q, want %q", got.Source, c.wantSrc)
			}
		})
	}
}
