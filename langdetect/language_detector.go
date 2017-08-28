package langdetect

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/aryahadii/ngram"
	"github.com/aryahadii/shiva/model"
)

var (
	randomGeneratorSeed = 75
	alphaWidth          = 0.05
	alphaDefault        = 0.5
	trials              = 7
	convThreshold       = 0.9999
	iterationsLimit     = 1000
	baseFreq            = 10000
)

type LanguageDetector struct {
	WordProbMap map[string][]int
	Languages   []string
}

// New creates new instance of LanguageDetector
func New(profilesDirectory string) *LanguageDetector {
	wordProbMap, langs, err := initWordProbMap(profilesDirectory)
	if err != nil {
		// return nil
		fmt.Println(err)
	}

	return &LanguageDetector{
		WordProbMap: wordProbMap,
		Languages:   langs,
	}
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
	probList := make([]model.LanguageProbablity, len(ld.Languages))
	for index := range probList {
		probList[index].Code = model.LanguageCode(ld.Languages[index])
	}

	text = cleanText(text)

	// Generate ngrams (n=2,3)
	var ngramSlice []string
	for n := 2; n < 4; n++ {
		ngramTemp, err := ngram.Get(text, n)
		if err != nil {
			return probList, err
		}
		ngramSlice = append(ngramSlice, ngramTemp...)
	}
	if len(ngramSlice) == 0 {
		return probList, errors.New("Text doesn't contain enough features")
	}

	// Update probabilities
	randomGenerator := rand.New(rand.NewSource(75))
	for trial := 0; trial < trials; trial++ {
		prob := make([]float64, len(ld.Languages))
		for index := range prob {
			prob[index] = 1.0 / float64(len(ld.Languages))
		}
		trialAlpha := alphaDefault + (randomGenerator.NormFloat64()/math.MaxFloat64)*alphaWidth

		for iter := 0; ; iter++ {
			wordIndex := randomGenerator.Intn(len(ngramSlice))
			ld.updateProbs(prob, ngramSlice[wordIndex], trialAlpha)
			if iter%5 == 0 {
				if maxProbability(prob) >= convThreshold || iter >= iterationsLimit {
					break
				}
			}
		}
		for index := range prob {
			probList[index].Probability += prob[index] / float64(trials)
		}
	}

	sortProbability(probList)
	return probList, nil
}

func (ld *LanguageDetector) updateProbs(prob []float64, word string, alpha float64) {
	if len(word) == 0 {
		return
	}
	if wordProb, ok := ld.WordProbMap[word]; ok {
		weight := alpha / float64(baseFreq)
		for index := range prob {
			prob[index] *= weight + float64(wordProb[index])
		}
	}
}
