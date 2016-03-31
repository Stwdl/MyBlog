package routers

import (
	"MyBlog/controllers"
	"log"

	"github.com/astaxie/beego"
)

func init() {
	log.Printf("初始化路由器")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/default", &controllers.MainController{}, "get:Default") //控制器/动作
}
