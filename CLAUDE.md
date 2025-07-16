# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Build and Install

```bash
# Build the binary in the current directory
make build

# Install the binary to $GOPATH/bin
make install

# Clean the test cache
make clean
```

### Testing

```bash
# Run all tests
make test

# Run a specific test
go test -v -run TestName

# Run tests with coverage
go test -cover ./...
```

### Linting

```bash
# Run linting checks
make lint
# or
golangci-lint run
```

The project uses `golangci-lint` with the following linters enabled:
- revive
- gofmt
- goimports

## Architecture Overview

The `gh-release-install` tool is a Go CLI application that installs executable binaries from GitHub release assets. It allows users to specify patterns to match release assets and install them as executable binaries.

### Core Components

1. **Main Application Flow**:
   - `main.go`: Entry point that handles CLI arguments using Cobra
   - `application.go`: Core application service that coordinates the overall process

2. **Asset Management**:
   - `asset.go`: Defines the Asset type and methods for extracting binary content
   - `asset_github.go`: Repository implementation for GitHub-hosted assets
   - `asset_external.go`: Repository implementation for non-GitHub hosted assets

3. **Binary Management**:
   - `exec_binary.go`: Interface for executable binary operations
   - `exec_binary_fs.go`: Filesystem implementation for writing binaries

4. **Pattern Matching**:
   - `pattern.go`: Handles pattern matching of asset URLs and binary name templates
   - Uses regular expressions to match asset URLs and templates to determine binary names

5. **Repository**:
   - `repository.go`: Defines repository types and parsing logic
   - `release.go`: Handles release information and semantic versioning

### Workflow

1. User provides a GitHub repository and release tag
2. The tool fetches assets from the GitHub release
3. It matches the assets against user-provided patterns
4. It downloads the matching asset
5. It extracts the executable binary from the asset (handling formats like tar, zip, gz)
6. It installs the binary to the specified directory

### Dependencies

- `github.com/spf13/cobra`: CLI framework
- `github.com/cli/go-gh`: GitHub CLI integration
- `github.com/google/go-github`: GitHub API client
- `github.com/gabriel-vasile/mimetype`: MIME type detection
- `github.com/cheggaaa/pb`: Progress bar display