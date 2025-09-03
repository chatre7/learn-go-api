#!/bin/bash

# Run tests with coverage
echo "Running tests with coverage..."
go test -coverprofile=coverage.out ./tests/... ./internal/...

# Display coverage
echo "Coverage report:"
go tool cover -func=coverage.out

# Optionally open HTML coverage report
echo "Generating HTML coverage report..."
go tool cover -html=coverage.out -o coverage.html

echo "Coverage report saved to coverage.html"
echo "You can open it with: open coverage.html (on macOS) or xdg-open coverage.html (on Linux)"