package main

import (
	"log"

	"github.com/zhy76/k8s-operator-study/client-go-study/client-go-demo/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1.config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		inClusterConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("can't get config")
		}
		config = inClusterConfig
	}

	// 2.client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("can't create client")
	}

	// 3.informer
	factory := informers.NewSharedInformerFactory(clientSet, 0)
	serviceInformer := factory.Core().V1().Services()
	ingressInformer := factory.Networking().V1().Ingresses()

	// 4.add event handler

	// 5.informer.start
	controller := pkg.NewController(clientSet, serviceInformer, ingressInformer)
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)

	controller.Run(stopCh)
}
