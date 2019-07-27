package routers

import (
	"myweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{});
	beego.Router("/index", &controllers.LoginController{});
	
}
