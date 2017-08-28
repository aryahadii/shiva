package model

// LanguageCode is integer representation of language name
type LanguageCode int

const (
	UnknownLanguageCode LanguageCode = 0
	PersianLanguageCode LanguageCode = 1
	EnglishLanguageCode LanguageCode = 2
)

// LanguageProbablity is a struct which contains language code and it's prob
type LanguageProbablity struct {
	Code        LanguageCode
	Probability float64
}
