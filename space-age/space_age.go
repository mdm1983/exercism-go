package space

import "fmt"

//Planet type represents planet
type Planet string

const (
	earth   = 60 * 60 * 24 * 365.25
	mercury = 0.2408467 * earth
	venus   = 0.61519726 * earth
	mars    = 1.8808158 * earth
	jupiter = 11.862615 * earth
	saturn  = 29.447498 * earth
	uranus  = 84.016846 * earth
	neptune = 164.79132 * earth
)

//Age Function calculates Age within planets
func Age(s float64, p Planet) float64 {

	var age float64
	switch p {
	case "Earth":
		age = s / earth
	case "Mercury":
		age = s / mercury
	case "Venus":
		age = s / venus
	case "Mars":
		age = s / mars
	case "Jupiter":
		age = s / jupiter
	case "Saturn":
		age = s / saturn
	case "Uranus":
		age = s / uranus
	case "Neptune":
		age = s / neptune
	}

	fmt.Println("Seconds", s, "on planet", p, "are equivalent to", age, "ages")
	return age
}
