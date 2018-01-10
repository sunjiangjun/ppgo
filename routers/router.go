package routers

import (
	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/controllers"
)

func init() {
	// 默认登录
	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	//自定义--sunhongtao
	beego.Router("log/logs",&controllers.LogController{},"*:Logs")

	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.ApiDocController{})
	beego.AutoRouter(&controllers.ApiMonitorController{})
	beego.AutoRouter(&controllers.EnvController{})
	beego.AutoRouter(&controllers.CodeController{})

	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

}
