// Package raindrops offers the solution to the corresponding sound of a raindrop given a number.
package raindrops

import (
	"strconv"
)

var sounds = []struct {
	Number int
	Sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert gets the sound of the raindrop equivalent to the number.
func Convert(number int) (sound string) {
	for _, value := range sounds {
		if number%value.Number == 0 {
			sound += value.Sound
		}
	}

	if sound == "" {
		sound = strconv.Itoa(number)
	}

	return
}
