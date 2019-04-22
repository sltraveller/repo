package controllers

import (
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"fmt"
)

// TSql is
func (c *ModelController) TSql() {
	o := orm.NewOrm()
	// var maps [] orm.Params
	// num, err := o.Raw("select * from up").Values(&maps)
	// var up models.Up
	// err := o.Raw("select * from Up where id=6").QueryRow(&up)
	var up []models.Up
	num, err := o.Raw("select * from Up").QueryRows(&up)
	if err != nil {
		fmt.Println("q")
	}
	c.Ctx.WriteString(fmt.Sprintf("共查询了 num:%d 条数据。。\n", num))
	for _, m:= range up {
		//c.Ctx.WriteString(fmt.Sprintf("%s\n", m["name"]),)
		c.Ctx.WriteString(fmt.Sprintf("%s\n", m.Name))
	}
	//c.Ctx.WriteString(fmt.Sprintf("id:%d, name:%s",up.Id,up.Name))
}