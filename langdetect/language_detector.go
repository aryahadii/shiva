package langdetect

import (
	"errors"

	"github.com/aryahadii/ngram"
	"github.com/aryahadii/shiva/model"
)

type LanguageDetector struct {
}

// New creates new instance of LanguageDetector
func New() *LanguageDetector {
	return nil
}

// Detect detects language of text
func (ld *LanguageDetector) Detect(text string) (model.LanguageCode, error) {
	probs, err := ld.DetectByProbability(text)
	if err != nil {
		return model.UnknownLanguageCode, err
	}
	if len(probs) != 0 {
		return probs[0].Code, nil
	}

	return model.UnknownLanguageCode, nil
}

// DetectByProbability detects language of text also return probability of each languages
func (ld *LanguageDetector) DetectByProbability(text string) ([]model.LanguageProbablity, error) {
	probList := []model.LanguageProbablity{}
	text = cleanText(text)

	ngramSlice, err := ngram.Get(text, 3)
	if err != nil {
		return probList, err
	}
	if len(ngramSlice) == 0 {
		return probList, errors.New("Text doesn't contain enough features")
	}

	return probList, err
}
