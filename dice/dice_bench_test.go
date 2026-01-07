package dice_test

import (
	"testing"

	"github.com/philoserf/go/dice" //nolint:depguard
)

func BenchmarkD(b *testing.B) {
	for b.Loop() {
		dice.D(20)
	}
}

func BenchmarkD6(b *testing.B) {
	for b.Loop() {
		dice.D6()
	}
}

func BenchmarkD20(b *testing.B) {
	for b.Loop() {
		dice.D20()
	}
}

func BenchmarkRoll(b *testing.B) {
	for b.Loop() {
		dice.Roll(3, 6)
	}
}

func BenchmarkRoll2D6(b *testing.B) {
	for b.Loop() {
		dice.Roll2D6()
	}
}

func BenchmarkRoll3D6(b *testing.B) {
	for b.Loop() {
		dice.Roll3D6()
	}
}
