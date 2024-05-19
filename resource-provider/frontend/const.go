package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

type contextKey int

const (
	ContextKeyOriginalPath = iota
	ContextKeyBody
	ContextKeySystemData
)
