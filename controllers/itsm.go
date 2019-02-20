package controllers

import (
	"changeGo/api"
	"changeGo/common/monitor"
	"changeGo/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ItsminfoController struct {
	beego.Controller
}

func (a *ItsminfoController) GetData() {
	monitorcommon.Test()
}

func (a *ItsminfoController) PushData() {
	var getData models.ApolloGet
	//models.NewApolloInfo()
	fmt.Println(getData)
	//common.Publish()
}
func (a *ItsminfoController) LoadFile() {
	monitorcommon.GetItsmStatus()
}

func (a *ItsminfoController) ItsmCron() {
	api.ChangITSMSyncThreeDay()
}
