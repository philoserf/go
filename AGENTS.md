# AGENTS.md

Go monorepo with `dice` - high-performance dice rolling library for tabletop RPGs built on `math/rand/v2` with zero allocations per roll.

- **Module Path**: `github.com/philoserf/go`
- **Go Version**: 1.25+
- **Setup**: `brew bundle` or `task setup` to install gofumpt, golangci-lint, go-task, markdownlint-cli

## Development Commands

Use [Task](https://taskfile.dev):

- `task build` - Build all packages
- `task test` - Run all tests with race detection
- `task test:one TEST=TestName` - Run specific test
- `task cover` - Coverage report (opens HTML)
- `task bench` - Benchmarks with memory stats
- `task fmt` - Format with gofumpt
- `task lint` - Lint Go and Markdown
- `task lint:go` - Go linting with auto-fix
- `task lint:md` - Markdown linting
- `task clean` - Remove build artifacts
- `task all` - Everything: setup, fmt, lint, build, test

## Code Style

- Tests use separate `_test` package (e.g., `package dice_test`)
- All tests use `t.Parallel()` for concurrent execution
- Table-driven tests for repetitive cases
- Test both valid and invalid inputs (panic cases)
- Loop tests 1000 times to verify randomness
- Test panic behavior with deferred recovery
- Format with gofumpt (stricter than gofmt)
- Use named constants for magic numbers
- PascalCase for exported identifiers, camelCase for internal
- Panic on invalid inputs with message: `"dice: <field> must be <constraint>, got <value>"`

## dice Package API

- **Single die**: `D(sides)`, `D2()`, `D4()`, `D6()`, `D8()`, `D10()`, `D12()`, `D20()`, `D100()`
- **Multiple dice**: `Roll(count, sides)`, `Roll2D6()`, `Roll3D6()`
- All validate inputs and panic on values < 1
- Zero allocations per operation
- Cryptographic randomness from `math/rand/v2`
