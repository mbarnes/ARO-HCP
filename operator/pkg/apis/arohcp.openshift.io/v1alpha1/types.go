package v1alpha1

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Operation is a specification for an Operation resource
type Operation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OperationSpec   `json:"spec"`
	Status OperationStatus `json:"status"`
}

// OperationSpec is the spec for an Operation resource
type OperationSpec struct {
	Request         string `json:"request"`
	ExternalID      string `json:"externalId"`
	InternalID      string `json:"internalId"`
	NotificationURI string `json:"notificationUri"`
}

// OperationStatus is the status for an Operation resource
type OperationStatus struct {
	LastProbeTime    metav1.Time `json:"lastProbeTime"`
	LastModifiedTime metav1.Time `json:"lastModifiedTime"`
	State            string      `json:"state"`
	Details          string      `json:"details"`
	Terminal         bool        `json:"terminal"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OperationList is a list of Operation resources
type OperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Operation `json:"items"`
}
