package controllers


import (
	"github.com/astaxie/beego"
	"fmt"
	"strings"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := strings.TrimSpace(c.GetString("username"))
	password := strings.TrimSpace(c.GetString("password"))

	fmt.Println("获取的用户信息--->", username, password)
	if (username == "" && password != "") {
		c.Data["Error"] = "用户名密码不能为空！"	
	}

	c.TplName = "login.html"
}

