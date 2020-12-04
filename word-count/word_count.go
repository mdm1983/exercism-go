package wordcount

import (
	"regexp"
	"strings"
)

//Frequency defines word counting object
type Frequency map[string]int

//WordCount counts the number of words given a phrase
func WordCount(phrase string) Frequency {

	var wordCount = make(Frequency)

	var validID = regexp.MustCompile("[a-zA-Z0-9]+['][a-zA-Z0-9]+|[a-zA-Z]+|[0-9]+")

	for _, val := range validID.FindAllString(phrase, -1) {
		wordCount[strings.ToLower(val)]++
	}

	return wordCount
}
