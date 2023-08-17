// Package dice provides a few simple functions to return dice-like results
package dice

import (
	"math/rand"
)

// D returns a random integer between 1 and sides.
func D(sides int) int {
	return rand.Intn(sides) + 1 //nolint
}

// D2 returns a random integer either 0 or 1.
func D2() int {
	return D(2) //nolint
}

// D6 returns a random integer between 1 and 6.
func D6() int {
	return D(6) //nolint
}

// D10 returns a random integer between 1 and 10.
func D10() int {
	return D(10) //nolint
}

// D100 returns a random integer between 1 and 100.
func D100() int {
	return D(100) //nolint
}
