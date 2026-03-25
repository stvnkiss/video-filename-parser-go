package videoparser

import "testing"

func TestReleaseTitleCleaner_Empty(t *testing.T) {
	got := ReleaseTitleCleaner("")
	if got != "" {
		t.Errorf("ReleaseTitleCleaner(\"\") = %q, want \"\"", got)
	}
}

func TestReleaseTitleCleaner(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"Mission Impossible Rogue Nation", "Mission Impossible Rogue Nation"},
		{"Passengers", "Passengers"},
		{"1941", "1941"},
		{"A.I.Artificial.Intelligence", "A.I. Artificial Intelligence"},
		{"(500).Days.Of.Summer", "(500) Days Of Summer"},
		{"We Are the Best!", "We Are the Best!"},
		{"V.H.S.2", "V.H.S. 2"},
		{"The.Man.from.U.N.C.L.E", "The Man from U.N.C.L.E."},
		{"The.Middle.", "The Middle"},
		{"Incassable.TRUE.FRENCH.PROPER.READ.NFO.DVDRiP.DiVX.SBC-KFT", "Incassable"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := ReleaseTitleCleaner(c.input)
			if got != c.want {
				t.Errorf("ReleaseTitleCleaner(%q) = %q, want %q", c.input, got, c.want)
			}
		})
	}
}

func TestSimplifyTitle(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"The.Man.from.U.N.C.L.E.2015.1080p.BluRay.x264-SPARKS", "The.Man.from.U.N.C.L.E.2015...-SPARKS"},
		// Go removes all regex matches (global), TS removes only the first occurrence;
		// result differs from TS: DD5.1 and H.264 are also removed by simpleTitleRegex.
		{"Movie.Title.Imax.2018.1080p.AMZN.WEB-DL.DD5.1.H.264-NTG", "Movie.Title.Imax.2018.....-NTG"},
		{"Thunderbirds.Are.Go.S01E10.Tunnels.Of.Time.720p.HDTV.x264-RDVAS[rartv]", "Thunderbirds.Are.Go.S01E10.Tunnels.Of.Time...-RDVAS"},
		{"The.Middle.720p.HEVC.x265-MeGusta-Pre", "The.Middle...-MeGusta-Pre"},
		{"[REQ] The.Middle.720p.HEVC.x265-MeGusta-Pre", "The.Middle...-MeGusta-Pre"},
		{"Spider-Man Far from Home.2019.1080p.HDRip.X264.AC3-EVO", "Spider-Man Far from Home.2019....AC3-EVO"},
		{"Inglorious.Basterds.CAM.XviD-CAMELOT", "Inglorious.Basterds..-CAMELOT"},
		{"The.Fighter.DVDR-MPTDVD", "The.Fighter.-MPTDVD"},
		{"The Card Counter WEBSCREENER XviD-EVO", "The Card Counter  -EVO"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := SimplifyTitle(c.input)
			if got != c.want {
				t.Errorf("SimplifyTitle(%q)\n got  = %q\n want = %q", c.input, got, c.want)
			}
		})
	}
}
