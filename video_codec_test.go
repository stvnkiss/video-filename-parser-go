package videoparser

import "testing"

func TestParseVideoCodec(t *testing.T) {
	cases := []struct {
		input     string
		wantCodec VideoCodec
	}{
		{"Terminator 3 Rise of The Machines 2003 HDDVD XvidHD 720p-NPW", VideoCodecXVID},
		{"Cloverfield 2008 BRRip XvidHD 720p-NPW", VideoCodecXVID},
		{"The Interview 2014 1080p WEB-DL x264 AAC MerryXmas", VideoCodecX264},
		{"Half Baked 1998 HDRip XviD AC3-FLAWL3SS", VideoCodecXVID},
		{"Hidden Figures 2016 DVDSCR XVID-FrangoAssado", VideoCodecXVID},
		{"Vice 2018 DVDScr Xvid AC3 HQ Hive-CM8", VideoCodecXVID},
		{"The Dark Knight[2008]DvDrip-aXXo [pendhu]", ""},         // undefined → ""
		{"Bridesmaids[2011][Unrated Edition]DvDrip AC3-aXXo", ""}, // undefined → ""
		{"Get Out 2017 BluRay 10Bit 1080p DD5 1 H265-d3g", VideoCodecH265},
		{"Minions 2015 720p HC HDRip X265 AC3 TiTAN", VideoCodecX265},
		{"Marvel's The Avengers 2012 BluRay 1080p DD5 1 10Bit H265-d3g", VideoCodecH265},
		{"Exodus Gods and Kings 2014 MULTi 2160p UHD BluRay x265-SESKAPiLE", VideoCodecX265},
		{"The Incredibles 2004 BluRay x264-jlw", VideoCodecX264},
		{"Jack Reacher 2012 720p BluRay X264-AMIABLE", VideoCodecX264},
		{"Super Troopers 2 2018 1080p WEB-DL H264 AC3-EVO", VideoCodecH264},
		{"The.Middle.720p.HEVC-MeGusta-Pre", VideoCodecX265},
		{"Cloud.Atlas.2012.BluRay.1080p.VC1.5.1.WMV-INSECTS", VideoCodecWMV},
		{"The.Book.Of.Eli.2010.Bluray.VC1.1080P.5.1.WMV-NOVO", VideoCodecWMV},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := parseVideoCodec(c.input)
			if got.Codec != c.wantCodec {
				t.Errorf("Codec = %q, want %q", got.Codec, c.wantCodec)
			}
		})
	}
}
