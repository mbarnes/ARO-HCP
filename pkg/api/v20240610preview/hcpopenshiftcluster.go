package v20240610preview

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net"

	"github.com/Azure/ARO-HCP/pkg/api/arm"
	"github.com/Azure/ARO-HCP/pkg/api/json"
)

// HCPOpenShiftCluster represents an ARO HCP OpenShift cluster resource.
type HCPOpenShiftCluster struct {
	arm.TrackedResource
	Properties HCPOpenShiftClusterProperties `json:"properties,omitempty"`
}

// HCPOpenShiftClusterProperties represents the property bag of a HCPOpenShiftCluster resource.
type HCPOpenShiftClusterProperties struct {
	ProvisioningState arm.ProvisioningState `json:"provisioningState,omitempty"       visibility:"read"`
	ClusterProfile    ClusterProfile        `json:"clusterProfile,omitempty"          visibility:"read,create,update"`
	APIProfile        APIProfile            `json:"apiProfile,omitempty"              visibility:"read,create"`
	ConsoleProfile    ConsoleProfile        `json:"consoleProfile,omitempty"          visibility:"read,create"`
	IngressProfiles   []IngressProfile      `json:"ingressProfiles,omitempty"         visibility:"read,create,update"`
	NetworkProfile    NetworkProfile        `json:"networkProfile,omitempty"          visibility:"read,create"`
	NodePoolProfiles  []NodePoolProfile     `json:"nodePoolProfiles,omitempty"        visibility:"read"`
	EtcdEncryption    EtcdEncryptionProfile `json:"etcdEncryption,omitempty"          visibility:"read,create"`
	AutoRepair        bool                  `json:"autoRepair,omitempty"              visibility:"read,create,update"`
	Labels            []string              `json:"labels,omitempty"                  visibility:"read,create,update"`
	OIDCIssuerURL     json.URL              `json:"oidcIssuerUrl,omitempty"           visibility:"read"`
}

// ClusterProfile represents a high level cluster configuration.
// Visibility for the entire struct is "read,create,update".
type ClusterProfile struct {
	ControlPlaneVersion string `json:"controlPlaneVersion,omitempty"`
	SubnetID            string `json:"subnetId,omitempty"`
}

type Visibility string

const (
	VisibilityPublic  Visibility = "public"
	VisibilityPrivate Visibility = "private"
)

// APIProfile represents a cluster API server configuration.
// Visibility for the entire struct is "read,create".
type APIProfile struct {
	URL        json.URL   `json:"url,omitempty"`
	IP         net.IP     `json:"ip,omitempty"`
	Visibility Visibility `json:"visibility,omitempty"`
}

// ConsoleProfile represents a cluster web console configuration.
// Visibility for the entire struct is "read,create".
type ConsoleProfile struct {
	URL json.URL `json:"url,omitempty"`
}

// IngressProfile represents a cluster ingress configuration.
// Visibility for the entire struct is "read,create,update".
type IngressProfile struct {
	IP         net.IP     `json:"ip,omitempty"`
	Name       string     `json:"name,omitempty"`
	Visibility Visibility `json:"visibility,omitempty"`
}

// OutboundType represents a routing strategy to provide egress to the Internet.
type OutboundType string

const (
	OutboundTypeLoadBalancer OutboundType = "loadBalancer"
)

// PreconfiguredNSGs represents whether to use a bring-your-own network security
// group (NSG) attached to the subnets.
// FIXME Maybe convert this to a boolean type with JSON marshal/unmarshal
//       methods, even if the TypeSpec can't represent it as a boolean?
// Visibility for the entire struct is "read,create".
type PreconfiguredNSGs struct {
	Enabled  string `json:"enabled,omitempty"`
	Disabled string `json:"disabled,omitempty"`
}

// NetworkProfile represents a cluster network configuration.
// Visibility for the entire struct is "read,create".
type NetworkProfile struct {
	PodCIDR           json.IPNet        `json:"podCidr,omitempty"`
	ServiceCIDR       json.IPNet        `json:"serviceCidr,omitempty"`
	MachineCIDR       json.IPNet        `json:"machineCidr,omitempty"`
	HostPrefix        int32             `json:"hostPrefix,omitempty"`
	OutboundType      OutboundType      `json:"outBoundType,omitempty"`
	PreconfiguredNSGs PreconfiguredNSGs `json:"preconfiguredNsgs,omitempty"`
}

// NodePoolAutoscaling represents a node pool autoscaling configuration.
// Visibility for the entire struct is "read".
type NodePoolAutoscaling struct {
	MinReplicas int32 `json:minReplicas,omitempty"`
	MaxReplicas int32 `json:maxReplicas,omitempty"`
}

// NodePoolProfile represents a worker node pool configuration.
// Visibility for the entire struct is "read".
type NodePoolProfile struct {
	Name             string              `json:"name,omitempty"`
	Replicas         int32               `json:"replicas,omitempty"`
	SubnetID         string              `json:"subnetId,omitempty"`
	EncryptionAtHost bool                `json:"encryptionAtHost,omitempty"`
	VMSize           string              `json:"vmSize,omitempty"`
	Autoscaling      NodePoolAutoscaling `json:autoscaling,omitempty"`
}

// EtcdEncryptionProfile represents the configuration needed for customer
// provided keys to encrypt etcd storage.
// Visibility for the entire struct is "read,create".
type EtcdEncryptionProfile struct {
	DiscEncryptionSetID string `json:"discEncryptionSetId,omitempty"`
}
