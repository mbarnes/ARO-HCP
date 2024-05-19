package arm

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"maps"
	"time"
)

// Resource represents a basic ARM resource
type Resource struct {
	ID         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Type       string      `json:"type,omitempty"`
	SystemData *SystemData `json:"systemData,omitempty"`
}

func (src *Resource) Copy(dst *Resource) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.Type = src.Type
	if src.SystemData == nil {
		dst.SystemData = nil
	} else {
		dst.SystemData = &SystemData{}
		src.SystemData.Copy(dst.SystemData)
	}
}

// TrackedResource represents a tracked ARM resource
type TrackedResource struct {
	Resource
	Location string            `json:"location,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

func (src *TrackedResource) Copy(dst *TrackedResource) {
	src.Resource.Copy(&dst.Resource)
	dst.Location = src.Location
	dst.Tags = maps.Clone(src.Tags)
}

// CreatedByType is the type of identity that created (or modified) the resource
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// SystemData includes creation and modification metadata for resources
// See https://eng.ms/docs/products/arm/api_contracts/resourcesystemdata
type SystemData struct {
	CreatedBy          string        `json:"createdBy,omitempty"`
	CreatedByType      CreatedByType `json:"createdByType,omitempty"`
	CreatedAt          *time.Time    `json:"createdAt,omitempty"`
	LastModifiedBy     string        `json:"lastModifiedBy,omitempty"`
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	LastModifiedAt     *time.Time    `json:"lastModifiedAt,omitempty"`
}

func (src *SystemData) Copy(dst *SystemData) {
	dst.CreatedBy = src.CreatedBy
	dst.CreatedByType = src.CreatedByType
	dst.CreatedAt = src.CreatedAt
	dst.LastModifiedBy = src.LastModifiedBy
	dst.LastModifiedByType = src.LastModifiedByType
	dst.LastModifiedAt = src.LastModifiedAt
}

// ProvisioningState represents the provisioning state of an ARM resource
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)
