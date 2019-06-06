/*
   Sentinel - Copyright (c) 2019 by www.gatblau.org

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software distributed under
   the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
   either express or implied.
   See the License for the specific language governing permissions and limitations under the License.

   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"strings"
	"time"
)

// a k8s controller that watches for changes to the state of a particular resource
// and triggers the execution of a publisher (e.g. calling a web hook,
// putting a message in a broker, etc.)
type Watcher struct {
	objType   string
	queue     workqueue.RateLimitingInterface
	informer  cache.SharedIndexInformer
	publisher Publisher
}

// creates a new controller to watch for changes in status of a specific resource
func newWatcher(informer cache.SharedIndexInformer, objType string, publisher Publisher) *Watcher {
	logrus.Tracef("Creating %s watcher.", strings.ToUpper(objType))
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	var change Change
	var err error

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			change.key, err = cache.MetaNamespaceKeyFunc(obj)
			change.changeType = "CREATE"
			change.objectType = objType
			change.namespace = getMetaData(obj).Namespace
			addToQueue(queue, change, err)
		},
		UpdateFunc: func(obj, new interface{}) {
			change.key, err = cache.MetaNamespaceKeyFunc(obj)
			change.changeType = "UPDATE"
			change.objectType = objType
			change.namespace = getMetaData(obj).Namespace
			addToQueue(queue, change, err)
		},
		DeleteFunc: func(obj interface{}) {
			change.key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			change.changeType = "DELETE"
			change.objectType = objType
			change.namespace = getMetaData(obj).Namespace
			addToQueue(queue, change, err)
		},
	})

	return &Watcher{
		objType:   objType,
		informer:  informer,
		queue:     queue,
		publisher: publisher,
	}
}

// runs the controller
func (w *Watcher) run() {
	// creates a stopCh channel to stop the controller when required
	stopCh := make(chan struct{})
	defer close(stopCh)

	// catches a crash and logs an error
	// TODOs: check if it can be removed as apiserver will handle panics
	defer runtime.HandleCrash()

	// shut downs the queue when it is time
	defer w.queue.ShutDown()

	logrus.Tracef("%s watcher starting.", strings.ToUpper(w.objType))
	startTime = time.Now().Local()

	// starts and runs the shared informer
	// the informer will be stopped when the stop channel is closed
	go w.informer.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, w.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync."))
		return
	}

	logrus.Tracef("%s watcher synchronised and ready.", strings.ToUpper(w.objType))

	// loops until the stop channel is closed, running the worker every second
	wait.Until(w.processQueue, time.Second, stopCh)
}

// process the items in the controller's queue
func (w *Watcher) processQueue() {
	// loops until the worker queue is shut down
	for w.nextItem() {
	}
}

// process the next item in the controller's queue
func (w *Watcher) nextItem() bool {
	// waits until there is a new item in the working queue
	key, shutdown := w.queue.Get()

	// if queue shuts down then quit
	if shutdown {
		logrus.Tracef("%s queue has shut down.", strings.ToUpper(w.objType))
		return false
	}

	// tells the queue that we are done processing this key
	// this unblocks the key for other workers and allows safe parallel processing because two pods
	// with the same key are never processed in parallel.
	defer w.queue.Done(key)

	// passes the queue item to the registered handler(s)
	err := w.publish(key.(Change))

	// handles the result of the previous operation
	// if something went wrong during the execution of the business logic, triggers a retry mechanism
	w.handleResult(err, key)

	// continues processing
	return true
}

// publish the state change
func (w *Watcher) publish(change Change) error {
	logrus.Tracef("Ready to publish %s changes for %s %s.", change.changeType, strings.ToUpper(change.objectType), change.key)
	obj, exists, err := w.informer.GetIndexer().GetByKey(change.key)
	if !exists {
		logrus.Tracef("%s %s does not exist anymore.", strings.ToUpper(change.objectType), change.key)
	}
	if err != nil {
		return fmt.Errorf("Failed to retrieve object with key %s: %s", change.key, err)
	} else {
		// get object metadata
		meta := getMetaData(obj)

		// publish events based on its type
		switch change.changeType {
		case "CREATE":
			// compare CreationTimestamp and serverStartTime and alert only on latest events
			// Could be Replaced by using Delta or DeltaFIFO
			if meta.CreationTimestamp.Sub(startTime).Seconds() > 0 {
				logrus.Tracef("Calling Publisher.OnCreate(change -> %+v).", change)
				w.publisher.OnCreate(change, meta)
				return nil
			}
		case "UPDATE":
			logrus.Tracef("Calling Publisher.OnUpdate(change -> %+v).", change)
			w.publisher.OnUpdate(change, meta)
			return nil
		case "DELETE":
			logrus.Tracef("Calling Publisher.OnDelete(change -> %+v).", change)
			w.publisher.OnDelete(change, meta)
			return nil
		}
	}
	return nil
}

// checks if an error has happened triggering retry
// or stops retrying if there is no error
func (w *Watcher) handleResult(err error, key interface{}) {
	if err == nil {
		// indicates that the item is finished being retried.
		// it doesn't matter whether it's for permanent failing or for success,
		// it stops the rate limiter from tracking it.
		logrus.Tracef("Change for %s has been processed.", key.(Change).key)
		w.queue.Forget(key)
		return
	} else if w.queue.NumRequeues(key) < maxRetries {
		// this controller retries a specified number of times if something goes wrong
		// after which, stops trying
		logrus.Errorf("Error processing %s (will retry): %s.", key.(Change).key, err)

		// re-queue the key rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the key will be processed later again.
		w.queue.AddRateLimited(key)
		return
	} else {
		// err != nil and too many retries, then give up
		w.queue.Forget(key)

		// reports to an external entity that, even after several retries,
		// the key could not be successfully handled
		runtime.HandleError(err)

		// logs the error
		if err != nil {
			logrus.Errorf("Error processing %s (giving up): %s.", key.(Change).key, err)
		} else {
			logrus.Errorf("Error processing %s: too many retries, giving up!", key.(Change).key)
		}
	}
}
