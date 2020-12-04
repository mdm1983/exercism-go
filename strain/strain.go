package strain

//wrapper for slice of int
type Ints []int

//wrapper for array of slice of int
type Lists [][]int

//wrapper for slice of strings
type Strings []string

//Keep (Ints) builds a collection of Ints given the funcion predicate
func (i Ints) Keep(predicate func(int) bool) Ints {

	var results Ints
	for _, v := range i {
		if predicate(v) == true {
			results = append(results, v)
		}
	}

	return results
}

//Discard (Ints) builds a collection of Ints given the funcion predicate
func (i Ints) Discard(predicate func(int) bool) Ints {

	var results Ints
	for _, v := range i {
		if predicate(v) == false {
			results = append(results, v)
		}
	}

	return results
}
func (l Lists) Keep(predicate func([]int) bool) Lists {

	var results Lists
	for _, v := range l {
		if predicate(v) == true {
			results = append(results, v)
		}
	}

	return results
}

func (s Strings) Keep(predicate func(string) bool) Strings {
	var results Strings
	for _, v := range s {
		if predicate(v) == true {
			results = append(results, v)
		}
	}

	return results
}
