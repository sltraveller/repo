package controllers

import (
	"github.com/astaxie/beego/orm"
)

//Query is
func (c *ModelController) Query() {
	o := orm.NewOrm()
	q := o.QueryTable("user")
	var users []*User
	num, err:=q.All(&users)
	
}