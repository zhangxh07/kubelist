package main

import (
	//"context"
	"fmt"
	"kubelist/pkg/conndb"
	"kubelist/server"
	"kubelist/setting"
	"os"
	"sync"

	"kubelist/pkg/kube"
)
var wg sync.WaitGroup
func init()  {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./kubelist conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	err := conndb.Connmysql(setting.Conf.MysqlConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
}


func main(){
	err := kube.CreateTable()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	wg.Add(2)
	go server.Forlist()
	go server.Ginrun()
	wg.Wait()


	//ListPod(client,"default")

	//if err = CreateDeployment(client,"default");err != nil{
	//	fmt.Printf("create deployment failed errror: %v\n",err)
	//}else {
	//	fmt.Println("successful!")
	//}
	//if err = CreateService(client,"default");err != nil{
	//	fmt.Printf("create service failed errror: %v\n",err)
	//}else {
	//	fmt.Println("successful!")
	//}

	//clean(client,"default","client-nginx","client-nginx-service")

}









