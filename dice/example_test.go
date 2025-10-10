package dice_test

import (
	"fmt"

	"github.com/philoserf/go/dice" //nolint:depguard
)

func ExampleD() {
	// Roll a 20-sided die
	result := dice.D(20)
	fmt.Printf("Rolled a d20: %d\n", result)
}

func ExampleD6() {
	// Roll a standard 6-sided die
	result := dice.D6()
	fmt.Printf("Rolled: %d\n", result)
}

func ExampleD20() {
	// Roll a 20-sided die for an attack roll
	result := dice.D20()
	fmt.Printf("Attack roll: %d\n", result)
}

func ExampleRoll() {
	// Roll 3 six-sided dice for ability score generation
	result := dice.Roll(3, 6)
	fmt.Printf("Ability score: %d\n", result)
}

func ExampleRoll2D6() {
	// Roll 2d6 for a common RPG mechanic
	result := dice.Roll2D6()
	fmt.Printf("2d6 roll: %d\n", result)
}

func ExampleRoll3D6() {
	// Roll 3d6 for ability score generation
	result := dice.Roll3D6()
	fmt.Printf("Ability score: %d\n", result)
}
