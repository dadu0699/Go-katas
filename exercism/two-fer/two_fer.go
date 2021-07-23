// Package twofer offers a funny sharing scheme.
package twofer

import "fmt"

// ShareWith tells you how to share with the given name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
