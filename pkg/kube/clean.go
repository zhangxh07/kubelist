package kube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func clean(c *kubernetes.Clientset,ns,deploy,svc string){
	emptyDeleteOptions := &metav1.DeleteOptions{}
	if err := c.CoreV1().Services(ns).Delete(svc, emptyDeleteOptions);err != nil{
		panic(err.Error())
	}

	if err := c.AppsV1().Deployments(ns).Delete(deploy, emptyDeleteOptions); err != nil{
		panic(err.Error())
	}

}
