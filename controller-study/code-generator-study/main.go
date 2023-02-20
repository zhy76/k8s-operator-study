package main

import (
	"context"
	clientset "github.com/zhy76/k8s-operator-study/controller-study/code-generator-study/pkg/generated/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := clientset.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	list, err := clientset.CrdV1().Foos("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, foo := range list.Items {
		println(foo.Name)
	}
}
