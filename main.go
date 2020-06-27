package main

import (
	"anydevelop.cn/recruit-server/redis"
	_ "anydevelop.cn/recruit-server/redis"
	_ "anydevelop.cn/recruit-server/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:xy942698.@tcp(127.0.0.1:3306)/recruit?charset=utf8", 30)
	defer redis.Rdb.Close()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
