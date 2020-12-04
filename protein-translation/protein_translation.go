package protein

import (
	"errors"
	"fmt"
)

var codonMapper = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

var (
	ErrInvalidBase = errors.New("Invalid base")
	ErrStop        = errors.New("Stop error")
)

//FromCodon translates codon into protein
func FromCodon(s string) (string, error) {

	if codonMapper[s] == "" {
		fmt.Println("found invalid base")
		return "", ErrInvalidBase
	}

	if codonMapper[s] == "STOP" {
		return "", ErrStop
	}

	return codonMapper[s], nil
}

//FromRNA returns proteins given an RNA sequence
func FromRNA(s string) ([]string, error) {

	proteinSequence := []string{}
	for i := 0; i < len(s); i += 3 {
		codon := extractCodon(s, i)
		fmt.Println("Codon:", codon, i)
		if codon != "" {
			sequence, err := FromCodon(codon)
			if err != nil && err == ErrStop {
				break
			} else if err != nil {
				return proteinSequence, err
			}

			proteinSequence = append(proteinSequence, sequence)
		}
	}

	return proteinSequence, nil
}

func extractCodon(s string, start int) string {
	if s != "" && len(s) >= start+3 {
		return s[start : start+3]
	}
	return ""

}
