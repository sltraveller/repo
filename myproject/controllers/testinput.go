package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

import()

type TestInputController struct {
	beego.Controller
}

func (c *TestInputController) TestInputGet() {
	id := c.GetString("id")
	c.Ctx.WriteString(id)

	c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/testinput" method="post">
						用户名：<input type="text" name="Username">
						<br/>
						密&nbsp&nbsp&nbsp码：<input type="password" name="pwd">
						<br/>
						<input type="submit" value="提交">
						</form>
						</html>`)
}

type User struct {
	Username string
	Password string `form:"pwd"`
}

func (c *TestInputController) TestInputPost() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		logs.Info("查询错误")
	}
	c.Ctx.WriteString(u.Username+u.Password)
}