package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
)

func main() {
	//config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	//get data
	pod := v1.Pod{}
	err = restClient.Get().Namespace("default").Resource("pods").Name("collector-sdjxn").Do(context.TODO()).Into(&pod)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("pod.Name: %v\n", pod.Name)
	}
}
