package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
)

const (
	// Literal path segments must be lowercase because
	// MiddlewareLowercase converts the request URL to
	// lowercase before multiplexing.
	ResourceProviderNamespace = "microsoft.redhatopenshift"

	SubscriptionPath  = "/subscriptions/{subscriptionId}"
	ResourceGroupPath = SubscriptionPath + "/resourcegroups/{resourceGroupName}"
	ResourceTypePath  = ResourceGroupPath + "/" + ResourceProviderNamespace + "/{resourceType}"
	ResourceNamePath  = ResourceTypePath + "/{resourceName}"
)

type PathValues struct {
	SubscriptionID    string
	ResourceGroupName string
	ResourceType      string
	ResourceName      string
}

func NewPathValues(request *http.Request) *PathValues {
	return &PathValues{
		SubscriptionID:    request.PathValue("subscriptionId"),
		ResourceGroupName: request.PathValue("resourceGroupName"),
		ResourceType:      request.PathValue("resourceType"),
		ResourceName:      request.PathValue("resourceName"),
	}
}

type Frontend struct {
	logger   *slog.Logger
	listener net.Listener
	server   http.Server
	done     chan struct{}
}

func NewFrontend(logger *slog.Logger, listener net.Listener) *Frontend {
	f := &Frontend{
		logger:   logger,
		listener: listener,
		server: http.Server{
			ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		},
		done: make(chan struct{}),
	}

	mux := NewMiddlewareMux(
		MiddlewarePanic,
		MiddlewareBody,
		MiddlewareLowercase,
		MiddlewareSystemData)
	// FIXME List post-multiplexing middleware like auth validation here.
	postMuxMiddleware := NewMiddleware()
	mux.Handle(
		"GET "+ResourceTypePath,
		postMuxMiddleware.HandlerFunc(f.ArmResourceListByParent))
	mux.Handle(
		"GET "+ResourceNamePath,
		postMuxMiddleware.HandlerFunc(f.ArmResourceRead))
	mux.Handle(
		"PUT "+ResourceNamePath,
		postMuxMiddleware.HandlerFunc(f.ArmResourceCreateOrUpdate))
	mux.Handle(
		"PATCH "+ResourceNamePath,
		postMuxMiddleware.HandlerFunc(f.ArmResourcePatch))
	mux.Handle(
		"DELETE "+ResourceNamePath,
		postMuxMiddleware.HandlerFunc(f.ArmResourceDelete))
	mux.Handle(
		"POST "+ResourceNamePath,
		postMuxMiddleware.HandlerFunc(f.ArmResourceAction))
	f.server.Handler = mux

	return f
}

func (f *Frontend) Run(ctx context.Context, stop <-chan struct{}) {
	if stop != nil {
		go func() {
			<-stop
			f.server.Shutdown(ctx)
		}()
	}

	f.logger.Info(fmt.Sprintf("listening on %s", f.listener.Addr().String()))

	err := f.server.Serve(f.listener)
	if err != http.ErrServerClosed {
		f.logger.Error(err.Error())
		os.Exit(1)
	}

	close(f.done)
}

func (f *Frontend) Join() {
	<-f.done
}

func (f *Frontend) ArmResourceListByParent(writer http.ResponseWriter, request *http.Request) {
	pathValues := NewPathValues(request)
	fmt.Println(*pathValues)
}

func (f *Frontend) ArmResourceRead(writer http.ResponseWriter, request *http.Request) {
	pathValues := NewPathValues(request)
	fmt.Println(*pathValues)
}

func (f *Frontend) ArmResourceCreateOrUpdate(writer http.ResponseWriter, request *http.Request) {
	pathValues := NewPathValues(request)
	fmt.Println(*pathValues)
}

func (f *Frontend) ArmResourcePatch(writer http.ResponseWriter, request *http.Request) {
	pathValues := NewPathValues(request)
	fmt.Println(*pathValues)
}

func (f *Frontend) ArmResourceDelete(writer http.ResponseWriter, request *http.Request) {
	pathValues := NewPathValues(request)
	fmt.Println(*pathValues)
}

func (f *Frontend) ArmResourceAction(writer http.ResponseWriter, request *http.Request) {
	pathValues := NewPathValues(request)
	fmt.Println(*pathValues)
}
