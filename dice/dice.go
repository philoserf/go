// Package dice provides a few simple functions to return dice-like results
package dice

import (
	crand "crypto/rand"
	"math/rand"
)

func init() { //nolint
	bits := 64
	seed, _ := crand.Prime(crand.Reader, bits)
	rand.Seed(seed.Int64())
}

// D returns a random integer, either 0 or 1.
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
