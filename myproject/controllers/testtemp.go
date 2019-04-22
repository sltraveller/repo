package controllers

import (
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"fmt"
)

type A struct{
	Name string
	Age int
}


//Temp is
func (c *ModelController) Temp() {
	c.Data["a"] = &A{Name:"hh", Age:22}
	
	mp := make(map[string]string)
	mp["name"]="hh"
	mp["buc"]="ru"
	c.Data["b"]=mp

	ss := []string{"a","b","c"}
	c.Data["c"]=ss
	

	o := orm.NewOrm()
	var up []models.Up
	num, err := o.Raw("select * from Up").QueryRows(&up)
	if err != nil {
		fmt.Println("q")
	}
	c.Data["num"]=num
	c.Data["user"]=up

	c.TplName = "template.html"
}