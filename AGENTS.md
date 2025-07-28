# Agent Guidelines

## Build/Test Commands
- `go test ./...` - Run all tests
- `go test ./api/test/...` - Run specific test package
- `go build ./...` - Build all packages
- `go fmt ./...` - Format code
- `go vet ./...` - Run static analysis

## Code Style
- Use standard Go formatting (`go fmt`)
- Package names: lowercase, no underscores
- Function names: PascalCase for exported, camelCase for unexported
- Error handling: Always check errors, use `fmt.Errorf` with context
- Imports: stdlib first, then third-party, then local
- HTTP handlers: Use `func(w http.ResponseWriter, r *http.Request)` signature
- JSON responses: Use proper JSON format (fix malformed responses)
- No global state in handlers - use constants or dependency injection