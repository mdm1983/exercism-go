package raindrops

import (
	"strconv"
)

//Convert function calculates the factors of a number
func Convert(number int) string {

	var factors = ""

	//check factor 3
	if number%3 == 0 {
		//Append the correct sound
		factors += "Pling"
	}

	//check factor 5
	if number%5 == 0 {
		//Append the correct sound
		factors += "Plang"
	}

	//check factor 5
	if number%7 == 0 {
		//Append the correct sound
		factors += "Plong"
	}

	if factors == "" {
		factors += strconv.Itoa(number)
	}

	return factors

}
