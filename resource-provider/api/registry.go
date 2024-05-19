package api

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

type VersionedHCPOpenShiftCluster interface {
	Normalize(*HCPOpenShiftCluster)
	ValidateStatic() error
}

type Version interface {
	NewHCPOpenShiftCluster(*HCPOpenShiftCluster) VersionedHCPOpenShiftCluster
}

// APIs is the map of registered API versions
var APIs = map[string]Version{}
