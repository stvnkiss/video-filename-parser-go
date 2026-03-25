package videoparser

import "testing"

func TestParseEdition(t *testing.T) {
	cases := []struct {
		input string
		want  Edition
	}{
		{"Prometheus 2012 Directors Cut", Edition{Directors: true}},
		{"Star Wars Episode IV - A New Hope 1999 (Despecialized).mkv", Edition{FanEdit: true}},
		{"Prometheus.2012.(Special.Edition.Remastered).[Bluray-1080p].mkv", Edition{Remastered: true}},
		{"Prometheus 2012 Extended", Edition{Extended: true}},
		{"Prometheus 2012 The Extended Cut", Edition{Extended: true}},
		{"Prometheus 2012 Extended Directors Cut Fan Edit", Edition{Directors: true, FanEdit: true, Extended: true}},
		{"Prometheus.2012.(Extended.Theatrical.Version.IMAX).BluRay.1080p.2012.asdf", Edition{IMAX: true, Theatrical: true, Extended: true}},
		{"2001: A Space Odyssey 1968 (Extended Directors Cut FanEdit)", Edition{FanEdit: true, Directors: true, Extended: true}},
		{"A Fake Movie 2035 2012 Directors.mkv", Edition{Directors: true}},
		{"Blade Runner 2049 Director's Cut.mkv", Edition{Directors: true}},
		{"Prometheus 2012 50th Anniversary Edition.mkv", Edition{Remastered: true}},
		{"Movie 2012 IMAX.mkv", Edition{IMAX: true}},
		{"Movie 2012 Restored.mkv", Edition{Remastered: true}},
		{"Prometheus.Special.Edition.Fan Edit.2012..BRRip.x264.AAC-m2g", Edition{FanEdit: true}},
		{"Star Wars Episode IV - A New Hope (Despecialized) 1999.mkv", Edition{FanEdit: true}},
		{"Prometheus.(Special.Edition.Remastered).2012.[Bluray-1080p].mkv", Edition{Remastered: true}},
		{"Prometheus Extended 2012", Edition{Extended: true}},
		{"Prometheus Extended Directors Cut Fan Edit 2012", Edition{Extended: true, Directors: true, FanEdit: true}},
		{"Prometheus.(Extended.Theatrical.Version.IMAX).2012.BluRay.1080p.asdf", Edition{Extended: true, Theatrical: true, IMAX: true}},
		{"2001: A Space Odyssey (Extended Directors Cut FanEdit) 1968 Bluray 1080p", Edition{Extended: true, FanEdit: true, Directors: true}},
		{"Prometheus 50th Anniversary Edition 2012.mkv", Edition{Remastered: true}},
		{"X-Men Days of Future Past 2014 THE ROGUE CUT BRRip XviD AC3-EVO", Edition{Extended: true}},
		{"Alita Battle Angel 2019 INTERNAL HDR 2160p WEB H265-DEFLATE", Edition{Internal: true, HDR: true}},
		{"Wonder.Woman.1984.2020.IMAX.3D.1080p.BluRay.Half-SBS.DTS-HD.MA.5.1.X264-EVO", Edition{IMAX: true, ThreeD: true, HSBS: true}},
		{"Warcraft.The.Beginning.3D.HOU.2016.German.DL.1080p.BluRay.x264-COiNCiDENCE", Edition{ThreeD: true, HOU: true}},
		{"Iron.Man.2008.INTERNAL.REMASTERED.2160p.UHD.BluRay.X265-IAMABLE", Edition{UHD: true, Remastered: true, Internal: true}},
		{"Long Shot 2019 DV 2160p WEB H265-SLOT", Edition{DolbyVision: true}},
		{"Sicario 2015 Hybrid 2160p UHD BluRay REMUX DV HDR10+ HEVC TrueHD 7.1 Atmos-WiLDCAT", Edition{UHD: true, DolbyVision: true}},
		{"Babylon.2022.OAR.1080p.WEB.H264-SLOT", Edition{OAR: true}},
		{"Solo.A.Star.Wars.Story.2018.BONUS.DELETED.SCENES.1080p.BluRay.x264-PussyFoot", Edition{DeletedScenes: true}},
		{"The.Golden.Compass.2007.BONUS.1080p.BluRay.H264-REFRACTiON", Edition{BonusContent: true}},
		{"The.Mist.2007.BW.2160p.UHD.BluRay.x265-GUHZER", Edition{BW: true, UHD: true}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := ParseEdition(c.input)
			if got != c.want {
				t.Errorf("ParseEdition(%q)\n got  = %+v\n want = %+v", c.input, got, c.want)
			}
		})
	}
}

func TestParseEditionHardcodedSubs(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"Movie.Title.2016.1080p.KORSUB.WEBRip.x264.AAC2.0-RADARR", true},
		{"Movie.Title.2016.1080p.KORSUBS.WEBRip.x264.AAC2.0-RADARR", true},
		{"Movie Title 2017 HC 720p HDRiP DD5 1 x264-LEGi0N", true},
		{"Movie.Title.2017.720p.SUBBED.HDRip.V2.XViD-26k.avi", true},
		{"Movie.Title.2000.1080p.BlueRay.x264.DTS.RoSubbed-playHD", false},
		{"Movie Title! 2018 [Web][MKV][h264][480p][AAC 2.0][Softsubs]", false},
		{"Movie Title! 2019 [HorribleSubs][Web][MKV][h264][848x480][AAC 2.0][Softsubs(HorribleSubs)]", false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := ParseEdition(c.input).HardcodedSubs
			if got != c.want {
				t.Errorf("HardcodedSubs(%q) = %v, want %v", c.input, got, c.want)
			}
		})
	}
}
