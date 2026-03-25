// Package videoparser parses video release / file names and extracts metadata
// such as title, year, codec, resolution, source, language, and TV show
// season/episode information.
//
// It is a Go port of the TypeScript @ctrl/video-filename-parser library which
// is itself based on Radarr/Sonarr's filename parsing logic.
package videoparser

// FilenameParse is the main entry point. It parses a release/file name and
// returns a ParsedFilename populated with all detected metadata.
//
// Set isTv to true to also parse season/episode information.
func FilenameParse(name string, isTv bool) ParsedFilename {
	titleAndYear := ParseTitleAndYear(name)
	parsedTitle := titleAndYear.Title

	var title, year string
	if !isTv {
		title = titleAndYear.Title
		year = titleAndYear.Year
	}

	edition := ParseEdition(name, parsedTitle)
	vc := parseVideoCodec(name)
	ac := parseAudioCodec(name)
	ch := parseAudioChannels(name)
	group := ParseGroup(name, parsedTitle)
	languages := ParseLanguage(name, parsedTitle)
	quality := ParseQuality(name, vc.Codec)
	multi := IsMulti(name)
	complete := isComplete(name)

	result := ParsedFilename{
		Title:         title,
		Year:          year,
		Edition:       edition,
		Resolution:    quality.Resolution,
		Sources:       quality.Sources,
		VideoCodec:    vc.Codec,
		AudioCodec:    ac.Codec,
		AudioChannels: ch.Channels,
		Revision:      quality.Revision,
		Group:         group,
		Languages:     languages,
		Multi:         multi,
		Complete:      complete,
	}

	if isTv {
		season, err := ParseSeason(name)
		if err == nil && season != nil {
			result.IsTv = true
			result.Title = season.SeriesTitle
			if result.Title == "" {
				result.Title = title
			}
			result.Seasons = season.Seasons
			result.EpisodeNumbers = season.EpisodeNumbers
			result.AirDate = season.AirDate
			result.FullSeason = season.FullSeason
			result.IsPartialSeason = season.IsPartialSeason
			result.IsMultiSeason = season.IsMultiSeason
			result.IsSeasonExtra = season.IsSeasonExtra
			result.IsSpecial = season.IsSpecial
			result.SeasonPart = season.SeasonPart
		}
	}

	return result
}
