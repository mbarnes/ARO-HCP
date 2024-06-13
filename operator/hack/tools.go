//go:build tools
// +build tools

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import _ "k8s.io/code-generator"
