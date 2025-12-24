# AGENTS.md

Guide for AI agents working in this repository.

## Project Overview

**webedex** - Declarative browser bookmark configuration tool written in Go.

- **Language**: Go 1.25.5
- **Module**: `github.com/AWDDude/webedex`
- **Status**: Early development (scaffolding stage)

## Commands

### Build

```bash
CGO_ENABLED=0 go build ./...
```

### Run

```bash
CGO_ENABLED=0 go run ./cmd/main.go
```

### Test

```bash
go test ./...
```

### Format & Lint

```bash
go fmt ./...
go vet ./...
```

## Project Structure

```
webedex/
├── cmd/
│   └── main.go      # Application entry point
├── go.mod           # Go module definition
├── .gitignore
├── LICENSE
└── README.md
```

## Code Conventions

### Go Standards

- Follow standard Go conventions and idioms
- Use `gofmt` for formatting
- Package names should be lowercase, single-word
- Entry point lives in `cmd/` directory

### Naming

- Use camelCase for unexported identifiers
- Use PascalCase for exported identifiers
- Prefer short, descriptive names

## Development Notes

- This is a new project with minimal code - patterns will evolve
- No external dependencies yet (only standard library)
- No CI/CD pipeline configured
- No Makefile - use standard `go` commands

## Environment

- `.env` files are gitignored - use for local configuration
- Binary artifacts (`*.exe`, `*.dll`, `*.so`, `*.dylib`) are gitignored
