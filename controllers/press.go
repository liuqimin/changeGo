package controllers

import (
	"changeGo/api"
	"changeGo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type PressController struct {
	beego.Controller
}

// @Title GetList
// @Description list for press
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [GET]
func (o *PressController) Get() {
	fmt.Println(12345)
	data, _ := models.NewPress().GetAll()
	o.Data["json"] = data
	o.ServeJSON()
}

// @Title AddPress
// @Description list for press
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [POST]
func (o *PressController) Post() {
	var pushData api.PressAddPost
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &pushData)
	if err != nil {
		o.Ctx.WriteString(fmt.Sprint(err))
	}
	models.NewPress().Insert(pushData.Name)

}
