#!/bin/bash

echo "Installing goimports..."
go install golang.org/x/tools/cmd/goimports@latest

echo "Setting up pre-commit hook..."
ln -s -f ../../hooks/pre-commit .git/hooks/pre-commit

echo "Installation complete."
