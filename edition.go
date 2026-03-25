package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var (
	editionInternalExp      = regexp2.MustCompile(`\b(INTERNAL)\b`, regexp2.IgnoreCase)
	editionRemasteredExp    = regexp2.MustCompile(`\b(Remastered|Anniversary|Restored)\b`, regexp2.IgnoreCase)
	editionImaxExp          = regexp2.MustCompile(`\b(IMAX)\b`, regexp2.IgnoreCase)
	editionUnratedExp       = regexp2.MustCompile(`\b(Uncensored|Unrated)\b`, regexp2.IgnoreCase)
	editionExtendedExp      = regexp2.MustCompile(`\b(Extended|Uncut|Ultimate|Rogue|Collector)\b`, regexp2.IgnoreCase)
	editionTheatricalExp    = regexp2.MustCompile(`\b(Theatrical)\b`, regexp2.IgnoreCase)
	editionDirectorsExp     = regexp2.MustCompile(`\b(Directors?)\b`, regexp2.IgnoreCase)
	editionFanExp           = regexp2.MustCompile(`\b(Despecialized|Fan.?Edit)\b`, regexp2.IgnoreCase)
	editionLimitedExp       = regexp2.MustCompile(`\b(LIMITED)\b`, regexp2.IgnoreCase)
	editionHdrExp           = regexp2.MustCompile(`\b(HDR)\b`, regexp2.IgnoreCase)
	editionThreeDExp        = regexp2.MustCompile(`\b(3D)\b`, regexp2.IgnoreCase)
	editionHsbsExp          = regexp2.MustCompile(`\b(Half-?SBS|HSBS)\b`, regexp2.IgnoreCase)
	editionSbsExp           = regexp2.MustCompile(`\b((?<!H|HALF-)SBS)\b`, regexp2.IgnoreCase)
	editionHouExp           = regexp2.MustCompile(`\b(HOU)\b`, regexp2.IgnoreCase)
	editionUhdExp           = regexp2.MustCompile(`\b(UHD)\b`, regexp2.IgnoreCase)
	editionOarExp           = regexp2.MustCompile(`\b(OAR)\b`, regexp2.IgnoreCase)
	editionDolbyVisionExp   = regexp2.MustCompile(`\b(DV(\b(HDR10|HLG|SDR))?)\b`, regexp2.IgnoreCase)
	editionHardcodedSubsExp = regexp2.MustCompile(`\b((?<hcsub>(\w+(?<!SOFT|HORRIBLE)SUBS?))|(?<hc>(HC|SUBBED)))\b`, regexp2.IgnoreCase)
	editionDeletedScenesExp = regexp2.MustCompile(`\b((Bonus.)?Deleted.Scenes)\b`, regexp2.IgnoreCase)
	editionBonusContentExp  = regexp2.MustCompile(`\b((Bonus|Extras|Behind.the.Scenes|Making.of|Interviews|Featurettes|Outtakes|Bloopers|Gag.Reel).(?!(Deleted.Scenes)))\b`, regexp2.IgnoreCase)
	editionBwExp            = regexp2.MustCompile(`\b(BW)\b`, regexp2.IgnoreCase)
)

// ParseEdition detects edition flags in a release title.
func ParseEdition(title string, parsedTitle ...string) Edition {
	var pt string
	if len(parsedTitle) > 0 {
		pt = parsedTitle[0]
	} else {
		pt = ParseTitleAndYear(title).Title
	}
	withoutTitle := strings.ToLower(strings.ReplaceAll(strings.Replace(title, ".", " ", 1), pt, ""))

	return Edition{
		Internal:      mustMatch(editionInternalExp, withoutTitle),
		Limited:       mustMatch(editionLimitedExp, withoutTitle),
		Remastered:    mustMatch(editionRemasteredExp, withoutTitle),
		Extended:      mustMatch(editionExtendedExp, withoutTitle),
		Theatrical:    mustMatch(editionTheatricalExp, withoutTitle),
		Directors:     mustMatch(editionDirectorsExp, withoutTitle),
		Unrated:       mustMatch(editionUnratedExp, withoutTitle),
		IMAX:          mustMatch(editionImaxExp, withoutTitle),
		FanEdit:       mustMatch(editionFanExp, withoutTitle),
		HDR:           mustMatch(editionHdrExp, withoutTitle),
		ThreeD:        mustMatch(editionThreeDExp, withoutTitle),
		HSBS:          mustMatch(editionHsbsExp, withoutTitle),
		SBS:           mustMatch(editionSbsExp, withoutTitle),
		HOU:           mustMatch(editionHouExp, withoutTitle),
		UHD:           mustMatch(editionUhdExp, withoutTitle),
		OAR:           mustMatch(editionOarExp, withoutTitle),
		DolbyVision:   mustMatch(editionDolbyVisionExp, withoutTitle),
		HardcodedSubs: mustMatch(editionHardcodedSubsExp, withoutTitle),
		DeletedScenes: mustMatch(editionDeletedScenesExp, withoutTitle),
		BonusContent:  mustMatch(editionBonusContentExp, withoutTitle),
		BW:            mustMatch(editionBwExp, withoutTitle),
	}
}
