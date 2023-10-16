package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

var configFile = "C:/root/.kube/config"
var nameSpace = "default"

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
	jobList, _ := clientSet.BatchV1().Jobs(nameSpace).List(context.TODO(), metav1.ListOptions{Limit: 100})
	var deleteJobName string
	for _, v := range jobList.Items {
		fmt.Printf("Job Name: %v\n", v.Name)
		deleteJobName = v.Name
	}
	// 查询pod
	podList, _ := clientSet.CoreV1().Pods(nameSpace).List(context.TODO(), metav1.ListOptions{})
	for _, v := range podList.Items {
		fmt.Printf("Pod Name:%v\n", v.Name)
	}
	// 删除一个Job，设置级联删除
	propagationPolicy := metav1.DeletePropagationBackground
	err = clientSet.BatchV1().Jobs(nameSpace).Delete(context.TODO(), deleteJobName, metav1.DeleteOptions{PropagationPolicy: &propagationPolicy})
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 30)
	// 查询剩余pod，观察是否会级联删除pod
	podList, _ = clientSet.CoreV1().Pods(nameSpace).List(context.TODO(), metav1.ListOptions{})
	for _, v := range podList.Items {
		fmt.Printf("Pod Name:%v\n", v.Name)
	}
}
