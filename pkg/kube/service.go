package kube

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateService(c *kubernetes.Clientset,ns string) error{
	svcClient := c.CoreV1().Services(ns)

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "client-nginx-service",
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Name: "http",
					Port: 80,
					NodePort: 30080,
				},
			},
			Selector: map[string]string{
				"app":"nginx",
			},
			Type: apiv1.ServiceTypeNodePort,
		},
	}
	result ,err := svcClient.Create(service)
	if err!=nil {
		return err
	}

	fmt.Printf("Create service %s \n", result.GetName())
	return nil
}
