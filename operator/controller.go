package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	aro_clientset "github.com/Azure/ARO-HCP/operator/pkg/generated/clientset/versioned"
	aro_scheme "github.com/Azure/ARO-HCP/operator/pkg/generated/clientset/versioned/scheme"
	aro_informers "github.com/Azure/ARO-HCP/operator/pkg/generated/informers/externalversions/arohcp.openshift.io/v1alpha1"
	aro_listers "github.com/Azure/ARO-HCP/operator/pkg/generated/listers/arohcp.openshift.io/v1alpha1"
)

const controllerAgentName = "operation-controller"

// Controler is the controller implementation for Operation resources
type Controller struct {
	// kubeClientset is a standard Kubernetes clientset
	kubeClientset kubernetes.Interface
	// selfClientset is a clientset for our own API group
	selfClientset aro_clientset.Interface

	operationsLister aro_listers.OperationLister
	operationsSynced cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// NewController returns a new operations controller
func NewController(
	ctx context.Context,
	kubeClientset kubernetes.Interface,
	selfClientset aro_clientset.Interface,
	operationInformer aro_informers.OperationInformer) *Controller {

	logger := klog.FromContext(ctx)

	// Create event broadcaster
	utilruntime.Must(aro_scheme.AddToScheme(scheme.Scheme))
	logger.V(4).Info("Creating event broadcaster")

	eventBroadcaster := record.NewBroadcaster(record.WithContext(ctx))
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})
	ratelimiter := workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(5*time.Millisecond, 1000*time.Second),
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(50), 300)},
	)

	controller := &Controller{
		kubeClientset:    kubeClientset,
		selfClientset:    selfClientset,
		operationsLister: operationInformer.Lister(),
		operationsSynced: operationInformer.Informer().HasSynced,
		workqueue:        workqueue.NewRateLimitingQueue(ratelimiter),
		recorder:         recorder,
	}

	logger.Info("Setting up event handlers")
	// Set up an event handler for when Operation resources change
	operationInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueOperation,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueOperation(new)
		},
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(ctx context.Context, workers int) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()
	logger := klog.FromContext(ctx)

	// Start the informer factories to begin populating the informer cache
	logger.Info("Starting Operation controller")

	// Wait for the caches to be synced before starting workers
	logger.Info("Waiting for informer caches to sync")

	if ok := cache.WaitForCacheSync(ctx.Done(), c.operationsSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	logger.Info("Starting workers", "count", workers)
	// Launch workers to process Operation resources
	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	logger.Info("Started workers")
	<-ctx.Done()
	logger.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message
// on the workqueue.
func (c *Controller) runWorker(ctx context.Context) {
	for c.processNextWorkItem(ctx) {
	}
}

// processNextWorkItem will read a single work item off the workqueue
// and attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem(ctx context.Context) bool {
	obj, shutdown := c.workqueue.Get()
	logger := klog.FromContext(ctx)

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func() error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget
		// if we do not want this work item being re-queued.  For
		// example, we do not call Forget if a transient error occurs.
		// Instead the item is put back on the workqueue and attempted
		// again after a back-off period.
		defer c.workqueue.Done(obj)
		// Run the syncHandler, passing it the namespace/name string
		// of the Operation resource to be synced.
		if err := c.syncHandler(ctx, obj.(string)); err != nil {
			// Put the item back on the workqueue to handle any
			// transient errors.
			c.workqueue.AddRateLimited(obj)
			return fmt.Errorf("error syncing '%s': %s, requeuing", obj, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does
		// not get queued again until another change happens.
		c.workqueue.Forget(obj)
		logger.Info("Successfully synced", "resourceName", obj)
		return nil
	}()

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

// syncHandler fetches the current resource status from Cluster Service
// and updates the Status block of the Operation resource.
func (c *Controller) syncHandler(ctx context.Context, key string) error {
	logger := klog.LoggerWithValues(klog.FromContext(ctx), "resourceName", key)

	// Convert the namespace/name string into a distinct namespace and name.
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the Operation resource with this namespace/name.
	operation, err := c.operationsLister.Operations(namespace).Get(name)
	if err != nil {
		// The Operation resource may no longer exist, in which case we
		// stop processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("operation '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	// FIXME Finish business logic here.

	return nil
}

// enqueueOperation takes an Operation resource and converts it into a
// namespace/name string which is then put onto the work queue. This method
// should *not* be passed resources of any type other than Operation.
func (c *Controller) enqueueOperation(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}
