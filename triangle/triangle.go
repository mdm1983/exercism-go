// Package Triangle hepls udentifying triangle types
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ = iota // equilateral
	Iso = iota // isosceles
	Sca = iota // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT
	k = NaT // Init with Not a triangle
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return k
	}

	if a > 0 && b > 0 && c > 0 {
		if (a+b >= c) && (a+c >= b) && (b+c >= a) {
			if a == b && b == c {
				k = Equ
			} else if (a == b && c != a) || (a == c && b != a) || (b == c && a != c) {
				k = Iso
			} else {
				k = Sca
			}
		}
	}
	return k
}
