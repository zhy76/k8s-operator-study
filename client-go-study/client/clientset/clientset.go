package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	coreV1 := clientSet.CoreV1()
	pod, err := coreV1.Pods("default").Get(context.TODO(), "collector-sdjxn", v1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		println(pod.Name)
	}
}
