package videoparser

import "github.com/dlclark/regexp2"

var audioCodecExp = regexp2.MustCompile(
	`\b(?<mp3>(LAME(?:\d)+-?(?:\d)+)|(mp3))\b`+
		`|\b(?<mp2>(mp2))\b`+
		`|\b(?<dolby>(Dolby)|(Dolby-?Digital)|(DD)|(AC3D?))\b`+
		`|\b(?<dolbyatmos>(Dolby-?Atmos))\b`+
		`|\b(?<aac>(AAC))(\d?.?\d?)(ch)?\b`+
		`|\b(?<eac3>(EAC3|DDP|DD\+))\b`+
		`|\b(?<flac>(FLAC))\b`+
		`|\b(?<dtshd>(DTS-?HD)|(DTS(?=-?MA)|(DTS-X)))\b`+
		`|\b(?<dts>(DTS))\b`+
		`|\b(?<truehd>(True-?HD))\b`+
		`|\b(?<opus>(Opus))\b`+
		`|\b(?<vorbis>(Vorbis))\b`+
		`|\b(?<pcm>(PCM))\b`+
		`|\b(?<lpcm>(LPCM))\b`,
	regexp2.IgnoreCase,
)

type audioCodecResult struct {
	Codec  AudioCodec
	Source string
}

func parseAudioCodec(title string) audioCodecResult {
	m, err := audioCodecExp.FindStringMatch(title)
	if err != nil || m == nil {
		return audioCodecResult{}
	}

	groupVal := func(name string) string {
		g := m.GroupByName(name)
		if g != nil && g.Length > 0 {
			return g.String()
		}
		return ""
	}

	// Order matters: more specific patterns first
	if v := groupVal("aac"); v != "" {
		return audioCodecResult{Codec: AudioCodecAAC, Source: v}
	}
	if v := groupVal("dolbyatmos"); v != "" {
		return audioCodecResult{Codec: AudioCodecEAC3, Source: v}
	}
	if v := groupVal("dolby"); v != "" {
		return audioCodecResult{Codec: AudioCodecDOLBY, Source: v}
	}
	if v := groupVal("dtshd"); v != "" {
		return audioCodecResult{Codec: AudioCodecDTSHD, Source: v}
	}
	if v := groupVal("dts"); v != "" {
		return audioCodecResult{Codec: AudioCodecDTS, Source: v}
	}
	if v := groupVal("flac"); v != "" {
		return audioCodecResult{Codec: AudioCodecFLAC, Source: v}
	}
	if v := groupVal("truehd"); v != "" {
		return audioCodecResult{Codec: AudioCodecTRUEHD, Source: v}
	}
	if v := groupVal("mp3"); v != "" {
		return audioCodecResult{Codec: AudioCodecMP3, Source: v}
	}
	if v := groupVal("mp2"); v != "" {
		return audioCodecResult{Codec: AudioCodecMP2, Source: v}
	}
	if v := groupVal("pcm"); v != "" {
		return audioCodecResult{Codec: AudioCodecPCM, Source: v}
	}
	if v := groupVal("lpcm"); v != "" {
		return audioCodecResult{Codec: AudioCodecLPCM, Source: v}
	}
	if v := groupVal("opus"); v != "" {
		return audioCodecResult{Codec: AudioCodecOPUS, Source: v}
	}
	if v := groupVal("vorbis"); v != "" {
		return audioCodecResult{Codec: AudioCodecVORBIS, Source: v}
	}
	if v := groupVal("eac3"); v != "" {
		return audioCodecResult{Codec: AudioCodecEAC3, Source: v}
	}
	return audioCodecResult{}
}
