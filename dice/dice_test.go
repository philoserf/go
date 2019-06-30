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
		{"D6 is between 1 and 6", 6, 1},
		{"D10 is between 1 and 10", 10, 1},
		{"D100 is between 1 and 100", 100, 1},
	}
	for _, test := range tests {
		test := test
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
