//Package twofer package generates the two fer string
package twofer

import "fmt"

// ShareWith function create the fer string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
