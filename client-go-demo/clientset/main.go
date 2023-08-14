package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var configFile = "/root/.kube/config"
var nameSpace = "kube-system"

func main() {
	// 生成config
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err)
	}

	// 生成clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deploymentList, err := clientSet.AppsV1().Deployments(nameSpace).List(context.TODO(), metav1.ListOptions{Limit: 100})
	for _, v := range deploymentList.Items {
		fmt.Printf("NameSpace: %v  Name: %v\n", v.Namespace, v.Name)
	}
}
