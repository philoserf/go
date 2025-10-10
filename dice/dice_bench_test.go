package dice_test

import (
	"testing"

	"github.com/philoserf/go/dice" //nolint:depguard
)

func BenchmarkD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.D(20)
	}
}

func BenchmarkD6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.D6()
	}
}

func BenchmarkD20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.D20()
	}
}

func BenchmarkRoll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.Roll(3, 6)
	}
}

func BenchmarkRoll2D6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.Roll2D6()
	}
}

func BenchmarkRoll3D6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dice.Roll3D6()
	}
}
