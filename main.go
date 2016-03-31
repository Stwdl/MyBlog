package main

import (
	"log"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	beego.BeeLogger.Info("设置静态路径")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "Home.html"
}

func (c *MainController) Default() {
	beego.BeeLogger.Info("设置静态路径")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func main() {

	log.Printf("程序开始")
	beego.Router("/", &MainController{})
	beego.Router("/default", &MainController{}, "get:Default") //控制器/动作

	beego.Run("0.0.0.0:5050")
}
