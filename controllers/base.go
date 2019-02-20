package controllers

import (
	"changeGo/common"
	"changeGo/models"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) Prepare() {
	auth := b.Ctx.Input.Header("Authorization")
	if auth == "" {
		http.Error(b.Ctx.ResponseWriter, "没有认证信息", 405)
	} else {
		auth_split := strings.Split(auth, " ")
		if len(auth_split) != 2 || auth_split[0] != "jwt" {
			http.Error(b.Ctx.ResponseWriter, "非jwt认证", 405)
		} else {
			//ss,_ := common.GetJwtAuth("lqm")
			aa, err := common.GetJwtName(auth_split[1])
			if err != nil {
				fmt.Println(err)
				http.Error(b.Ctx.ResponseWriter, fmt.Sprintln(err), 405)
			} else {
				fmt.Println("afg")
				fmt.Println(aa)
				models.NewUser().AddName(aa)
			}
		}
	}

}
