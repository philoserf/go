package dice_test

import (
	"testing"

	"github.com/philoserf/go/dice"
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

func TestD100(t *testing.T) {
	t.Parallel()

	for range [1000]int{} {
		got := dice.D100()
		if !(got <= 100) || !(got >= 1) {
			t.Errorf("D100() got %v, want a number between 1 and 100", got)
		}
	}
}
