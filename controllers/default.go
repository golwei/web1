package controllers

import (
	"github.com/astaxie/beego"
)

//Main{}
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// User{}
type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {
	c.TplName = "user/profile.html"
}
