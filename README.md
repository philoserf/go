# Go tools, toys, & libraries

A collection of Go packages for various purposes.

## Contents

- [dice](#dice) - Fast, simple dice rolling library for tabletop RPGs and games

## dice

[![Go Reference](https://pkg.go.dev/badge/github.com/philoserf/go/dice.svg)](https://pkg.go.dev/github.com/philoserf/go/dice)

A lightweight, high-performance dice rolling library using Go's modern `math/rand/v2` package.

### Features

- Zero allocations per roll
- Automatic seeding with cryptographically sound randomness
- Support for common dice types (d2, d4, d6, d8, d10, d12, d20, d100)
- Multi-dice rolling functions
- Input validation with clear panic messages
- 100% test coverage

### Installation

```bash
go get github.com/philoserf/go/dice
```

### Quick Start

```go
package main

import (
    "fmt"
    "github.com/philoserf/go/dice"
)

func main() {
    // Roll a single die
    result := dice.D20()
    fmt.Printf("Attack roll: %d\n", result)

    // Roll multiple dice
    damage := dice.Roll(3, 6)  // 3d6
    fmt.Printf("Damage: %d\n", damage)

    // Use convenience functions
    initiative := dice.D20()
    stats := dice.Roll3D6()
    fmt.Printf("Initiative: %d, Stats: %d\n", initiative, stats)
}
```

### API Reference

#### Single Die Functions

All single die functions return a random integer between 1 and the die size (inclusive).

```go
dice.D(sides int) int    // Custom die (panics if sides < 1)
dice.D2() int            // 1-2
dice.D4() int            // 1-4
dice.D6() int            // 1-6
dice.D8() int            // 1-8
dice.D10() int           // 1-10
dice.D12() int           // 1-12
dice.D20() int           // 1-20
dice.D100() int          // 1-100
```

#### Multi-Die Functions

Roll multiple dice and return the sum.

```go
dice.Roll(count, sides int) int  // Roll 'count' dice with 'sides' (panics if either < 1)
dice.Roll2D6() int               // 2-12
dice.Roll3D6() int               // 3-18
```

### Performance

Built on `math/rand/v2` with zero memory allocations per operation:

```text
BenchmarkD-10          267827300    4.400 ns/op    0 B/op    0 allocs/op
BenchmarkD6-10         271751463    4.446 ns/op    0 B/op    0 allocs/op
BenchmarkD20-10        269157292    4.462 ns/op    0 B/op    0 allocs/op
BenchmarkRoll-10        80306948   14.93 ns/op     0 B/op    0 allocs/op
BenchmarkRoll2D6-10    120718734   10.05 ns/op     0 B/op    0 allocs/op
BenchmarkRoll3D6-10     81663901   15.04 ns/op     0 B/op    0 allocs/op
```

Run benchmarks yourself:

```bash
task bench
# or
go test -bench=. -benchmem ./dice
```

### Examples

#### Character Creation (D&D Style)

```go
func rollAbilityScores() [6]int {
    var scores [6]int
    for i := range scores {
        scores[i] = dice.Roll3D6()
    }
    return scores
}
```

#### Combat Mechanics

```go
func attackRoll(bonus int) int {
    return dice.D20() + bonus
}

func rollDamage(diceCount, diceSides, modifier int) int {
    return dice.Roll(diceCount, diceSides) + modifier
}
```

#### Random Encounters

```go
func randomEncounter() string {
    roll := dice.D100()
    switch {
    case roll <= 20:
        return "No encounter"
    case roll <= 60:
        return "Minor encounter"
    case roll <= 90:
        return "Moderate encounter"
    default:
        return "Deadly encounter"
    }
}
```

### Development

```bash
# Run tests
task test

# Run tests with coverage
task cover

# Run benchmarks
task bench

# Format and lint
task fmt
task lint

# Run everything
task all
```

---

[home] | [philoserf.com] | [source]

[home]: https://philoserf.github.io/
[philoserf.com]: https://philoserf.com/
[source]: https://github.com/philoserf/go
