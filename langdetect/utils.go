package langdetect

import "strconv"

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
				nonLatinText += strconv.QuoteRune(runeVal)
			}
		}
		return nonLatinText
	}

	return text
}
