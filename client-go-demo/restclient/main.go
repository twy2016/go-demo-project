package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var configFile = "/root/.kube/config"
var nameSpace = "kube-system"
var resource = "deployments"

func main() {
	// 生成config
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err)
	}
	// 参考api请求路径 /apis/apps/v1/namespaces/{namespace}/deployments
	config.APIPath = "apis"
	// 资源的group
	config.GroupVersion = &appsv1.SchemeGroupVersion
	// 指定序列化工具
	config.NegotiatedSerializer = scheme.Codecs

	// 生成restClient
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// 声明空结构体
	deploymentList := &appsv1.DeploymentList{}
	// 发起get请求
	if err = restClient.Get().
		Namespace(nameSpace).
		Resource(resource).
		// 指定大小限制和序列化工具
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).
		// 发送请求
		Do(context.TODO()).
		// 结果存入deploymentList
		Into(deploymentList); err != nil {
		panic(err)
	}
	for _, v := range deploymentList.Items {
		fmt.Printf("NameSpace: %v  Name: %v\n", v.Namespace, v.Name)
	}
}
