package luhn

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var validID = regexp.MustCompile("[a-zA-Z]+")

//Valid returns luhn check of a string
func Valid(s string) bool {

	if len(s) == 1 || len(strings.Trim(s, " ")) == 1 {
		return false
	}

	if validID.MatchString(s) {
		return false
	}

	checksum := 0
	count := 0

	for len(s) > 0 {

		r, size := utf8.DecodeLastRuneInString(s)
		if unicode.IsPunct(r) {
			return false
		}

		if unicode.IsDigit(r) {
			num := int(r - '0')
			if count%2 == 1 {
				num = num * 2
				if num > 9 {
					num -= 9
				}
			}
			checksum += num
			count++
		}
		s = s[:len(s)-size]
	}

	return checksum%10 == 0

}
