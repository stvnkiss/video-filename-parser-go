package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var (
	blurayExp    = regexp2.MustCompile(`\b(?<bluray>M?Blu-?Ray|HDDVD|BD|UHDBD|BDISO|BDMux|BD25|BD50|BR.?DISK|Bluray(1080|720)p?|BD(1080|720)p?)\b`, regexp2.IgnoreCase)
	webdlExp     = regexp2.MustCompile(`\b(?<webdl>WEB[-_. ]DL|HDRIP|WEBDL|WEB-DLMux|NF|APTV|NETFLIX|NetflixU?HD|DSNY|DSNP|HMAX|AMZN|AmazonHD|iTunesHD|MaxdomeHD|WebHD\b|[. ]WEB[. ](?:[xh]26[45]|DD5[. ]1)|\d+0p[. ]WEB[. ]|\b\s\/\sWEB\s\/\s\b|AMZN[. ]WEB[. ])\b`, regexp2.IgnoreCase)
	webripExp    = regexp2.MustCompile(`\b(?<webrip>WebRip|Web-Rip|WEBCap|WEBMux)\b`, regexp2.IgnoreCase)
	hdtvExp      = regexp2.MustCompile(`\b(?<hdtv>HDTV)\b`, regexp2.IgnoreCase)
	bdripExp     = regexp2.MustCompile(`\b(?<bdrip>BDRip|UHDBDRip|HD[-_. ]?DVDRip)\b`, regexp2.IgnoreCase)
	brripExp     = regexp2.MustCompile(`\b(?<brrip>BRRip)\b`, regexp2.IgnoreCase)
	dvdrExp      = regexp2.MustCompile(`\b(?<dvdr>DVD-R|DVDR)\b`, regexp2.IgnoreCase)
	dvdExp       = regexp2.MustCompile(`\b(?<dvd>DVD9?|DVDRip|NTSC|PAL|xvidvd|DvDivX)\b`, regexp2.IgnoreCase)
	dsrExp       = regexp2.MustCompile(`\b(?<dsr>WS[-_. ]DSR|DSR)\b`, regexp2.IgnoreCase)
	regionalExp  = regexp2.MustCompile(`\b(?<regional>R[0-9]{1}|REGIONAL)\b`, regexp2.IgnoreCase)
	ppvExp       = regexp2.MustCompile(`\b(?<ppv>PPV)\b`, regexp2.IgnoreCase)
	scrExp       = regexp2.MustCompile(`\b(?<scr>SCR|SCREENER|DVDSCR|(DVD|WEB).?SCREENER)\b`, regexp2.IgnoreCase)
	tsExp        = regexp2.MustCompile(`\b(?<ts>TS|TELESYNC|HD-TS|HDTS|PDVD|TSRip|HDTSRip)\b`, regexp2.IgnoreCase)
	tcExp        = regexp2.MustCompile(`\b(?<tc>TC|TELECINE|HD-TC|HDTC)\b`, regexp2.IgnoreCase)
	camExp       = regexp2.MustCompile(`\b(?<cam>CAMRIP|CAM|HDCAM|HD-CAM)\b`, regexp2.IgnoreCase)
	workprintExp = regexp2.MustCompile(`\b(?<workprint>WORKPRINT|WP)\b`, regexp2.IgnoreCase)
	pdtvExp      = regexp2.MustCompile(`\b(?<pdtv>PDTV)\b`, regexp2.IgnoreCase)
	sdtvExp      = regexp2.MustCompile(`\b(?<sdtv>SDTV)\b`, regexp2.IgnoreCase)
	tvripExp     = regexp2.MustCompile(`\b(?<tvrip>TVRip)\b`, regexp2.IgnoreCase)
)

type sourceGroups struct {
	bluray    bool
	webdl     bool
	webrip    bool
	hdtv      bool
	bdrip     bool
	brrip     bool
	scr       bool
	dvdr      bool
	dvd       bool
	dsr       bool
	regional  bool
	ppv       bool
	ts        bool
	tc        bool
	cam       bool
	workprint bool
	pdtv      bool
	sdtv      bool
	tvrip     bool
}

func mustMatch(re *regexp2.Regexp, s string) bool {
	ok, _ := re.MatchString(s)
	return ok
}

func normalizeSourceTitle(title string) string {
	r := strings.NewReplacer("_", " ", "[", " ", "]", " ")
	return strings.TrimSpace(r.Replace(title))
}

func parseSourceGroups(title string) sourceGroups {
	n := normalizeSourceTitle(title)
	return sourceGroups{
		bluray:    mustMatch(blurayExp, n),
		webdl:     mustMatch(webdlExp, n),
		webrip:    mustMatch(webripExp, n),
		hdtv:      mustMatch(hdtvExp, n),
		bdrip:     mustMatch(bdripExp, n),
		brrip:     mustMatch(brripExp, n),
		scr:       mustMatch(scrExp, n),
		dvdr:      mustMatch(dvdrExp, n),
		dvd:       mustMatch(dvdExp, n),
		dsr:       mustMatch(dsrExp, n),
		regional:  mustMatch(regionalExp, n),
		ppv:       mustMatch(ppvExp, n),
		ts:        mustMatch(tsExp, n),
		tc:        mustMatch(tcExp, n),
		cam:       mustMatch(camExp, n),
		workprint: mustMatch(workprintExp, n),
		pdtv:      mustMatch(pdtvExp, n),
		sdtv:      mustMatch(sdtvExp, n),
		tvrip:     mustMatch(tvripExp, n),
	}
}

func parseSource(title string, groups ...sourceGroups) []Source {
	var g sourceGroups
	if len(groups) > 0 {
		g = groups[0]
	} else {
		g = parseSourceGroups(title)
	}

	var result []Source

	if g.bluray || g.bdrip || g.brrip {
		result = append(result, SourceBLURAY)
	}
	if g.webrip {
		result = append(result, SourceWEBRIP)
	}
	if !g.webrip && g.webdl {
		result = append(result, SourceWEBDL)
	}
	if g.dvdr || (g.dvd && !g.scr) {
		result = append(result, SourceDVD)
	}
	if g.ppv {
		result = append(result, SourcePPV)
	}
	if g.workprint {
		result = append(result, SourceWORKPRINT)
	}
	if g.pdtv || g.sdtv || g.dsr || g.tvrip || g.hdtv {
		result = append(result, SourceTV)
	}
	if g.cam {
		result = append(result, SourceCAM)
	}
	if g.ts {
		result = append(result, SourceTELESYNC)
	}
	if g.tc {
		result = append(result, SourceTELECINE)
	}
	if g.scr {
		result = append(result, SourceSCREENER)
	}
	return result
}
