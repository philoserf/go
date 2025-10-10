package dice_test

import (
	"testing"

	"github.com/philoserf/go/dice" //nolint:depguard
)

func TestD(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		upper int
		lower int
	}{
		{"D2 is between 1 and 2", 2, 1},
		{"D4 is between 1 and 4", 4, 1},
		{"D6 is between 1 and 6", 6, 1},
		{"D8 is between 1 and 8", 8, 1},
		{"D10 is between 1 and 10", 10, 1},
		{"D12 is between 1 and 12", 12, 1},
		{"D20 is between 1 and 20", 20, 1},
		{"D100 is between 1 and 100", 100, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for range [1000]int{} {
				got := dice.D(test.upper)
				if (got > test.upper) || (got < test.lower) {
					t.Errorf("D(%v) %v >= %v >= %v", test.upper, test.upper, got, test.lower)
				}
			}
		})
	}
}

func TestDInvalidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		sides int
	}{
		{"zero sides panics", 0},
		{"negative sides panics", -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r == nil {
					t.Errorf("D(%v) did not panic", test.sides)
				}
			}()

			dice.D(test.sides)
		})
	}
}

func TestSpecificDice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		fn    func() int
		upper int
		lower int
	}{
		{"D2", dice.D2, 2, 1},
		{"D4", dice.D4, 4, 1},
		{"D6", dice.D6, 6, 1},
		{"D8", dice.D8, 8, 1},
		{"D10", dice.D10, 10, 1},
		{"D12", dice.D12, 12, 1},
		{"D20", dice.D20, 20, 1},
		{"D100", dice.D100, 100, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for range [1000]int{} {
				got := test.fn()
				if (got > test.upper) || (got < test.lower) {
					t.Errorf("%s() got %v, want a number between %v and %v",
						test.name, got, test.lower, test.upper)
				}
			}
		})
	}
}

func TestRoll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		count      int
		sides      int
		lowerBound int
		upperBound int
	}{
		{"2d6 is between 2 and 12", 2, 6, 2, 12},
		{"3d6 is between 3 and 18", 3, 6, 3, 18},
		{"4d4 is between 4 and 16", 4, 4, 4, 16},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for range [1000]int{} {
				got := dice.Roll(test.count, test.sides)
				if (got > test.upperBound) || (got < test.lowerBound) {
					t.Errorf("Roll(%v, %v) got %v, want between %v and %v",
						test.count, test.sides, got, test.lowerBound, test.upperBound)
				}
			}
		})
	}
}

func TestRollInvalidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		count int
		sides int
	}{
		{"zero count panics", 0, 6},
		{"negative count panics", -1, 6},
		{"zero sides panics", 2, 0},
		{"negative sides panics", 2, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Roll(%v, %v) did not panic", test.count, test.sides)
				}
			}()

			dice.Roll(test.count, test.sides)
		})
	}
}

func TestSpecificRolls(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		fn         func() int
		lowerBound int
		upperBound int
	}{
		{"Roll2D6", dice.Roll2D6, 2, 12},
		{"Roll3D6", dice.Roll3D6, 3, 18},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for range [1000]int{} {
				got := test.fn()
				if (got > test.upperBound) || (got < test.lowerBound) {
					t.Errorf("%s() got %v, want a number between %v and %v",
						test.name, got, test.lowerBound, test.upperBound)
				}
			}
		})
	}
}
