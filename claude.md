# Claude Helper Guide for philoserf/go

## Task Commands (using Task)

- Build project: `task build`
- Run all tests: `task test`
- Run specific test: `task test:one TEST=TestName`
- Run with coverage: `task cover`
- Format code: `task fmt`
- Lint all files: `task lint`
- Lint Go code: `task lint:go`
- Lint Markdown: `task lint:md`
- Format Markdown: `task pretty`
- Run all tasks: `task all`
- Install dependencies: `task setup`

## Go Commands (without Task)

- Build project: `go build ./...`
- Run all tests: `go test ./...`
- Run specific test: `go test -run TestName ./package`
- Race detection: `go test -race ./...`
- Format code: `gofumpt -l -w .`
- Lint check: `golangci-lint run --fix`

## Code Style Guidelines

- Package doc comments start above package declaration
- Function doc comments use `// FuncName ...` format
- Test files use separate test package with `_test` suffix
- Tests use `t.Parallel()` for concurrent execution
- Group imports with standard library first
- Use table-driven tests for repetitive tests
- Mark lint exceptions with `//nolint:rule` comments
- Follow Go standard naming conventions (PascalCase for exported, camelCase for internal)
- Prefer early returns for error handling
- Test error cases explicitly
