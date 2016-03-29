package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "wowauc.info"
	c.Data["Email"] = "webmaster@wowauc.info"
	c.TplName = "index.tpl"
}
