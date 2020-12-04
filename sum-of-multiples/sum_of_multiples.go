package summultiples

//SumMultiples calculates sum of multiplies given a set of divisors
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	uniqueMultiplies := map[int]int{}

	for _, val := range divisors {
		for i := 1; i < limit; i++ {
			if i*val < limit {
				uniqueMultiplies[i*val] = i * val
			}
		}
	}

	for _, val := range uniqueMultiplies {
		sum += val
	}

	return sum
}
