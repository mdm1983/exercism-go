package isogram

import (
	"unicode"
)

//IsIsogram calculates if a word is an isogram (no repeating letters)
func IsIsogram(s string) bool {
	var count = map[rune]bool{}

	for _, value := range s {
		if value != '-' && value != ' ' {
			if count[unicode.ToLower(value)] {
				return false
			}

			count[unicode.ToLower(value)] = true
		}
	}
	return true
}
