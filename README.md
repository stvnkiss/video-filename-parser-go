# video-filename-parser (Go)

A Go port of [@ctrl/video-filename-parser](https://github.com/scttcper/video-filename-parser), a TypeScript library for parsing video release/file names. The parsing logic is based on [Radarr](https://github.com/Radarr/Radarr) and [Sonarr](https://github.com/Sonarr/Sonarr).

Given a release name like `Whats.Eating.Gilbert.Grape.1993.720p.BluRay.x264-SiNNERS`, it extracts structured metadata: title, year, resolution, source, video/audio codec, audio channels, release group, languages, edition flags, and TV season/episode information.

## Installation

```sh
go get github.com/stvnkiss/video-filename-parser-go
```

## Usage

```go
import videoparser "github.com/stvnkiss/video-filename-parser-go"

// Parse a movie
result := videoparser.FilenameParse("Whats.Eating.Gilbert.Grape.1993.720p.BluRay.x264-SiNNERS", false)
// result.Title      → "Whats Eating Gilbert Grape"
// result.Year       → "1993"
// result.Resolution → "720P"
// result.Sources    → ["BLURAY"]
// result.VideoCodec → "x264"
// result.Group      → "SiNNERS"

// Parse a TV episode
tv := videoparser.FilenameParse("Its Always Sunny in Philadelphia S14E04 720p WEB H264-METCON", true)
// tv.IsTv           → true
// tv.Title          → "Its Always Sunny in Philadelphia"
// tv.Seasons        → [14]
// tv.EpisodeNumbers → [4]
// tv.Resolution     → "720P"
// tv.Sources        → ["WEBDL"]
// tv.VideoCodec     → "h264"
// tv.Group          → "METCON"
```

Set `isTv` to `true` to enable TV-specific parsing (season/episode numbers, air dates, specials).

## Parsed fields

`FilenameParse` returns a `ParsedFilename` struct:

| Field             | Type            | Description                                     |
|-------------------|-----------------|-------------------------------------------------|
| `Title`           | `string`        | Movie or series title                           |
| `Year`            | `string`        | Release year (movies only)                      |
| `Resolution`      | `Resolution`    | `480P` / `720P` / `1080P` / `2160P` / …        |
| `Sources`         | `[]Source`      | `BLURAY`, `WEBDL`, `WEBRIP`, `DVD`, `TV`, …    |
| `VideoCodec`      | `VideoCodec`    | `x264`, `x265`, `h264`, `h265`, `xvid`, …      |
| `AudioCodec`      | `AudioCodec`    | `AAC`, `DTS`, `DTS-HD`, `Dolby Digital`, …     |
| `AudioChannels`   | `Channels`      | `stereo`, `mono`, `5.1`, `7.1`                 |
| `Group`           | `string`        | Release group                                   |
| `Languages`       | `[]Language`    | Detected languages                              |
| `Revision`        | `Revision`      | `Version` (PROPER=2) and `Real` counters        |
| `Edition`         | `Edition`       | Struct of bool flags (see below)                |
| `Multi`           | `bool`          | Multi-language release                          |
| `Complete`        | `bool`          | Complete series/season pack                     |
| `IsTv`            | `bool`          | Whether TV parsing matched                      |
| `Seasons`         | `[]int`         | Season numbers (TV)                             |
| `EpisodeNumbers`  | `[]int`         | Episode numbers (TV)                            |
| `AirDate`         | `*time.Time`    | Air date for daily shows (TV)                   |
| `FullSeason`      | `bool`          | Full season pack (TV)                           |
| `IsPartialSeason` | `bool`          | Partial season pack (TV)                        |
| `IsMultiSeason`   | `bool`          | Multi-season pack (TV)                          |
| `IsSeasonExtra`   | `bool`          | Season extras (TV)                              |
| `IsSpecial`       | `bool`          | Special episode (TV)                            |
| `SeasonPart`      | `int`           | Season part number (TV)                         |

### Edition flags

The `Edition` struct contains boolean flags for:
`Internal`, `Limited`, `Remastered`, `Extended`, `Theatrical`, `Directors`, `Unrated`, `IMAX`, `FanEdit`, `HDR`, `BW`, `ThreeD`, `HSBS`, `SBS`, `HOU`, `UHD`, `OAR`, `DolbyVision`, `HardcodedSubs`, `DeletedScenes`, `BonusContent`

## Requirements

- Go 1.21+
- [`github.com/dlclark/regexp2`](https://github.com/dlclark/regexp2) — PCRE-compatible regex engine (handles lookahead/lookbehind patterns from the original library)

## Running tests

```sh
cd go
go test ./...
```

## Origin

This is a faithful Go port of the TypeScript [`@ctrl/video-filename-parser`](https://github.com/scttcper/video-filename-parser) library. All parsing patterns and logic are derived from that project, which is itself based on Radarr/Sonarr's release name parsing.
