package videoparser

import (
	"testing"
	"time"
)

func TestFilenameParse_Movies(t *testing.T) {
	cases := []struct {
		input      string
		title      string
		year       string
		resolution Resolution
		sources    []Source
		codec      VideoCodec
		audioCodec AudioCodec
		group      string
		revision   Revision
		languages  []Language
		multi      bool
		complete   bool
		edition    *Edition
	}{
		{
			input:      "Whats.Eating.Gilbert.Grape.1993.720p.BluRay.x264-SiNNERS",
			title:      "Whats Eating Gilbert Grape",
			year:       "1993",
			resolution: R720P,
			sources:    []Source{SourceBLURAY},
			codec:      VideoCodecX264,
			group:      "SiNNERS",
			revision:   Revision{Version: 1, Real: 0},
			languages:  []Language{LanguageEnglish},
		},
		{
			input:      "Timecop.1994.PROPER.1080p.BluRay.x264-Japhson",
			title:      "Timecop",
			year:       "1994",
			resolution: R1080P,
			sources:    []Source{SourceBLURAY},
			codec:      VideoCodecX264,
			group:      "Japhson",
			revision:   Revision{Version: 2, Real: 0},
			languages:  []Language{LanguageEnglish},
		},
		{
			input:      "This.is.40.2012.PROPER.UNRATED.720p.BluRay.MULti.x264-Felony",
			title:      "This is 40",
			year:       "2012",
			resolution: R720P,
			sources:    []Source{SourceBLURAY},
			codec:      VideoCodecX264,
			group:      "Felony",
			revision:   Revision{Version: 2, Real: 0},
			languages:  []Language{LanguageEnglish},
			multi:      true,
			edition:    &Edition{Unrated: true},
		},
		{
			input:      "Spider-Man Far from Home.2019.1080p.HDRip.X264.AC3-EVO",
			title:      "Spider-Man Far from Home",
			year:       "2019",
			resolution: R1080P,
			sources:    []Source{SourceWEBDL},
			codec:      VideoCodecX264,
			audioCodec: AudioCodecDOLBY,
			group:      "EVO",
			revision:   Revision{Version: 1, Real: 0},
			languages:  []Language{LanguageEnglish},
		},
		{
			input:      "Togo 2019 2160p HDR DSNP WEBRip DDPAtmos 5 1 X265-TrollUHD",
			title:      "Togo",
			year:       "2019",
			resolution: R2160P,
			sources:    []Source{SourceWEBRIP},
			group:      "TrollUHD",
		},
		{
			input:      "Ex Machina 2015 UHD BluRay 2160p DTS-X 7 1 HDR x265 10bit-CHD",
			title:      "Ex Machina",
			year:       "2015",
			group:      "CHD",
			resolution: R2160P,
		},
		{
			input:      "Apprentice.2016.COMPLETE.BLURAY-UNRELiABLE",
			title:      "Apprentice",
			year:       "2016",
			group:      "UNRELiABLE",
			resolution: R1080P,
			complete:   true,
		},
		{
			input:      "Indiana.Jones.and.the.Temple.of.Doom.1984.Complete.UHD.Bluray-JONES",
			title:      "Indiana Jones and the Temple of Doom",
			year:       "1984",
			group:      "JONES",
			resolution: R2160P,
			complete:   true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := FilenameParse(c.input, false)

			if c.title != "" && got.Title != c.title {
				t.Errorf("Title = %q, want %q", got.Title, c.title)
			}
			if c.year != "" && got.Year != c.year {
				t.Errorf("Year = %q, want %q", got.Year, c.year)
			}
			if c.resolution != "" && got.Resolution != c.resolution {
				t.Errorf("Resolution = %q, want %q", got.Resolution, c.resolution)
			}
			if c.sources != nil {
				if len(got.Sources) == 0 || got.Sources[0] != c.sources[0] {
					t.Errorf("Sources = %v, want %v", got.Sources, c.sources)
				}
			}
			if c.codec != "" && got.VideoCodec != c.codec {
				t.Errorf("VideoCodec = %q, want %q", got.VideoCodec, c.codec)
			}
			if c.audioCodec != "" && got.AudioCodec != c.audioCodec {
				t.Errorf("AudioCodec = %q, want %q", got.AudioCodec, c.audioCodec)
			}
			if c.group != "" && got.Group != c.group {
				t.Errorf("Group = %q, want %q", got.Group, c.group)
			}
			if c.revision.Version != 0 && got.Revision != c.revision {
				t.Errorf("Revision = %+v, want %+v", got.Revision, c.revision)
			}
			if c.languages != nil && len(got.Languages) > 0 && got.Languages[0] != c.languages[0] {
				t.Errorf("Languages[0] = %q, want %q", got.Languages[0], c.languages[0])
			}
			if c.multi && !got.Multi {
				t.Errorf("Multi = false, want true")
			}
			if c.complete && !got.Complete {
				t.Errorf("Complete = false, want true")
			}
			if c.edition != nil {
				if c.edition.Unrated && !got.Edition.Unrated {
					t.Errorf("Edition.Unrated = false, want true")
				}
			}
		})
	}
}

func TestFilenameParse_TV(t *testing.T) {
	cases := []struct {
		input      string
		title      string
		resolution Resolution
		sources    []Source
		codec      VideoCodec
		group      string
		revision   Revision
		languages  []Language
		seasons    []int
		episodes   []int
	}{
		{
			input:      "Its Always Sunny in Philadelphia S14E04 720p WEB H264-METCON",
			title:      "Its Always Sunny in Philadelphia",
			resolution: R720P,
			sources:    []Source{SourceWEBDL},
			codec:      VideoCodecH264,
			group:      "METCON",
			revision:   Revision{Version: 1, Real: 0},
			languages:  []Language{LanguageEnglish},
			seasons:    []int{14},
			episodes:   []int{4},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := FilenameParse(c.input, true)

			if !got.IsTv {
				t.Errorf("IsTv = false, want true")
			}
			if c.title != "" && got.Title != c.title {
				t.Errorf("Title = %q, want %q", got.Title, c.title)
			}
			if c.resolution != "" && got.Resolution != c.resolution {
				t.Errorf("Resolution = %q, want %q", got.Resolution, c.resolution)
			}
			if c.sources != nil {
				if len(got.Sources) == 0 || got.Sources[0] != c.sources[0] {
					t.Errorf("Sources = %v, want %v", got.Sources, c.sources)
				}
			}
			if c.codec != "" && got.VideoCodec != c.codec {
				t.Errorf("VideoCodec = %q, want %q", got.VideoCodec, c.codec)
			}
			if c.group != "" && got.Group != c.group {
				t.Errorf("Group = %q, want %q", got.Group, c.group)
			}
			if c.revision.Version != 0 && got.Revision != c.revision {
				t.Errorf("Revision = %+v, want %+v", got.Revision, c.revision)
			}
			if c.seasons != nil {
				if len(got.Seasons) != len(c.seasons) || (len(c.seasons) > 0 && got.Seasons[0] != c.seasons[0]) {
					t.Errorf("Seasons = %v, want %v", got.Seasons, c.seasons)
				}
			}
			if c.episodes != nil {
				if len(got.EpisodeNumbers) != len(c.episodes) || (len(c.episodes) > 0 && got.EpisodeNumbers[0] != c.episodes[0]) {
					t.Errorf("EpisodeNumbers = %v, want %v", got.EpisodeNumbers, c.episodes)
				}
			}
		})
	}
}

func TestFilenameParse_DailyTV(t *testing.T) {
	cases := []struct {
		input      string
		title      string
		resolution Resolution
		sources    []Source
		codec      VideoCodec
		group      string
		revision   Revision
		languages  []Language
		airDate    time.Time
	}{
		{
			input:      "NFL 2019 10 06 Chicago Bears vs Oakland Raiders Highlights 720p HEVC x265-MeGusta",
			title:      "NFL",
			resolution: R720P,
			sources:    []Source{SourceWEBDL},
			codec:      VideoCodecX265,
			group:      "MeGusta",
			revision:   Revision{Version: 1, Real: 0},
			languages:  []Language{LanguageEnglish},
			airDate:    time.Date(2019, time.October, 6, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := FilenameParse(c.input, true)

			if !got.IsTv {
				t.Errorf("IsTv = false, want true")
			}
			if c.title != "" && got.Title != c.title {
				t.Errorf("Title = %q, want %q", got.Title, c.title)
			}
			if c.resolution != "" && got.Resolution != c.resolution {
				t.Errorf("Resolution = %q, want %q", got.Resolution, c.resolution)
			}
			if c.codec != "" && got.VideoCodec != c.codec {
				t.Errorf("VideoCodec = %q, want %q", got.VideoCodec, c.codec)
			}
			if c.group != "" && got.Group != c.group {
				t.Errorf("Group = %q, want %q", got.Group, c.group)
			}
			if got.AirDate == nil || !got.AirDate.Equal(c.airDate) {
				t.Errorf("AirDate = %v, want %v", got.AirDate, c.airDate)
			}
		})
	}
}
