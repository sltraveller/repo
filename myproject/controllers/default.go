package controllers

import (
	"github.com/astaxie/beego"
)

//MainController is
type MainController struct {
	beego.Controller
}

//Get is
func (c *MainController) Get() {

	c.Data["data"] = "ok"
	c.TplName = "text.html"
}

func (c *MainController) Post() {
	c.Data["data"] = "world"
	c.TplName = "text.html"
}
