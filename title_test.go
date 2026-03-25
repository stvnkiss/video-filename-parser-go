package videoparser

import "testing"

func TestParseTitleAndYear(t *testing.T) {
	cases := []struct {
		input     string
		wantTitle string
		wantYear  string
	}{
		{"The.Man.from.U.N.C.L.E.2015.1080p.BluRay.x264-SPARKS", "The Man from U.N.C.L.E.", "2015"},
		{"1941.1979.EXTENDED.720p.BluRay.X264-AMIABLE", "1941", "1979"},
		{"MY MOVIE (2016) [R][Action, Horror][720p.WEB-DL.AVC.8Bit.6ch.AC3].mkv", "MY MOVIE", "2016"},
		{"R.I.P.D.2013.720p.BluRay.x264-SPARKS", "R.I.P.D.", "2013"},
		{"V.H.S.2.2013.LIMITED.720p.BluRay.x264-GECKOS", "V.H.S. 2", "2013"},
		{
			"This Is A Movie (1999) [IMDB #] <Genre, Genre, Genre> {ACTORS} !DIRECTOR +MORE_SILLY_STUFF_NO_ONE_NEEDS ?",
			"This Is A Movie", "1999",
		},
		{"We Are the Best!.2013.720p.H264.mkv", "We Are the Best!", "2013"},
		{"(500).Days.Of.Summer.(2009).DTS.1080p.BluRay.x264.NLsubs", "(500) Days Of Summer", "2009"},
		{"To.Live.and.Die.in.L.A.1985.1080p.BluRay", "To Live and Die in L.A.", "1985"},
		{"A.I.Artificial.Intelligence.(2001)", "A.I. Artificial Intelligence", "2001"},
		{"A.Movie.Name.(1998)", "A Movie Name", "1998"},
		{"www.Torrenting.com - Revenge.2008.720p.X264-DIMENSION", "Revenge", "2008"},
		{"Thor: The Dark World 2013", "Thor The Dark World", "2013"},
		{"Resident.Evil.The.Final.Chapter.2016", "Resident Evil The Final Chapter", "2016"},
		{
			"Mission Impossible: Rogue Nation (2015)\xef\xbf\xbd[XviD - Ita Ac3 - SoftSub Ita]azione, spionaggio, thriller *Prima Visione* Team mulnic Tom Cruise",
			"Mission Impossible Rogue Nation", "2015",
		},
		{"Scary.Movie.2000.FRENCH..BluRay.-AiRLiNE", "Scary Movie", "2000"},
		{"My Movie 1999 German Bluray", "My Movie", "1999"},
		{"Leaving Jeruselem by Railway (1897) [DVD].mp4", "Leaving Jeruselem by Railway", "1897"},
		{"Climax.2018.1080p.AMZN.WEB-DL.DD5.1.H.264-NTG", "Climax", "2018"},
		{"Movie.Title.Imax.2018.1080p.AMZN.WEB-DL.DD5.1.H.264-NTG", "Movie Title", "2018"},
		{"The.Middle.720p.HEVC.x265-MeGusta-Pre", "The Middle", ""},
		{"The.Middle.HEVC.x265-MeGusta-Pre", "The Middle", ""},
		{"Blade Runner 2049 2017", "Blade Runner 2049", "2017"},
		{"Blade Runner 2049 (2017)", "Blade Runner 2049", "2017"},
		{"Scarface.Anniversary.Edition.1983.INTERNAL.DVDRip.XviD-VoMiT", "Scarface", "1983"},
		{"Ouija.Origin.of.Evil.2016.MULTi.TRUEFRENCH.1080p.BluRay.x264-MELBA", "Ouija Origin of Evil", "2016"},
		{"Appaloosa.1080p.Bluray.x264-1920", "Appaloosa", ""},
		{"Inglorious.Basterds.CAM.XviD-CAMELOT", "Inglorious Basterds", ""},
		{"Inglourious.Basterds.SCR.XViD-xSCR", "Inglourious Basterds", ""},
		{"No.Country.for.Old.Men.DVDRip.XviD-DiAMOND", "No Country for Old Men", ""},
		{"The.Fighter.DVDR-MPTDVD", "The Fighter", ""},
		{"Sunshine.Cleaning.DVDR-Replica", "Sunshine Cleaning", ""},
		{"Scarface.The.Uncut.Version.1983.DVDRip.Divx.AC3.iNTERNAL-FFM", "Scarface", "1983"},
		{"Casino.10TH.ANNiVERSARY.1995.iNTERNAL.DVDRiP.XViD-KiSS", "Casino", "1995"},
		{"Get.Him.To.The.Greek.UNRATED.FRENCH.720p.BluRay.x264-NERDHD", "Get Him To The Greek", ""},
		{
			"The.Social.Network.German.720p.BluRay.x264-DECENT",
			"The Social Network German", "",
		},
		{"The.Outsiders.DC.German.1983.AC3.BDRip.XviD.INTERNAL-ARC", "The Outsiders", "1983"},
		{"The.Girl.in.the.Spiders.Web.2018.2160p.UHD.BluRay.x265-VALiS", "The Girl in the Spiders Web", "2018"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := ParseTitleAndYear(c.input)
			if got.Title != c.wantTitle {
				t.Errorf("Title = %q, want %q", got.Title, c.wantTitle)
			}
			if got.Year != c.wantYear {
				t.Errorf("Year = %q, want %q", got.Year, c.wantYear)
			}
		})
	}
}
