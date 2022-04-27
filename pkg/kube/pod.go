package kube

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubelist/pkg/conndb"
	"strings"
)

type PodList struct {
	Name        string   `json:"name"`
	Namespace   string   `json:"namespace"`
	Image       string   `json:"image"`
	CpuLimit    *resource.Quantity   `json:"cpu_limet"`
	CpuRequest  *resource.Quantity   `json:"cpu_request"`
	MemoryLimit *resource.Quantity   `json:"memory_limit"`
	MemoryRequest *resource.Quantity `json:"memory_request"`
	PodIp       string   `json:"pod_ip"`
	NodeIp      string   `json:"node_ip"`
	Status      string   `json:"status"`
	ContainerId string   `json:"container_id"`

}

func CreatePodtable()(err error){
	err = conndb.DB.AutoMigrate(&PodList{})
	if err != nil{
		return err
	}
	return
}

func ListPod(c *kubernetes.Clientset,ns string){
	pods,err := c.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil{
		panic(err.Error())
	}
	for _,v := range pods.Items{
		var image,containerid string
		var cpulimit,memorylimit,cpurequest,memoryreques *resource.Quantity
		for _,i := range v.Spec.Containers{
			image = i.Image
			cpulimit = i.Resources.Limits.Cpu()
			memorylimit = i.Resources.Limits.Memory()
			cpurequest = i.Resources.Requests.Cpu()
			memoryreques = i.Resources.Requests.Memory()
		}
		for _,c := range v.Status.ContainerStatuses{
			dockerid := c.ContainerID
			list := strings.Split(dockerid,"//")
			containerid = list[1]
		}

		podlist := &PodList{
			Name: v.Name,
			Namespace: v.Namespace,
			Image: image,
			CpuLimit: cpulimit,
			MemoryLimit: memorylimit,
			CpuRequest: cpurequest,
			MemoryRequest: memoryreques,
			NodeIp: v.Status.HostIP,
			PodIp: v.Status.PodIP,
			Status: string(v.Status.Phase),
			ContainerId: containerid,
		}

		//var p PodList
		//conndb.DB.Where("name = ?",p.Name).Find(&p)
		//time.Sleep(time.Second*1)
		//fmt.Printf("PodName is :%#v\n",p.Name)
		//if p.Name == "" {
		//	conndb.DB.Create(&podlist)
		//}else {
		//	fmt.Println("----")
		//}


		//fmt.Printf("namespace: %v\n podname: %v\n podstatus: %v\n nodeip: %v\n podip: %v\n startTime: %v\n", v.Namespace, v.Name, v.Status.Phase,v.Status.HostIP,v.Status.PodIP,v.Status.StartTime.Time)
		fmt.Printf("PodName: %v\n namespace: %v\n image: %v\n cpulimit: %v\n memlimit: %v\n cpurequest: %v\n memrequest: %v\n nodeip: %v\n podip %v\n status: %v\n conid: %v\n ",
			podlist.Name,
			podlist.Namespace,
			podlist.Image,
			podlist.CpuLimit,
			podlist.MemoryLimit,
			podlist.CpuRequest,
			podlist.MemoryRequest,
			podlist.NodeIp,
			podlist.PodIp,
			podlist.Status,
			podlist.ContainerId)
	}
}
