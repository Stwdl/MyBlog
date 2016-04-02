package routers

import (
	"log"
	"myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	log.Printf("初始化路由器")
	beego.Router("/", &controllers.MainController{})

}
