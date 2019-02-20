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
			fmt.Println(err)
			return
		}
		initcommon.Etcdpush(cli, key, value)
		defer cli.Close()
	} else {
		fmt.Println("缺少参数")
	}
}

func (e *EtcdController) GetData() {
	//var data EtcdPushData
	//idStr := e.Ctx.Input.Param(":name")

	key := "/test/123"
	fmt.Println(key)
	cli, err := initcommon.Etcdinit()
	if err != nil {
		fmt.Println(err)
		return
	}
	initcommon.EtcdGet(cli, key)
	defer cli.Close()
}
