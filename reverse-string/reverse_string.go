package reverse

//Reverse the string
func Reverse(s string) string {

	stringRunes := []rune(s)
	for i, j := 0, len(stringRunes)-1; i < j; i, j = i+1, j-1 {
		stringRunes[i], stringRunes[j] = stringRunes[j], stringRunes[i]
	}

	return string(stringRunes)
}
