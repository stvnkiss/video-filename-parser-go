package videoparser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

var (
	langEnglishExp    = regexp2.MustCompile(`\b(english|eng|EN|FI)\b`, regexp2.IgnoreCase)
	langDanishExp     = regexp2.MustCompile(`\b(DK|DAN|danish)\b`, regexp2.IgnoreCase)
	langSwedishExp    = regexp2.MustCompile(`\b(SE|SWE|swedish)\b`, regexp2.IgnoreCase)
	langIcelandicExp  = regexp2.MustCompile(`\b(ice|Icelandic)\b`, regexp2.IgnoreCase)
	langChineseExp    = regexp2.MustCompile(`\b(chi|chinese)\b`, regexp2.IgnoreCase)
	langItalianExp    = regexp2.MustCompile(`\b(ita|italian)\b`, regexp2.IgnoreCase)
	langGermanExp     = regexp2.MustCompile(`\b(german|videomann)\b`, regexp2.IgnoreCase)
	langFlemishExp    = regexp2.MustCompile(`\b(flemish)\b`, regexp2.IgnoreCase)
	langGreekExp      = regexp2.MustCompile(`\b(greek)\b`, regexp2.IgnoreCase)
	langFrenchExp     = regexp2.MustCompile(`\b(FR|FRENCH|VOSTFR|VO|VFF|VFQ|VF2|TRUEFRENCH|SUBFRENCH)\b`, regexp2.IgnoreCase)
	langRussianExp    = regexp2.MustCompile(`\b(russian|rus)\b`, regexp2.IgnoreCase)
	langNorwegianExp  = regexp2.MustCompile(`\b(norwegian|NO)\b`, regexp2.IgnoreCase)
	langHungarianExp  = regexp2.MustCompile(`\b(HUNDUB|HUN|hungarian)\b`, regexp2.IgnoreCase)
	langHebDubExp     = regexp2.MustCompile(`\b(HebDub)\b`, regexp2.IgnoreCase)
	langCzSkExp       = regexp2.MustCompile(`\b(CZ|SK)\b`, regexp2.IgnoreCase)
	langUkrExp        = regexp2.MustCompile(`(?<ukrainian>\bukr\b)`, regexp2.IgnoreCase)
	langPolishExp     = regexp2.MustCompile(`\b(PL|PLDUB|POLISH)\b`, regexp2.IgnoreCase)
	langDutchExp      = regexp2.MustCompile(`\b(nl|dutch)\b`, regexp2.IgnoreCase)
	langHindiExp      = regexp2.MustCompile(`\b(HIN|Hindi)\b`, regexp2.IgnoreCase)
	langTamilExp      = regexp2.MustCompile(`\b(TAM|Tamil)\b`, regexp2.IgnoreCase)
	langArabicExp     = regexp2.MustCompile(`\b(Arabic)\b`, regexp2.IgnoreCase)
	langLatvianExp    = regexp2.MustCompile(`\b(Latvian)\b`, regexp2.IgnoreCase)
	langLithuanianExp = regexp2.MustCompile(`\b(Lithuanian)\b`, regexp2.IgnoreCase)
	langRomanianExp   = regexp2.MustCompile(`\b(RO|Romanian|rodubbed)\b`, regexp2.IgnoreCase)
	langSlovakExp     = regexp2.MustCompile(`\b(SK|Slovak)\b`, regexp2.IgnoreCase)
	langBrazilianExp  = regexp2.MustCompile(`\b(Brazilian)\b`, regexp2.IgnoreCase)
	langPersianExp    = regexp2.MustCompile(`\b(Persian)\b`, regexp2.IgnoreCase)
	langBengaliExp    = regexp2.MustCompile(`\b(Bengali)\b`, regexp2.IgnoreCase)
	langBulgarianExp  = regexp2.MustCompile(`\b(Bulgarian)\b`, regexp2.IgnoreCase)
	langSerbianExp    = regexp2.MustCompile(`\b(Serbian)\b`, regexp2.IgnoreCase)
	langNordicExp     = regexp2.MustCompile(`\b(nordic|NORDICSUBS)\b`, regexp2.IgnoreCase)
	// isMulti — must not match WEB-DL
	multiExp = regexp2.MustCompile(`(?<!(WEB-))\b(MULTi|DUAL|DL)\b`, regexp2.IgnoreCase)
)

// ParseLanguage extracts languages detected in a release title.
func ParseLanguage(title string, parsedTitle ...string) []Language {
	var pt string
	if len(parsedTitle) > 0 {
		pt = parsedTitle[0]
	} else {
		pt = ParseTitleAndYear(title).Title
	}

	languageTitle := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(title, ".", " "), pt, ""))

	var languages []Language
	add := func(l Language) { languages = append(languages, l) }

	if mustMatch(langEnglishExp, languageTitle) {
		add(LanguageEnglish)
	}
	if strings.Contains(languageTitle, "spanish") {
		add(LanguageSpanish)
	}
	if mustMatch(langDanishExp, languageTitle) {
		add(LanguageDanish)
	}
	if strings.Contains(languageTitle, "japanese") {
		add(LanguageJapanese)
	}
	if strings.Contains(languageTitle, "cantonese") {
		add(LanguageCantonese)
	}
	if strings.Contains(languageTitle, "mandarin") {
		add(LanguageMandarin)
	}
	if strings.Contains(languageTitle, "korean") {
		add(LanguageKorean)
	}
	if strings.Contains(languageTitle, "vietnamese") {
		add(LanguageVietnamese)
	}
	if mustMatch(langSwedishExp, languageTitle) {
		add(LanguageSwedish)
	}
	if strings.Contains(languageTitle, "finnish") {
		add(LanguageFinnish)
	}
	if strings.Contains(languageTitle, "turkish") {
		add(LanguageTurkish)
	}
	if strings.Contains(languageTitle, "portuguese") {
		add(LanguagePortuguese)
	}
	if strings.Contains(languageTitle, "hebrew") {
		add(LanguageHebrew)
	}
	if strings.Contains(languageTitle, "czech") {
		add(LanguageCzech)
	}
	if strings.Contains(languageTitle, "ukrainian") {
		add(LanguageUkrainian)
	}
	if strings.Contains(languageTitle, "catalan") {
		add(LanguageCatalan)
	}
	if strings.Contains(languageTitle, "estonian") {
		add(LanguageEstonian)
	}
	if mustMatch(langIcelandicExp, languageTitle) {
		add(LanguageIcelandic)
	}
	if mustMatch(langChineseExp, languageTitle) {
		add(LanguageChinese)
	}
	if strings.Contains(languageTitle, "thai") {
		add(LanguageThai)
	}
	if mustMatch(langItalianExp, languageTitle) {
		add(LanguageItalian)
	}
	if mustMatch(langGermanExp, languageTitle) {
		add(LanguageGerman)
	}
	if mustMatch(langFlemishExp, languageTitle) {
		add(LanguageFlemish)
	}
	if mustMatch(langGreekExp, languageTitle) {
		add(LanguageGreek)
	}
	if mustMatch(langFrenchExp, languageTitle) {
		add(LanguageFrench)
	}
	if mustMatch(langRussianExp, languageTitle) {
		add(LanguageRussian)
	}
	if mustMatch(langNorwegianExp, languageTitle) {
		add(LanguageNorwegian)
	}
	if mustMatch(langHungarianExp, languageTitle) {
		add(LanguageHungarian)
	}
	if mustMatch(langHebDubExp, languageTitle) {
		add(LanguageHebrew)
	}
	if mustMatch(langCzSkExp, languageTitle) {
		add(LanguageCzech)
	}
	if mustMatch(langUkrExp, languageTitle) {
		add(LanguageUkrainian)
	}
	if mustMatch(langPolishExp, languageTitle) {
		add(LanguagePolish)
	}
	if mustMatch(langDutchExp, languageTitle) {
		add(LanguageDutch)
	}
	if mustMatch(langHindiExp, languageTitle) {
		add(LanguageHindi)
	}
	if mustMatch(langTamilExp, languageTitle) {
		add(LanguageTamil)
	}
	if mustMatch(langArabicExp, languageTitle) {
		add(LanguageArabic)
	}
	if mustMatch(langLatvianExp, languageTitle) {
		add(LanguageLatvian)
	}
	if mustMatch(langLithuanianExp, languageTitle) {
		add(LanguageLithuanian)
	}
	if mustMatch(langRomanianExp, languageTitle) {
		add(LanguageRomanian)
	}
	if mustMatch(langSlovakExp, languageTitle) {
		add(LanguageSlovak)
	}
	if mustMatch(langBrazilianExp, languageTitle) {
		add(LanguageBrazilian)
	}
	if mustMatch(langPersianExp, languageTitle) {
		add(LanguagePersian)
	}
	if mustMatch(langBengaliExp, languageTitle) {
		add(LanguageBengali)
	}
	if mustMatch(langBulgarianExp, languageTitle) {
		add(LanguageBulgarian)
	}
	if mustMatch(langSerbianExp, languageTitle) {
		add(LanguageSerbian)
	}
	if mustMatch(langNordicExp, languageTitle) {
		add(LanguageNordic)
	}
	if IsMulti(languageTitle) {
		add(LanguageEnglish)
	}
	if len(languages) == 0 {
		add(LanguageEnglish)
	}

	return dedupeLanguages(languages)
}

func dedupeLanguages(langs []Language) []Language {
	seen := make(map[Language]bool)
	out := langs[:0]
	for _, l := range langs {
		if !seen[l] {
			seen[l] = true
			out = append(out, l)
		}
	}
	return out
}

// IsMulti returns true when the title contains a MULTI/DUAL/DL tag that
// isn't part of WEB-DL.
func IsMulti(title string) bool {
	noWeb, _ := regexp2.MustCompile(`\bWEB-?DL\b`, regexp2.IgnoreCase).Replace(title, "", -1, -1)
	ok, _ := multiExp.MatchString(noWeb)
	return ok
}
