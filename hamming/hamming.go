//Package hamming calculates hamming's code for two strings
package hamming

import (
	"fmt"
	"unicode/utf8"
)

//Distance calculates Hamming Code
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, fmt.Errorf("hamming len must be the same and greather than zero - found %s, %s", a, b)
	}

	hamming := 0
	for len(a) > 0 {
		r, size := utf8.DecodeRuneInString(a)
		r2, size2 := utf8.DecodeRuneInString(b)

		if r != r2 {
			hamming++
		}

		a = a[size:]
		b = b[size2:]
	}

	return hamming, nil

}
