// Package dice provides a few simple functions to return dice-like results
package dice

import (
	"fmt"
	"math/rand/v2"
)

// These constants define common dice parameters.
const (
	twoDice     = 2
	threeDice   = 3
	fourSides   = 4
	sixSides    = 6
	eightSides  = 8
	twelveSides = 12
)

// D returns a random integer between 1 and sides.
// Panics if sides is less than 1.
func D(sides int) int {
	if sides < 1 {
		panic(fmt.Sprintf("dice: sides must be positive, got %d", sides))
	}
	return rand.IntN(sides) + 1
}

// D2 returns a random integer between 1 and 2.
func D2() int {
	return D(twoDice)
}

// D4 returns a random integer between 1 and 4.
func D4() int {
	return D(fourSides)
}

// D6 returns a random integer between 1 and 6.
func D6() int {
	return D(sixSides)
}

// D8 returns a random integer between 1 and 8.
func D8() int {
	return D(eightSides)
}

// D10 returns a random integer between 1 and 10.
func D10() int {
	return D(10) //nolint
}

// D12 returns a random integer between 1 and 12.
func D12() int {
	return D(twelveSides)
}

// D20 returns a random integer between 1 and 20.
func D20() int {
	return D(20) //nolint
}

// D100 returns a random integer between 1 and 100.
func D100() int {
	return D(100) //nolint
}

// Roll returns the sum of multiple dice with the given number of sides.
// Panics if count or sides is less than 1.
func Roll(count, sides int) int {
	if count < 1 {
		panic(fmt.Sprintf("dice: count must be positive, got %d", count))
	}
	if sides < 1 {
		panic(fmt.Sprintf("dice: sides must be positive, got %d", sides))
	}
	total := 0
	for range count {
		total += D(sides)
	}

	return total
}

// Roll2D6 returns the sum of 2 six-sided dice (2-12).
func Roll2D6() int {
	return Roll(twoDice, sixSides)
}

// Roll3D6 returns the sum of 3 six-sided dice (3-18).
func Roll3D6() int {
	return Roll(threeDice, sixSides)
}
