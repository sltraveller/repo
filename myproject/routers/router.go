package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{},"get:TestGet;post:TestPost")
	beego.Router("/testinput", &controllers.TestInputController{},"get:TestInputGet;post:TestInputPost")
	beego.Router("/testlogin", &controllers.TestLoginController{},"get:Login;post:Post")
	beego.Router("/testmodel", &controllers.ModelController{})
	beego.Router("/testquery", &controllers.ModelController{},"get:Query")
	beego.Router("/testsql", &controllers.ModelController{},"get:TSql")
	beego.Router("/testtemp", &controllers.ModelController{},"get:Temp")
}
