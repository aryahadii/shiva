package langdetect

import "testing"

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
