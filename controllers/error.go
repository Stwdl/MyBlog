package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

///404未找到页面
func (c *ErrorController) Error404() {
	c.Data["appname"] = "醉翁二郎"
	c.Data["appurl"] = "localhost:8080"
	c.Data["errimg"] = "../static/img/404.jpg"
	c.TplName = "error.tpl"
}
