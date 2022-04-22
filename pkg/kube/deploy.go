package kube

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/utils/pointer"
	"k8s.io/client-go/kubernetes"
	"kubelist/pkg/conndb"
	"time"
)

//func init()  {
//	err := CreateTable()
//	if err != nil{
//		log.Println(err)
//	}
//}

type DeployInfo struct {
	ID           int
	Deployname   string `json:"deployname"`
	Namespace    string `json:"namespace"`
	Replicas     int32  `json:"replicas"`
	Ready        int32  `json:"ready"`
	Image        string `json:"image"`
	Createtime   string `json:"createtime"`
	Updatetime   string `json:"updatetime" gorm:"column:updatetime"`
}

func CreateTable()(err error){
	err = conndb.DB.AutoMigrate(&DeployInfo{})
	if err != nil{
		return err
	}
	return
}
func ListDeployment(c *kubernetes.Clientset,ns string) (err error){
	deployments,err := c.AppsV1().Deployments(ns).List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _,v := range deployments.Items{
		var image string
		t := v.Status.Conditions
		updatetime := t[len(t)-1].LastUpdateTime

		for _,i := range v.Spec.Template.Spec.Containers{
			image = i.Image
		}
		d := &DeployInfo{
			Deployname: v.Name,
			Namespace: v.Namespace,
			Replicas: v.Status.Replicas,
			Ready: v.Status.ReadyReplicas,
			Image: image,
			Createtime: v.CreationTimestamp.Time.Format("2006-01-02 15:04:05"),
			Updatetime: updatetime.Time.Format("2006-01-02 15:04:05"),
		}

		var s DeployInfo
		conndb.DB.Where("deployname = ?",d.Deployname).Find(&s)
		time.Sleep(time.Second*1)
		fmt.Printf("deployment is :%#v\n",s.Deployname)
		if s.Deployname == "" {
			//先写入数据
			conndb.DB.Create(&d)
		}else {
			fmt.Println(s.Deployname,"is exist in k8s")
			//再判断是否有信息变更
			if s.Replicas != d.Replicas{
				conndb.DB.Model(&s).Where("deployname = ?",d.Deployname).Update("replicas",d.Ready)
				fmt.Printf("replicas update %v\n",d.Replicas)
			}
			if s.Ready != d.Ready{
				conndb.DB.Model(&s).Where("deployname = ?",d.Deployname).Update("ready",d.Ready)
				fmt.Printf("ready update %v\n",d.Ready)
			}
			if s.Image != d.Image{
				conndb.DB.Model(&s).Where("deployname = ?",d.Deployname).Update("image",d.Image)
				fmt.Printf("image update %v\n",d.Image)
			}
			if s.Createtime != d.Createtime {
				conndb.DB.Model(&s).Where("deployname = ?",d.Deployname).Update("createtime",d.Createtime)
				fmt.Printf("createtime update %v\n",d.Createtime)
			}
			if s.Updatetime != d.Updatetime {
				conndb.DB.Model(&s).Where("deployname = ?",d.Deployname).Update("updatetime",d.Updatetime)
				fmt.Printf("updatetime update %v\n",d.Updatetime)
			}
		}
	}
	return
}

func FindDeploy() (service []*DeployInfo){
	var s []*DeployInfo
	conndb.DB.Find(&s)

	for k,v := range s{
		ser := &DeployInfo{
			ID: k,
			Deployname: v.Deployname,
			Namespace: v.Namespace,
			Replicas: v.Replicas,
			Ready: v.Ready,
			Image: v.Image,
			Createtime: v.Createtime,
			Updatetime: v.Updatetime,
		}
		service = append(service,ser)
	}
	//services = append(services,ser)
	return
}


func CreateDeployment(c *kubernetes.Clientset,ns string) error{
	deployClient  := c.AppsV1().Deployments(ns)

	deployment := &appsv1.Deployment{
		ObjectMeta : metav1.ObjectMeta{
			Name: "client-nginx",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta:metav1.ObjectMeta{
					Labels: map[string]string{
						"app":"nginx",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "nginx",
							Image: "nginx:1.16.1",
							ImagePullPolicy: "IfNotPresent",
							Ports: []apiv1.ContainerPort{
								{
									Name: "http",
									Protocol: apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	result ,err := deployClient.Create(deployment)
	if err!= nil{
		panic(err.Error())
	}
	fmt.Println("Create deployment %s \n", result.GetName())
	return err
}