package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
)

// TestLoginController is
type TestLoginController struct {
	beego.Controller
}

//UserInfo is
type UserInfo struct {
	Username string
	Password string `form:"pwd"`
}

//Login is
func (c *TestLoginController) Login() {

	username := c.Ctx.GetCookie("username")
	password := c.Ctx.GetCookie("password")
	fmt.Println("login:username:",username,",password",password)
	if username != "" {
		c.Ctx.WriteString("Username:" + username + ",Password:" + password)
	}else{
	c.Ctx.WriteString(`<html><form action="/testlogin" method="post">
						用户名：<input type="text" name="Username" />
						<br/>
						密&nbsp&nbsp&nbsp码：<input type="password" name="pwd">
						<br/>
						<input type="submit" value="提交">
						</form>
						</html>`)
	}
/*
	username := c.GetSession("username")
    password := c.GetSession("password")
    if nameString, ok := username.(string); ok && nameString != "" {
        c.Ctx.WriteString("Username:" + username.(string) + ",Password:" + password.(string))
	}else{
	c.Ctx.WriteString(`<html><form action="/testinput" method="post">
						用户名：<input type="text" name="Username" />
						<br/>
						密&nbsp&nbsp&nbsp码：<input type="password" name="pwd">
						<br/>
						<input type="submit" value="提交">
						</form>
						</html>`)
	}
*/
}

//Post is
func (c *TestLoginController) Post() {
	u := UserInfo{}
	if err := c.ParseForm(&u); err != nil {
		logs.Info("")
	}
	fmt.Println("post:username:",u.Username,"password",u.Password)
	c.Ctx.SetCookie("username",u.Username,20,"/")
	c.Ctx.SetCookie("password",u.Password,20,"/")

	//c.SetSession("username",u.Username)
	//c.SetSession("password",u.Password)
	c.Ctx.WriteString("Username:" + u.Username + ",Password:" + u.Password)
}