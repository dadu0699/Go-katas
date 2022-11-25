// Package hamming offers the solution to the Hamming distance.
package hamming

import "errors"

// Distance calculates the Hamming distance between two strings of equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("the inputs have different lengths")
	}

	a2 := []rune(a)
	b2 := []rune(b)
	distance := 0

	for i := range a2 {
		if b2[i] != a2[i] {
			distance++
		}
	}

	return distance, nil
}
