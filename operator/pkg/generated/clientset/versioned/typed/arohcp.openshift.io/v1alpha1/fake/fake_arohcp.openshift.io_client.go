// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/Azure/ARO-HCP/operator/pkg/generated/clientset/versioned/typed/arohcp.openshift.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeArohcpV1alpha1 struct {
	*testing.Fake
}

func (c *FakeArohcpV1alpha1) Operations(namespace string) v1alpha1.OperationInterface {
	return &FakeOperations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeArohcpV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
