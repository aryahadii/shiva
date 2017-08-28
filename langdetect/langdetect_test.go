package langdetect

import (
	"testing"

	"github.com/aryahadii/shiva/model"
)

func TestInitWordProbMap(t *testing.T) {
	wordProb, langs, err := initWordProbMap("../langdetect_profiles")
	if err != nil {
		t.Errorf("Error! %v\n", err)
	}
	if len(langs) == 0 {
		t.Errorf("Languages length is zero\n")
	}
	if len(wordProb) == 0 {
		t.Errorf("Wordsprob is empty\n")
	}
}

func TestDetectByProbability(t *testing.T) {
	languageDetector := New("../langdetect_profiles")

	result, err := languageDetector.Detect("This text is for test")
	if err != nil {
		t.Errorf("Error occured: %v\n", err)
	}
	if result != model.LanguageCode("fa") {
		t.Errorf("Wrong language detection!\n")
	}
}
