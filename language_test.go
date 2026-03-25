package videoparser

import (
	"reflect"
	"testing"
)

func TestParseLanguage(t *testing.T) {
	cases := []struct {
		input string
		want  []Language
	}{
		{"Castle.2009.S01E14.English.HDTV.XviD-LOL", []Language{LanguageEnglish}},
		{"Castle.2009.S01E14.French.HDTV.XviD-LOL", []Language{LanguageFrench}},
		{"Ouija.Origin.of.Evil.2016.MULTi.TRUEFRENCH.1080p.BluRay.x264-MELBA", []Language{LanguageFrench, LanguageEnglish}},
		{"Everest.2015.FRENCH.VFQ.BDRiP.x264-CNF30", []Language{LanguageFrench}},
		{"Showdown.In.Little.Tokyo.1991.MULTI.VFQ.VFF.DTSHD-MASTER.1080p.BluRay.x264-ZombiE", []Language{LanguageFrench, LanguageEnglish}},
		{"The.Polar.Express.2004.MULTI.VF2.1080p.BluRay.x264-PopHD", []Language{LanguageFrench, LanguageEnglish}},
		{"Castle.2009.S01E14.Spanish.HDTV.XviD-LOL", []Language{LanguageSpanish}},
		{"Castle.2009.S01E14.German.HDTV.XviD-LOL", []Language{LanguageGerman}},
		{"Castle.2009.S01E14.Italian.HDTV.XviD-LOL", []Language{LanguageItalian}},
		{"Castle.2009.S01E14.Danish.HDTV.XviD-LOL", []Language{LanguageDanish}},
		{"Castle.2009.S01E14.Dutch.HDTV.XviD-LOL", []Language{LanguageDutch}},
		{"Castle.2009.S01E14.Japanese.HDTV.XviD-LOL", []Language{LanguageJapanese}},
		{"Castle.2009.S01E14.Cantonese.HDTV.XviD-LOL", []Language{LanguageCantonese}},
		{"Castle.2009.S01E14.Mandarin.HDTV.XviD-LOL", []Language{LanguageMandarin}},
		{"Castle.2009.S01E14.Korean.HDTV.XviD-LOL", []Language{LanguageKorean}},
		{"Castle.2009.S01E14.Russian.HDTV.XviD-LOL", []Language{LanguageRussian}},
		{"Castle.2009.S01E14.Ukrainian.HDTV.XviD-LOL", []Language{LanguageUkrainian}},
		{"Castle.2009.S01E14.Ukr.HDTV.XviD-LOL", []Language{LanguageUkrainian}},
		{"Castle.2009.S01E14.Polish.HDTV.XviD-LOL", []Language{LanguagePolish}},
		{"Castle.2009.S01E14.Vietnamese.HDTV.XviD-LOL", []Language{LanguageVietnamese}},
		{"Castle.2009.S01E14.Swedish.HDTV.XviD-LOL", []Language{LanguageSwedish}},
		{"Castle.2009.S01E14.Norwegian.HDTV.XviD-LOL", []Language{LanguageNorwegian}},
		{"Castle.2009.S01E14.Finnish.HDTV.XviD-LOL", []Language{LanguageFinnish}},
		{"Castle.2009.S01E14.Turkish.HDTV.XviD-LOL", []Language{LanguageTurkish}},
		{"Castle.2009.S01E14.Czech.HDTV.XviD-LOL", []Language{LanguageCzech}},
		{"Castle.2009.S01E14.Portuguese.HDTV.XviD-LOL", []Language{LanguagePortuguese}},
		{"Revolution S01E03 No Quarter 2012 WEB-DL 720p Nordic-philipo mkv", []Language{LanguageNordic}},
		{"Constantine.2014.S01E01.WEBRiP.H264.AAC.5.1-NL.SUBS", []Language{LanguageDutch}},
		{"Castle.2009.S01E14.HDTV.XviD.HUNDUB-LOL", []Language{LanguageHungarian}},
		{"Castle.2009.S01E14.HDTV.XviD.ENG.HUN-LOL", []Language{LanguageEnglish, LanguageHungarian}},
		{"Castle.2009.S01E14.HDTV.XviD.HUN-LOL", []Language{LanguageHungarian}},
		{"Castle.2009.S01E14.HDTV.XviD.CZ-LOL", []Language{LanguageCzech}},
		{"Peter.Rabbit.2.The.Runaway.2021.LATViAN.2160p.UHD.BLURAY.x265-UNDERDOG", []Language{LanguageLatvian}},
		{"Peter.Rabbit.2.The.Runaway.2021.LiTHUANiAN.2160p.UHD.BLURAY.x265-UNDERDOG", []Language{LanguageLithuanian}},
		{"Passengers.2016.German.DL.AC3.Dubbed.1080p.WebHD.h264.iNTERNAL-PsO", []Language{LanguageGerman, LanguageEnglish}},
		{"Smurfs.\u200b\u200bThe.\u200b\u200bLost.\u200b\u200bVillage.\u200b\u200b2017.\u200b\u200b1080p.\u200b\u200bBluRay.\u200b\u200bHebDub.\u200b\u200bx264-\u200b\u200biSrael", []Language{LanguageHebrew}},
		{"The Danish Girl 2015", []Language{LanguageEnglish}},
		{"Nocturnal Animals (2016) MULTi VFQ English [1080p] BluRay x264-PopHD", []Language{LanguageEnglish, LanguageFrench}},
		{"Wonder.Woman.2017.720p.BluRay.DD5.1.x264-TayTO.CZ-FTU", []Language{LanguageCzech}},
		{"Fantastic.Beasts.The.Crimes.Of.Grindelwald.2018.2160p.WEBRip.x265.10bit.HDR.DD5.1-GASMASK", []Language{LanguageEnglish}},
		{"Nick.and.Norahs.Infinite.Playlist.2008.CATALAN.MULTi.1080p.BluRay.x264-DESPACiTO", []Language{LanguageCatalan, LanguageEnglish}},
		{"Harry.Potter.And.The.Order.Of.The.Phoenix.2007.CHINESE.2160p.UHD.BluRay.X265-HOA", []Language{LanguageChinese}},
		{"Seven.Years.of.Night.2018.PL.DUAL.1080p.BluRay.x264-FLAME", []Language{LanguagePolish, LanguageEnglish}},
		{"Tenet.2020.THAI.2160p.UHD.BLURAY.x265-HOA", []Language{LanguageThai}},
		{"Tenet 2020 1080p Multi Eng Hin Tam iMax BluRay 10Bit DD5 1 H265-IPT", []Language{LanguageEnglish, LanguageHindi, LanguageTamil}},
		{"The Flying Guillotine 1975 CHI ENG DTS-HD DTS 1080p BluRay x264 HQ-TUSAHD", []Language{LanguageEnglish, LanguageChinese}},
		{"The Incredible Story Of The Giant Pear 2017 SWE DAN DTS-HD DTS MULTISUBS 1080p BluRay x264 HQ-TUSAHD", []Language{LanguageDanish, LanguageSwedish}},
		{"Wonder.Woman.1984.2020.PLDUB.DUAL.HDR10Plus.2160p.UHD.BluRay.x265.iNTERNAL-PLHD", []Language{LanguagePolish, LanguageEnglish}},
		{"Wadjda.2012.ARABiC.1080p.BluRay.x264-CONSTANT", []Language{LanguageArabic}},
		{"Arabic.12.1982.1080p.BluRay.x264-ROVERS", []Language{LanguageEnglish}},
		{"No.Country.for.Old.Men.1080p.BluRay.x264-HiGHTiMES", []Language{LanguageEnglish}},
		{"Cars.2.2011.ESTONiAN.DVDRip.x264-EMX", []Language{LanguageEstonian}},
		{"Cars.2.2011.EN.SE.FI.PAL.DVDR-AMIRITE", []Language{LanguageEnglish, LanguageSwedish}},
		{"Cars.2.2011.ENG.DK.NO.ICE.READ.NFO.PAL.DVDR-WILDER", []Language{LanguageEnglish, LanguageDanish, LanguageIcelandic, LanguageNorwegian}},
		{"Scarface.1983.CE.UNCUT.DVDRip.XviD.iNT-TURKiSO", []Language{LanguageEnglish}},
		{"Scarface.1983.20th.AE.iNTERNAL.DVDRip.XviD-MHQ", []Language{LanguageEnglish}},
		{"The.Conjuring.The.Devil.Made.Me.Do.It.2021.SUBFRENCH.2160p.WEB.H265-McNULTY", []Language{LanguageFrench}},
		{"Get.Him.To.The.Greek.UNRATED.FRENCH.720p.BluRay.x264-NERDHD", []Language{LanguageFrench}},
		{"Maennertrip.UNRATED.German.AC3.Dubbed.1080p.Bluray.x264-CIS", []Language{LanguageGerman}},
		{"Maennertrip.TS.MD.German.XViD.iNTERNAL-AOE", []Language{LanguageGerman}},
		{"Maennertrip.EXTENDED.German.AC3.BDRip.XviD-RedRay", []Language{LanguageGerman}},
		{"Get.Him.To.The.Greek.TRUEFRENCH.DVDRip.XviD-REVOLTE", []Language{LanguageFrench}},
		{"The.Social.Network.R5.LD.German.XviD-CinePlexx", []Language{LanguageGerman}},
		{"The.Social.Network.R5.LiNE.XviD-TWiZTED", []Language{LanguageEnglish}},
		{"Incassable.TRUE.FRENCH.PROPER.READ.NFO.DVDRiP.DiVX.SBC-KFT", []Language{LanguageFrench}},
		{"Space.Jam.A.New.Legacy.2021.ROMANiAN.2160p.UHD.BLURAY.x265-UNDERDOG", []Language{LanguageRomanian}},
		{"Space.Jam.A.New.Legacy.2021.RoDubbed.2160p.UHD.BLURAY.x265-UNDERDOG", []Language{LanguageRomanian}},
		{"Space.Jam.A.New.Legacy.2021.RO.2160p.UHD.BLURAY.x265-UNDERDOG", []Language{LanguageRomanian}},
		{"Spider-Man.No.Way.Home.2021.SLOVAK.2160p.UHD.BLURAY.x265-UNDERDOG", []Language{LanguageSlovak}},
		{"A.Serbian.Film.2010.SERBIAN.UnCut.DTS-HD.DTS.NORDICSUBS.1080p.BluRay.x264.HQ-TUSAHD", []Language{LanguageSerbian, LanguageNordic}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := ParseLanguage(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("ParseLanguage(%q)\n got  = %v\n want = %v", c.input, got, c.want)
			}
		})
	}
}

func TestIsMulti(t *testing.T) {
	cases := []string{
		"Ouija.Origin.of.Evil.2016.MULTi.TRUEFRENCH.1080p.BluRay.x264-MELBA",
		"Showdown.In.Little.Tokyo.1991.MULTI.VFQ.VFF.DTSHD-MASTER.1080p.BluRay.x264-ZombiE",
		"The.Polar.Express.2004.MULTI.VF2.1080p.BluRay.x264-PopHD",
		"Star.Trek.Der.Film.1979.German.DL.2160p.UHD.BluRay.HEVC-UNTHEVC",
	}

	for _, input := range cases {
		input := input
		t.Run(input, func(t *testing.T) {
			if !IsMulti(input) {
				t.Errorf("IsMulti(%q) = false, want true", input)
			}
		})
	}
}
