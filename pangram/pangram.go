package pangram

import "unicode"

//IsPangram checks if a string contains all the letters of the alhpabet
func IsPangram(s string) bool {

	alphabet := make(map[rune]bool)

	for _, val := range s {
		if unicode.IsLetter(val) {
			alphabet[unicode.ToLower(val)] = true
		}
	}

	return len(alphabet) == 26
}
