package anagram

import (
	"strings"
	"unicode"
)

//Detect identifies anagram of a string over a list of candidates
func Detect(subject string, candidates []string) []string {

	anagrams := []string{}

	for _, val := range candidates {
		if checkAnagram(subject, val) {
			anagrams = append(anagrams, val)
		}
	}

	return anagrams

}

func checkAnagram(subject string, candidate string) bool {

	var composition = make(map[rune]int)

	if strings.ToLower(subject) == strings.ToLower(candidate) {
		return false
	}

	for _, val := range subject {
		composition[unicode.ToLower(val)] = composition[unicode.ToLower(val)] + 1
	}

	for _, val := range candidate {
		composition[unicode.ToLower(val)] = composition[unicode.ToLower(val)] - 1
	}

	for _, val := range composition {
		if val != 0 {
			return false
		}
	}

	return true

}
