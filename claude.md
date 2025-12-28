# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go monorepo containing tools, toys, and libraries. The primary package is `dice`, a high-performance dice rolling library for tabletop RPGs and games built on `math/rand/v2` with zero allocations per roll.

- **Module Path**: `github.com/philoserf/go`
- **Go Version**: 1.25+
- **Dependencies**: Managed via Homebrew using a Brewfile

Run `brew bundle` or `task setup` to install all required development tools (gofumpt, golangci-lint, go-task, markdownlint-cli).

## Development Commands

### Task (Preferred)

The project uses [Task](https://taskfile.dev) for build automation:

- `task build` - Build all packages
- `task test` - Run all tests with race detection
- `task test:one TEST=TestName` - Run a specific test
- `task cover` - Run tests with coverage report (opens HTML)
- `task bench` - Run benchmarks with memory allocation stats
- `task fmt` - Format Go code using gofumpt
- `task lint` - Run all linters (Go + Markdown)
- `task lint:go` - Lint Go code using golangci-lint with auto-fix
- `task lint:md` - Lint Markdown files using markdownlint-cli
- `task pretty` - Format Markdown files using prettier
- `task clean` - Remove build artifacts (coverage.out)
- `task all` - Run everything (setup, fmt, lint, build, test)
- `task setup` - Install development dependencies via Brewfile (gofumpt, golangci-lint, go-task, markdownlint-cli)

### Direct Go Commands

- `go build ./...` - Build all packages
- `go test ./...` - Run all tests
- `go test -race ./...` - Run tests with race detection
- `go test -run TestName ./package` - Run specific test
- `go test -bench=. -benchmem ./...` - Run benchmarks
- `gofumpt -l -w .` - Format code
- `golangci-lint run --fix` - Lint and fix issues

## Code Style Guidelines

### Package Structure

- Tests use separate test package with `_test` suffix (e.g., `package dice_test`)
- Tests import the package under test explicitly
- Package documentation comments go above the package declaration

### Testing Conventions

- All tests use `t.Parallel()` for concurrent execution
- Use table-driven tests for repetitive test cases
- Test both valid and invalid inputs (panic cases)
- Run tests 1000 times in loops to verify randomness behavior
- Test error cases explicitly with deferred panic recovery

### Code Formatting

- Use `gofumpt` (stricter than `gofmt`) for formatting
- Group imports with standard library first
- Use named constants for magic numbers (e.g., `sixSides = 6`)
- Mark intentional lint exceptions with `//nolint:rule` comments

### Naming Conventions

- PascalCase for exported identifiers
- camelCase for internal identifiers
- Descriptive test names: `TestFunctionName`, `TestFunctionNameInvalidInput`

### Error Handling

- Prefer early returns for error handling
- Use `panic()` for invalid inputs with descriptive messages (dice library convention)
- Format panic messages consistently: `"dice: <field> must be <constraint>, got <value>"`

## Architecture Notes

### dice Package

The dice package provides a simple API built on `math/rand/v2`:

- **Single die functions**: `D(sides)`, `D2()`, `D4()`, `D6()`, `D8()`, `D10()`, `D12()`, `D20()`, `D100()`
- **Multi-die functions**: `Roll(count, sides)`, `Roll2D6()`, `Roll3D6()`
- All functions validate inputs and panic on invalid values (< 1)
- Zero allocations per operation for maximum performance
- Automatic seeding with cryptographically sound randomness from `math/rand/v2`

### Testing Strategy

- Validate output ranges (1000 iterations per test)
- Test panic behavior for invalid inputs using deferred recovery
- Benchmark all functions with memory allocation tracking
- Maintain 100% test coverage
