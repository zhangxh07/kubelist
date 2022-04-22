package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubelist/pkg/kube"
	"kubelist/setting"
	"sync"
	"time"
)
var wg sync.WaitGroup
type Result struct {
	Status int         		`json:"status"`
	Msg    string      		`json:"msg"`
	Services   interface{}  `json:"services"`
	Total  int         		`json:"total"`
}

func Forlist()  {
	config, err := clientcmd.BuildConfigFromFlags("","certs/config")
	if err != nil {
		fmt.Println("k8s config failed",err)
	}
	client, _ := kubernetes.NewForConfig(config)
	defer wg.Done()
	for {
		if err = kube.ListDeployment(client,"default");err != nil{
			fmt.Printf("list deployment error: %v", err)
		}
		time.Sleep(time.Second*1)
	}
}

func Ginrun(){
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	defer wg.Done()
	r := gin.Default()
	r.GET("/api/v1/kubelist", func(c *gin.Context) {
		c.JSON(200,Result{
			Status: 200,
			Msg: "Sucessful",
			Services: kube.FindDeploy(),
			Total: len(kube.FindDeploy()),
		})
	})
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
