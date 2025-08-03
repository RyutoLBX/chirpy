#!/bin/bash
set -e

echo "Building chirpy..."
go build -o ./bin/chirpy ./cmd/main.go
echo "✅ Done!"
echo
