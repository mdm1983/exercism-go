package romannumerals

import "errors"

type atr struct {
	a int
	r string
}

var conv = []atr{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

//ToRomanNumeral converts arabic numbers into equivalent roman representation
//max 3000 value allowed
func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("number out of range 1-3000")
	}
	var s string
	for _, v := range conv {
		for ; arabic >= v.a; arabic -= v.a {
			s = s + v.r
		}
	}
	return s, nil
}
