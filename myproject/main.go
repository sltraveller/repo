package main

import (
	_ "myproject/routers"
	_ "myproject/models"
	"github.com/astaxie/beego"
)


func main() {
	beego.Run()
}
