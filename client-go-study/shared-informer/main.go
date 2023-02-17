package main

import (
	"fmt"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	// create config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	// create client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// get informer
	// factory := informers.NewSharedInformerFactory(clientSet, 0)
	factory := informers.NewSharedInformerFactoryWithOptions(clientSet, 0, informers.WithNamespace("default"))
	informer := factory.Core().V1().Pods().Informer()

	// add workqueue
	rateLimitingQueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller")

	// add event handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add Event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				fmt.Println("can't get key")
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Update Event")
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err != nil {
				fmt.Println("can't get key")
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Delete Event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				fmt.Println("can't get key")
			}
			rateLimitingQueue.AddRateLimited(key)
		},
	})

	// start informer
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
