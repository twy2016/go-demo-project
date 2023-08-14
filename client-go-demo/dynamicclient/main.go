package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var configFile = "/root/.kube/config"
var nameSpace = "kube-system"
var gvr = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

func main() {
	// 生成config
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err)
	}
	// 生成dynamicClient
	dynamicClient, err := dynamic.NewForConfig(config)
	unstructObj, err := dynamicClient.
		Resource(gvr).
		Namespace(nameSpace).
		List(context.TODO(), metav1.ListOptions{Limit: 100})
	if err != nil {
		panic(err)
	}
	// 声明空结构体
	deploymentList := &appsv1.DeploymentList{}
	// 转换
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), deploymentList)
	if err != nil {
		panic(err.Error())
	}
	for _, v := range deploymentList.Items {
		fmt.Printf("NameSpace: %v  Name: %v\n", v.Namespace, v.Name)
	}
}
