package controllers

import (
	"changeGo/common/init"
	"fmt"
	"github.com/astaxie/beego"
)

type EtcdController struct {
	beego.Controller
}

type EtcdPushData struct {
	key   string
	value string
}

func (e *EtcdController) Push() {
	//var data EtcdPushData
	//idStr := e.Ctx.Input.Param(":name")
	key := e.GetString("key")
	value := e.GetString("value")
	if len(key) != 0 || len(value) != 0 {
		cli, err := initcommon.Etcdinit()
		if err != nil {
			e.Data["json"] = map[string]interface{}{"success": 1, "message": "缺少参数"}
			e.ServeJSON()
		}
		initcommon.Etcdpush(cli, key, value)
		e.Data["json"] = map[string]interface{}{"success": 0, "message": "变更成功"}
		e.ServeJSON()
		defer cli.Close()
	} else {
		fmt.Println("缺少参数")
		e.Data["json"] = map[string]interface{}{"success": 1, "message": "缺少参数"}
		e.ServeJSON()
	}
}

func (e *EtcdController) GetData() {
	//var data EtcdPushData
	//idStr := e.Ctx.Input.Param(":name")

	//key := "/test/123"
	//key 命名  appid/应用名 value cront格式  * * * * * * aaaa
	key := e.GetString("key")
	fmt.Println(key)
	cli, err := initcommon.Etcdinit()
	if err != nil {
		fmt.Println(err)
		return
	}
	initcommon.EtcdGet(cli, key)
	defer cli.Close()
}
