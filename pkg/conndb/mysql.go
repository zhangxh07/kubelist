package conndb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"kubelist/setting"
)

var (
	DB *gorm.DB
	//err error
)
func Connmysql(cfg *setting.MysqlConfig)  (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	//dsn := "root:hao!2345@tcp(192.168.47.50:3306)/kube?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Println(err)
	}
	return nil
}

func Closedb()  {
	sqldb,_ := DB.DB()
	defer sqldb.Close()
}

