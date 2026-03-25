package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var fileExtensions = map[string]bool{
	// Unknown
	".webm": true,
	// SDTV
	".m4v": true, ".3gp": true, ".nsv": true, ".ty": true, ".strm": true,
	".rm": true, ".rmvb": true, ".m3u": true, ".ifo": true, ".mov": true,
	".qt": true, ".divx": true, ".xvid": true, ".bivx": true, ".nrg": true,
	".pva": true, ".wmv": true, ".asf": true, ".asx": true, ".ogm": true,
	".ogv": true, ".m2v": true, ".avi": true, ".bin": true, ".dat": true,
	".dvr-ms": true, ".mpg": true, ".mpeg": true, ".mp4": true, ".avc": true,
	".vp3": true, ".svq3": true, ".nuv": true, ".viv": true, ".dv": true,
	".fli": true, ".flv": true, ".wpl": true,
	// DVD
	".img": true, ".iso": true, ".vob": true,
	// HD
	".mkv": true, ".mk3d": true, ".ts": true, ".wtv": true,
	// Bluray
	".m2ts": true,
}

var fileExtensionExp = regexp2.MustCompile(`\.[a-z0-9]{2,4}$`, regexp2.IgnoreCase)

// RemoveFileExtension strips recognised video file extensions from a title string.
func RemoveFileExtension(title string) string {
	m, err := fileExtensionExp.FindStringMatch(title)
	if err != nil || m == nil {
		return title
	}
	ext := strings.ToLower(m.String())
	if fileExtensions[ext] {
		return title[:len(title)-len(m.String())]
	}
	return title
}
