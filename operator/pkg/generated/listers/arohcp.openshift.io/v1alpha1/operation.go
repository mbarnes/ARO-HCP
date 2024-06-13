// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Azure/ARO-HCP/operator/pkg/apis/arohcp.openshift.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OperationLister helps list Operations.
// All objects returned here must be treated as read-only.
type OperationLister interface {
	// List lists all Operations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Operation, err error)
	// Operations returns an object that can list and get Operations.
	Operations(namespace string) OperationNamespaceLister
	OperationListerExpansion
}

// operationLister implements the OperationLister interface.
type operationLister struct {
	indexer cache.Indexer
}

// NewOperationLister returns a new OperationLister.
func NewOperationLister(indexer cache.Indexer) OperationLister {
	return &operationLister{indexer: indexer}
}

// List lists all Operations in the indexer.
func (s *operationLister) List(selector labels.Selector) (ret []*v1alpha1.Operation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Operation))
	})
	return ret, err
}

// Operations returns an object that can list and get Operations.
func (s *operationLister) Operations(namespace string) OperationNamespaceLister {
	return operationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OperationNamespaceLister helps list and get Operations.
// All objects returned here must be treated as read-only.
type OperationNamespaceLister interface {
	// List lists all Operations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Operation, err error)
	// Get retrieves the Operation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Operation, error)
	OperationNamespaceListerExpansion
}

// operationNamespaceLister implements the OperationNamespaceLister
// interface.
type operationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Operations in the indexer for a given namespace.
func (s operationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Operation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Operation))
	})
	return ret, err
}

// Get retrieves the Operation from the indexer for a given namespace and name.
func (s operationNamespaceLister) Get(name string) (*v1alpha1.Operation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("operation"), name)
	}
	return obj.(*v1alpha1.Operation), nil
}
