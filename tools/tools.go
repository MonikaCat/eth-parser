//go:build tools
// +build tools

// This file uses the recommended method for tracking developer tools in a Go
// module. It does not define any buildable source code, but instead lists tools
// REFERENCE: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)