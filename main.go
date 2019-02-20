package main

import (
	"changeGo/cron"
	"changeGo/models"
	_ "changeGo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
	go cron.CronInit()
	go cron.CronWatch()
	//orm.RegisterDriver("mysql", orm.DRMySQL)    //可以不加
	//orm.RegisterDataBase("default", "mysql", "dbtools:Test123@123@tcp(172.16.8.100:3306)/bloggo?charset=utf8")
	//orm.RunSyncdb("default", false, true)
	//ORM 必须注册一个别名为default的数据库，作为默认使用。
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
