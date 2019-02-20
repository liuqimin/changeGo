// @APIVersion 1.0.0
// @Title dbtools Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"changeGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//jwt认证
	beego.Router("/api-auth/", &controllers.UserController{}, "post:UserJWTLogin")
	//beego.Router("/apllo/version", &controllers.AplloVersioninfoController{}, "post:UserJWTLogin")
	beego.Router("/user-add/test/test/", &controllers.UserController{}, "post:AddUser")
	//rocketmq的数据收取和推送
	beego.Router("/api/thridinterface/apollo/", &controllers.AplloVersioninfoController{}, "get:GetData")
	beego.Router("/api/thridinterface/apolloPush/", &controllers.AplloVersioninfoController{}, "get:Test")
	//写普罗米修斯的落地文件，供监控使用
	beego.Router("/cron/monitor/itsm/", &controllers.ItsminfoController{}, "get:LoadFile")
	//定时任务同步
	beego.Router("/cron/monitor/itest/", &controllers.ItsminfoController{}, "get:ItsmCron")
	//定时任务注入etcd，通过变更etcd的值来达到所有定期任务进行变更
	beego.Router("/utils/etcd/push", &controllers.EtcdController{}, "get:Push")
	//beego.Router("/utils/etcd/watch", &controllers.EtcdController{}, "get:Push")
	beego.Router("/utils/etcd/get", &controllers.EtcdController{}, "get:GetData")
	//sql脚本检查，抽出字段，table名，字段名，sql类型 其他检验在python后台处理
	beego.Router("/sql/srcipt", &controllers.SqlParseController{}, "post:Srcipt")
	//sql语句检查，抽出字段，table名，字段名，sql类型
	beego.Router("/sql/sql", &controllers.SqlParseController{}, "post:GetData")
}
