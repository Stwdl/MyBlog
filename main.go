package main

import (
	"myblog/controllers"
	//	"myblog/models"
	_ "myblog/routers"

	"github.com/astaxie/beego"
)

///返回存字符串的限定字数的文字
func toHtml(in string, count int) (out string) {
	content := beego.Substr(beego.Htmlunquote(beego.HTML2str(in)), 0, count)
	if len(in) > count {
		return content + "......"
	} else {
		return content
	}
}
func add(a int, b int) (out int) {
	return a + b
}
func main() {

	beego.AddFuncMap("add", add)
	beego.AddFuncMap("toHtml", toHtml)
	//绑定全局错误处理控制器
	beego.ErrorController(&controllers.ErrorController{})

	if beego.AppConfig.String("RunMode") == "dev" {
		beego.Run()
	} else {
		beego.Run("0.0.0.0:5050")
	}

}
