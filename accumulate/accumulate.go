package accumulate

//Accumulate function applies a given converter to a slice of string
func Accumulate(given []string, converter func(string) string) []string {

	var ret = make([]string, len(given))

	for i, s := range given {
		ret[i] = converter(s)
	}
	return ret
}
