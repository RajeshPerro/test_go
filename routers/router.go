package routers

import (
	"github.com/astaxie/beego"
	"test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{}, "get:Register")
	beego.Router("/register", &controllers.RegisterController{}, "post:DataCatch")
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/login", &controllers.LoginController{}, "post:LoginCheck")

}
