package etl

import "strings"

//Transform function
func Transform(oldData map[int][]string) map[string]int {

	var newSystem = map[string]int{}
	for key, value := range oldData {
		//Loop over slice of strings
		for _, letters := range value {
			newSystem[strings.ToLower(letters)] = key
		}
	}

	return newSystem
}
