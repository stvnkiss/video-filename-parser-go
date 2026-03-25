package videoparser

import (
	"reflect"
	"testing"
)

func TestParseSource_Single(t *testing.T) {
	cases := []struct {
		input string
		want  []Source
	}{
		{"Whats.Eating.Gilbert.Grape.1993.720p.BluRay", []Source{SourceBLURAY}},
		{"Whats.Eating.Gilbert.Grape.1993.720p.BluRay.x264-SiNNERS", []Source{SourceBLURAY}},
		{"Oceans.Thirteen.2007.iNTERNAL.720p.BluRay.x264-MHQ", []Source{SourceBLURAY}},
		{"Rocketman 2019 2160p UHD BluRay x265-TERMiNAL", []Source{SourceBLURAY}},
		{"Alita Battle Angel 2019 2160p WEB-DL DD+5 1 HEVC-DEFLATE[NO RAR]", []Source{SourceWEBDL}},
		{"Alita: Battle Angel 2019 BRRip AC3 x264-CMRG", []Source{SourceBLURAY}},
		{"The Hateful Eight 2015 DVDScr XVID AC3 HQ Hive-CM8", []Source{SourceSCREENER}},
		{"This Is 40 2012 DVD Screener Xvid UnKnOwN", []Source{SourceSCREENER}},
		{"Brooklyns Finest DVDSCREENER XviD-MENTiON", []Source{SourceSCREENER}},
		{"50 50 2011 SCREENER XviD-REFiLL", []Source{SourceSCREENER}},
		{"True Grit 2010 SCR XViD - IMAGiNE", []Source{SourceSCREENER}},
		{"Tracers 2015 PPV XVID AC3 HQ Hive-CM8", []Source{SourcePPV}},
		{"X-Men Origins Wolverine 2009 WORKPRINT XviD-NoGRP", []Source{SourceWORKPRINT}},
		{"Half Baked 1998 720p HDDVD XVID-AC3-PULSAR", []Source{SourceBLURAY}},
		{"Teenage Mutant Ninja Turtles Turtles Forever 2009 WS PDTV XviD-DVSKY", []Source{SourceTV}},
		{"The Interview 2014 1080p WEB-DL x264 AAC MerryXmas", []Source{SourceWEBDL}},
		{"John Wick Chapter 2 2017 720p WEB-DL X264 AC3-EVO", []Source{SourceWEBDL}},
		{"Into the Storm 720 WEBDL (RUSSiAN & ENGLISH AUDIO)", []Source{SourceWEBDL}},
		{"Michael 1996 1080p AMZN WEBCap DD+5.1 x264-LEGi0N", []Source{SourceWEBRIP}},
		{"Avengers Infinity War 2018 NEW PROPER 720p HD-CAM X264 HQ-LPG", []Source{SourceCAM}},
		{"The Hunger Games Mockingjay - Part 1 (2014) 576p NEW CAM XViD", []Source{SourceCAM}},
		{"Suicide Squad 2016 CAM UnKnOwN", []Source{SourceCAM}},
		{"Star Trek Beyond (2016) ENG Cam V2 XviD UnKnOwN", []Source{SourceCAM}},
		{"Parasite.2019.MULTi.VFI.WEBrip.2160p.HDR.x265.True.HD-Tokuchi", []Source{SourceWEBRIP}},
		{"How You Look At Me 2019 720p AMZN WEBRip AAC2 0 X 264-EVO", []Source{SourceWEBRIP}},
		{"Togo 2019 2160p HDR DSNP WEBRip DDPAtmos 5 1 X265-TrollUHD", []Source{SourceWEBRIP}},
		{"Palmer.2021.1080p.APTV.H264.Atmos-EVO", []Source{SourceWEBDL}},
		{"Palmer.2021.1080p.APTV.WEB-RIP.H264.Atmos-EVO", []Source{SourceWEBRIP}},
		{"Finding.Ohana.2021.720p.NF.AAC2.0.X.264-EVO", []Source{SourceWEBDL}},
		{"Finding.Ohana.2021.720p.NF.WEBRIP.AAC2.0.X.264-EVO", []Source{SourceWEBRIP}},
		{"300.2006.iNTERNAL.NTSC.DVD9-FaiLED", []Source{SourceDVD}},
		{"The Card Counter 2021 1080p WEBSCREENER X264-EVO", []Source{SourceSCREENER}},
		{"Movie.Name.2016.German.DTS.DL.1080p.UHDBD.x265-TDO", []Source{SourceBLURAY}},
		{"127.Hours.DVDSCR.NTSC.DVDR-GALAXY", []Source{SourceDVD, SourceSCREENER}},
		{"Movie.Title.2019.1080p.AMZN.WEB-Rip.DDP.5.1.HEVC", []Source{SourceWEBRIP}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := parseSource(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("parseSource(%q)\n got  = %v\n want = %v", c.input, got, c.want)
			}
		})
	}
}

func TestParseSource_Multi(t *testing.T) {
	cases := []struct {
		input string
		want  []Source
	}{
		{"The Office S01-S09 720p BluRay WEB-DL nHD x264-NhaNc3", []Source{SourceBLURAY, SourceWEBDL}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := parseSource(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("parseSource(%q)\n got  = %v\n want = %v", c.input, got, c.want)
			}
		})
	}
}
