//Package strand manipualtes DNA sequences
package strand

import (
	"strings"
)

//ToRNA calculates RNA from DNA
func ToRNA(dna string) string {
	rna := []string{}
	if dna != "" {
		var replacement = ""
		for _, c := range dna {
			switch c {
			case rune('C'):
				replacement = "G"
			case rune('G'):
				replacement = "C"
			case rune('A'):
				replacement = "U"
			case rune('T'):
				replacement = "A"

			}
			rna = append(rna, replacement)
		}
	}
	return strings.Join(rna, "")
}
