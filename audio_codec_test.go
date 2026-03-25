package videoparser

import "testing"

func TestParseAudioCodec(t *testing.T) {
	cases := []struct {
		input     string
		wantCodec AudioCodec
		wantSrc   string
	}{
		{
			"Hannibal 2001 4K UHD Dolby Vision MP4 DD+5 1 H265-d3g",
			AudioCodecDOLBY, "Dolby",
		},
		{
			"Aladdin 2019 720p BluRay x264 AC3 5 1-OMEGA",
			AudioCodecDOLBY, "AC3",
		},
		// no codec found → empty result
		{
			"Trespass Against Us (2017) 1080p BluRay x265 6ch -Dtech mkv",
			"", "",
		},
		{
			"Abbot and Costello Meet Frankenstein 1948 BluRay 1080p HEVC Dts Stereo-D3FiL3R",
			AudioCodecDTS, "Dts",
		},
		{
			"The.Daily.Show.2015.07.01.Kirsten.Gillibrand.Extended.720p.Comedy.Central.WEBRip.AAC2.0.x264-BTW.mkv",
			AudioCodecAAC, "AAC",
		},
		{
			"Girl on the Third Floor 2019 BRRip x264 AAC-SSN",
			AudioCodecAAC, "AAC",
		},
		{
			"New Eden S01E01 Who Are These Women CRAV WEB-DL AAC2 0 H 264-BTW",
			AudioCodecAAC, "AAC",
		},
		{
			"South Park S20E08 Members Only Uncensored 1080p WEB-DL HEVC x265 AAC2ch-NEBO666",
			AudioCodecAAC, "AAC",
		},
		{
			"Behind the Candelabra 2013 BDRip 1080p DTS-HD extra-HighCode",
			AudioCodecDTSHD, "DTS-HD",
		},
		{
			"Ex Machina 2015 UHD BluRay 2160p DTS-X 7 1 HDR x265 10bit-CHD",
			AudioCodecDTSHD, "DTS-X",
		},
		{
			"Frozen.2.2019.German.DL.EAC3.1080p.DSNP.WEB.H265-ZeroTwo",
			AudioCodecEAC3, "EAC3",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := parseAudioCodec(c.input)
			if got.Codec != c.wantCodec {
				t.Errorf("Codec = %q, want %q", got.Codec, c.wantCodec)
			}
			if got.Source != c.wantSrc {
				t.Errorf("Source = %q, want %q", got.Source, c.wantSrc)
			}
		})
	}
}
