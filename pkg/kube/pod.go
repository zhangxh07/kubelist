package kube

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPod(c *kubernetes.Clientset,ns string){
	pods,err := c.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil{
		panic(err.Error())
	}
	for _,v := range pods.Items{
		fmt.Printf("namespace: %v\n podname: %v\n podstatus: %v\n nodeip: %v\n podip: %v\n startTime: %v\n", v.Namespace, v.Name, v.Status.Phase,v.Status.HostIP,v.Status.PodIP,v.Status.StartTime.Time)
	}
}

