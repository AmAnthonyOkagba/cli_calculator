# Calculator (Go)

A small first Go project — a simple command-line calculator demonstrating basic Go project structure, argument parsing.

## Features
- Basic arithmetic: add, sub, mul, div
- CLI usage and/or interactive mode
- Small, testable codebase suitable for learning Go tooling

## Requirements
- Go 1.20+ installed (https://go.dev/dl/)
- Windows PowerShell or Command Prompt (examples below use PowerShell)

## Getting started

Clone or open this repository in VS Code (project path shown for this workspace):

Install dependencies (if any):
```powershell
go mod tidy
```

Build:
```powershell
go build -o calculator .
```

Run without building (useful during development):
```powershell
go run .
# or run a specific file
go run main.go
```
