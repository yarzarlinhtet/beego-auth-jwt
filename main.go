package main

import (
	_ "auth-api/routers"
	"github.com/beego/beego/v2/client/orm"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	url, _ := beego.AppConfig.String("sqlconn")
	orm.RegisterDataBase("default", "mysql", url)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
