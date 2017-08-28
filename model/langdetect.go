package model

// LanguageCode is integer representation of language name
type LanguageCode string

const (
	UnknownLanguageCode LanguageCode = "unknown"
)

// LanguageProbablity is a struct which contains language code and it's prob
type LanguageProbablity struct {
	Code        LanguageCode
	Probability float64
}

type ByProbability []LanguageProbablity

func (p ByProbability) Len() int {
	return len(p)
}
func (p ByProbability) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p ByProbability) Less(i, j int) bool {
	return p[i].Probability < p[j].Probability
}
