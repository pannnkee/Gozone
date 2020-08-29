package routers

import (
	"Gozone/src/zone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{}, "*:Home")
    beego.Router("/login", &controllers.HomeController{}, "*:Login")
    beego.Router("/register", &controllers.HomeController{}, "*:Register")

    beego.Router("/article", &controllers.HomeController{},"*:Article")


	v1 := beego.NewNamespace("/v1/api",
		beego.NSNamespace("/user",
			beego.NSRouter("/register", &controllers.UserController{}, "post,options:Register"),
			beego.NSRouter("/login", &controllers.UserController{}, "post,options:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "*:Logout"),
			),

		beego.NSNamespace("/article",
				beego.NSRouter("/page", &controllers.ArticleController{}, "*:PageList"),
			),
	)
	beego.AddNamespace(v1)
}
