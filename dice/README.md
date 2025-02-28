# dice

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/philoserf/go/dice)
[![Go Report Card](https://goreportcard.com/badge/github.com/philoserf/go/dice)](https://goreportcard.com/report/github.com/philoserf/go/dice)

Package dice provides simple functions to return dice-like results for simulations and games.

## Functions

### Basic Dice

### func [D](https://github.com/philoserf/go/blob/main/dice/dice.go#L16)

`func D(sides int) int`

D returns a random integer between 1 and sides.

### func [D2](https://github.com/philoserf/go/blob/main/dice/dice.go#L21)

`func D2() int`

D2 returns a random integer between 1 and 2.

### func [D6](https://github.com/philoserf/go/blob/main/dice/dice.go#L26)

`func D6() int`

D6 returns a random integer between 1 and 6.

### func [D10](https://github.com/philoserf/go/blob/main/dice/dice.go#L31)

`func D10() int`

D10 returns a random integer between 1 and 10.

### func [D20](https://github.com/philoserf/go/blob/main/dice/dice.go#L36)

`func D20() int`

D20 returns a random integer between 1 and 20.

### func [D100](https://github.com/philoserf/go/blob/main/dice/dice.go#L41)

`func D100() int`

D100 returns a random integer between 1 and 100.

### Multiple Dice

### func [Roll](https://github.com/philoserf/go/blob/main/dice/dice.go#L46)

`func Roll(count, sides int) int`

Roll returns the sum of multiple dice with the given number of sides.

### func [Roll2D6](https://github.com/philoserf/go/blob/main/dice/dice.go#L55)

`func Roll2D6() int`

Roll2D6 returns the sum of 2 six-sided dice (2-12).

### func [Roll3D6](https://github.com/philoserf/go/blob/main/dice/dice.go#L60)

`func Roll3D6() int`

Roll3D6 returns the sum of 3 six-sided dice (3-18).
