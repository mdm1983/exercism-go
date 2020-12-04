//Package proverb generates proverbs
package proverb

import "fmt"

//Proverb dumps out proverbs according to string list.
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return []string{}
	}

	var r = make([]string, len(rhyme))

	for i := 0; i < len(rhyme)-1; i++ {
		r[i] = fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1])
	}

	r[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %v.", rhyme[0])

	return r
}
