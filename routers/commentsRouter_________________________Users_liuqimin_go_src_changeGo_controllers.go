package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["changeGo/controllers:AuthorController"] = append(beego.GlobalControllerRouter["changeGo/controllers:AuthorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:AuthorController"] = append(beego.GlobalControllerRouter["changeGo/controllers:AuthorController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:AuthorController"] = append(beego.GlobalControllerRouter["changeGo/controllers:AuthorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"DELETE"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:BookController"] = append(beego.GlobalControllerRouter["changeGo/controllers:BookController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:BookController"] = append(beego.GlobalControllerRouter["changeGo/controllers:BookController"],
		beego.ControllerComments{
			Method:           "BookUpdate",
			Router:           `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["changeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["changeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["changeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["changeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["changeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:PressController"] = append(beego.GlobalControllerRouter["changeGo/controllers:PressController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:PressController"] = append(beego.GlobalControllerRouter["changeGo/controllers:PressController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["changeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["changeGo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams:     param.Make(),
			Params:           nil})

}
