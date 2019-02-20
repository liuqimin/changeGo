package controllers

import (
	"changeGo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title getAllUser
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [GET]
func (u *UserController) GetUser() {

	data, err := models.NewUser().GetAllUser()
	if err != nil {
		u.Data["json"] = fmt.Sprintln(err)
	} else {
		u.Data["json"] = data
	}

	u.ServeJSON()

}

func (u *UserController) AddUser() {
	var user_post models.UserPost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user_post)
	if err != nil {
		u.Ctx.WriteString(fmt.Sprintln(err))
	}
	if len(user_post.Username) == 0 || len(user_post.Password) == 0 {
		u.Ctx.WriteString(fmt.Sprintln("缺少参数username或password"))
	} else {
		err := models.NewUser().Add(user_post.Username, user_post.Password)
		if err != nil {
			u.Ctx.WriteString(fmt.Sprintln(err))
		} else {
			u.Ctx.WriteString(fmt.Sprintln("增加用户成功"))
		}

	}
}

func (u *UserController) UserJWTLogin() {

	var user_post models.UserPost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user_post)
	if err != nil {
		u.Ctx.WriteString(fmt.Sprintln(err))
	}
	if len(user_post.Username) == 0 || len(user_post.Password) == 0 {
		u.Ctx.WriteString(fmt.Sprintln("缺少参数username或password"))
	} else {
		models.NewUser().Verification(user_post.Username, user_post.Password)
		u.Ctx.WriteString(fmt.Sprintln("登陆成功"))
	}
}
