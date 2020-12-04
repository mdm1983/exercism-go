package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//RunLengthEncode encodes the message input string
func RunLengthEncode(message string) string {

	if message == "" {
		return ""
	}

	var count int = 1
	var b strings.Builder

	for i := 1; i <= len(message)-1; i++ {
		if message[i-1] == message[i] {
			count++
		} else {
			if count == 1 {
				fmt.Fprintf(&b, "%s", string(message[i-1]))

			} else {
				fmt.Fprintf(&b, "%d%s", count, string(message[i-1]))
			}
			count = 1
		}

		if i == len(message)-1 {
			if count == 1 {
				fmt.Fprintf(&b, "%s", string(message[i]))
			} else {
				fmt.Fprintf(&b, "%d%s", count, string(message[i]))
			}
		}

	}
	return b.String()
}

//RunLengthDecode decode the string
func RunLengthDecode(message string) string {

	var b strings.Builder
	var decodedString strings.Builder

	var lastNumber bool = false

	for _, val := range message {

		if unicode.IsDigit(val) {
			fmt.Fprintf(&b, "%d", int(val-'0'))
			lastNumber = true
		} else {

			if lastNumber {
				num, _ := strconv.Atoi(b.String())
				for i := 0; i < num; i++ {
					fmt.Fprintf(&decodedString, "%s", string(val))
				}
				lastNumber = false
				b.Reset()
			} else {
				fmt.Fprintf(&decodedString, "%s", string(val))
			}
		}
	}
	return decodedString.String()
}
