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
		{"D6 is between 1 and 6", 6, 1},
		{"D10 is between 1 and 10", 10, 1},
		{"D20 is between 1 and 20", 20, 1},
		{"D100 is between 1 and 100", 100, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for range [1000]int{} {
				got := dice.D(test.upper)
				if !(got <= test.upper) || !(got >= test.lower) {
					t.Errorf("D(%v) %v >= %v >= %v", test.upper, test.upper, got, test.lower)
				}
			}
		})
	}
}

func TestD2(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.D2()
		if !(got <= 2) || !(got >= 1) {
			t.Errorf("D2() got %v, want a number between 1 and 2", got)
		}
	}
}

func TestD6(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.D6()
		if !(got <= 6) || !(got >= 1) {
			t.Errorf("D6() got %v, want a number between 1 and 6", got)
		}
	}
}

func TestD10(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.D10()
		if !(got <= 10) || !(got >= 1) {
			t.Errorf("D10() got %v, want a number between 1 and 10", got)
		}
	}
}

func TestD20(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.D20()
		if !(got <= 20) || !(got >= 1) {
			t.Errorf("D20() got %v, want a number between 1 and 20", got)
		}
	}
}

func TestD100(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.D100()
		if !(got <= 100) || !(got >= 1) {
			t.Errorf("D100() got %v, want a number between 1 and 100", got)
		}
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
				if !(got <= test.upperBound) || !(got >= test.lowerBound) {
					t.Errorf("Roll(%v, %v) got %v, want between %v and %v",
						test.count, test.sides, got, test.lowerBound, test.upperBound)
				}
			}
		})
	}
}

func TestRoll2D6(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.Roll2D6()
		if !(got <= 12) || !(got >= 2) {
			t.Errorf("Roll2D6() got %v, want a number between 2 and 12", got)
		}
	}
}

func TestRoll3D6(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.Roll3D6()
		if !(got <= 18) || !(got >= 3) {
			t.Errorf("Roll3D6() got %v, want a number between 3 and 18", got)
		}
	}
}
