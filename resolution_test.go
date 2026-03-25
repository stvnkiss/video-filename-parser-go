package videoparser

import "testing"

func TestParseResolution(t *testing.T) {
	cases := []struct {
		input string
		want  Resolution
	}{
		{"Oceans.Thirteen.2007.iNTERNAL.720p.BluRay.x264-MHQ", R720P},
		{"Rocketman 2019 2160p UHD BluRay x265-TERMiNAL", R2160P},
		{"Alita Battle Angel 2019 1080p BluRay x264-SPARKS", R1080P},
		{"Alita Battle Angel 2019 HDRip AC3 x264-CMRG", ""},
		{"Alita Battle Angel 2019 2160p WEB-DL DD+5 1 HEVC-DEFLATE[NO RAR]", R2160P},
		{"Alita: Battle Angel 2019 BRRip AC3 x264-CMRG", ""},
		{"Revolution.S01E02.Chained.Heat.[Bluray720p].mkv", R720P},
		{"WEEDS.S03E01-06.DUAL.720p.Blu-ray.AC3.-HELLYWOOD.avi", R720P},
		{"Revolution.S01E02.Chained.Heat.[Bluray1080p].mkv", R1080P},
		{"27.Dresses.2008.REMUX.2160p.Bluray.AVC.DTS-HR.MA.5.1-LEGi0N", R2160P},
		{"Deadpool 2016 2160p 4K UltraHD BluRay DTS-HD MA 7 1 x264-Whatevs", R2160P},
		{"Deadpool 2016 4K 2160p UltraHD BluRay AAC2 0 HEVC x265", R2160P},
		{"The Martian 2015 2160p Ultra HD BluRay DTS-HD MA 7 1 x264-Whatevs", R2160P},
		{"The Revenant 2015 2160p UHD BluRay FLAC 7 1 x264-Whatevs", R2160P},
		{"Into the Inferno 2016 2160p Netflix WEBRip DD5 1 x264-Whatevs", R2160P},
		{"Indiana.Jones.and.the.Temple.of.Doom.1984.Complete.UHD.Bluray-JONES", R2160P},
		{"Orphan Black S05E09 WEBRip 1080p10bit DD5 1 x265 HEVC D0ct0rLew", R1080P},
		{"[SubsPlease] Movie Title (540p) [AB649D32]", R540P},
		{"Series.Title.S04E13.960p.WEB-DL.AAC2.0.H.264-squalor", R720P},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := parseResolution(c.input).Resolution
			if got != c.want {
				t.Errorf("parseResolution(%q) = %q, want %q", c.input, got, c.want)
			}
		})
	}
}

func TestParseResolution_AssumedFromSource(t *testing.T) {
	cases := []struct {
		input string
		want  Resolution
	}{
		{"127.Hours.DVDSCR.NTSC.DVDR-GALAXY", R480P},
		{"127.Hours.GERMAN.2010.DL.PAL.DVDR-OldsMan", R480P},
		{"12.Angry.Men.1957.DvDivX-SMB", R480P},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := parseResolution(c.input).Resolution
			if got != c.want {
				t.Errorf("parseResolution(%q) = %q, want %q", c.input, got, c.want)
			}
		})
	}
}
