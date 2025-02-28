// Package dice provides a few simple functions to return dice-like results
package dice

import (
	"math/rand"
	"time"
)

// These constants define common dice parameters.
const (
	twoDice   = 2
	threeDice = 3
	sixSides  = 6
)

// This avoids using init() as flagged by linters.
func getRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano())) //nolint
}

// D returns a random integer between 1 and sides.
func D(sides int) int {
	return getRand().Intn(sides) + 1
}

// D2 returns a random integer between 1 and 2.
func D2() int {
	return D(twoDice)
}

// D6 returns a random integer between 1 and 6.
func D6() int {
	return D(sixSides)
}

// D10 returns a random integer between 1 and 10.
func D10() int {
	return D(10) //nolint
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
func Roll(count, sides int) int {
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
