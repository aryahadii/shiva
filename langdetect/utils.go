package langdetect

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/aryahadii/shiva/model"
)

func isLatinRune(r rune) bool {
	if r >= 'A' && r <= 'z' {
		return true
	}
	return false
}

func isLatinExtendedAdditionalRune(r rune) bool {
	if r >= 'Ḁ' && r <= 'ỿ' {
		return true
	}
	return false
}

// cleanText removes URLs, e-mail address and Latin sentence if it is not written in Latin alphabet
func cleanText(text string) string {
	var latinCharsCount, nonLatinCharsCount int
	for _, runeVal := range text {
		if isLatinRune(runeVal) {
			latinCharsCount++
		} else if runeVal >= '\u0300' && !isLatinExtendedAdditionalRune(runeVal) {
			nonLatinCharsCount++
		}
	}

	// If text isn't latin
	if latinCharsCount*2 < nonLatinCharsCount {
		var nonLatinText string
		for _, runeVal := range text {
			if !isLatinRune(runeVal) {
				nonLatinText += string(runeVal)
			}
		}
		return nonLatinText
	}

	return text
}

type profile struct {
	Frequencies map[string]int `json:"freq"`
}

func initWordProbMap(profilesDir string) (wordProb map[string][]int, langs []string, err error) {
	langs = []string{}
	wordProb = map[string][]int{}

	// Get all files in profilesDir
	files, err := ioutil.ReadDir(profilesDir)
	if err != nil {
		return wordProb, langs, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Load profile
		prof := &profile{}
		fileBytes, err := ioutil.ReadFile(filepath.Join(profilesDir, file.Name()))
		if err != nil {
			return wordProb, langs, err
		}
		if err := json.Unmarshal(fileBytes, &prof); err != nil {
			return wordProb, langs, err
		}

		// Add new language
		langs = append(langs, file.Name())
		for key := range wordProb {
			wordProb[key] = append(wordProb[key], 0)
		}
		// Add profile to wordProb
		for str, freq := range prof.Frequencies {
			if _, ok := wordProb[str]; !ok {
				wordProb[str] = make([]int, len(langs))
			}
			wordProb[str][len(wordProb[str])-1] = freq
		}
	}

	return wordProb, langs, nil
}

func maxProbability(prob []float64) float64 {
	var sum, max float64
	for _, probability := range prob {
		sum += probability
		if probability > max {
			max = probability
		}
	}

	if sum == 0 {
		return 0
	}
	return max / sum
}

func sortProbability(langs []model.LanguageProbablity) {
	sort.Sort(sort.Reverse(model.ByProbability(langs)))
}
