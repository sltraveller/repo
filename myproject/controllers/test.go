package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) TestGet() {
	c.Data["data"] = "id"
	c.TplName = "text.html"
}

func (c *TestController) TestPost() {
	c.Data["data"] = "testpost"
	c.TplName = "text.html"
}
