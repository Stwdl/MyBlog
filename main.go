package main

import (
	//	"MyBlog/models"
	_ "MyBlog/routers"

	"github.com/astaxie/beego"
)

func main() {

	//	models.Init()
	//	o := orm.NewOrm()
	//	user := models.User{Name: "李逍遥"}
	//	//插入数据
	//	id, err := o.Insert(&user)
	//	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	//	user.Name = "赵灵儿"
	//	num, err := o.Update(&user)
	//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//	u := models.User{Id: user.Id}
	//	err = o.Read(&u)
	//	fmt.Printf("ERR: %v\n", err)

	//	// delete
	//	num, err = o.Delete(&u)
	//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	beego.SetStaticPath("/assets/", "static/assets/")
	beego.Run()
}
