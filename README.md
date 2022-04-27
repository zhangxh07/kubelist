
动态获取K8s资源
动态获取K8s资源

go build   
./kubelist conf/config.ini

访问http://localhost:9090/api/v1/kubelist 返回如下json格式数据：

```json
{
    "status":200,
    "msg":"Sucessful",
    "services":[
        {
            "ID":2,
            "deployname":"",
            "namespace":"",
            "replicas":1,
            "ready":1,
            "image":"",
            "createtime":"",
            "updatetime":""
        }
    ],
    "total":2
}
```
git config --global http.sslVerify "false"