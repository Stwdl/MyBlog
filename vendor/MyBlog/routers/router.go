package routers

import (
	"MyBlog/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/default", &controllers.MainController{}, "get:Default") //控制器/动作
}
