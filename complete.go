package videoparser

import "github.com/dlclark/regexp2"

var (
	completeDvdExp = regexp2.MustCompile(`\b(NTSC|PAL)?.DVDR\b`, regexp2.IgnoreCase)
	completeExp    = regexp2.MustCompile(`\b(COMPLETE)\b`, regexp2.IgnoreCase)
)

func isCompleteDvd(title string) bool {
	ok, _ := completeDvdExp.MatchString(title)
	return ok
}

func isComplete(title string) bool {
	ok, _ := completeExp.MatchString(title)
	return ok || isCompleteDvd(title)
}
