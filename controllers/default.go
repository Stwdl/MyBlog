package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

//db.Using("default")  切换数据库实体，可以定义多个数据库的呀，用这种方式切换
func (c *MainController) Get() {
	db := orm.NewOrm()
	var images []*models.Image
	var articles []*models.Article
	_, err := db.QueryTable(new(models.Image)).Filter("IsShowHome", true).OrderBy("-Createtime").Limit(8).All(&images)
	if err == nil {
		c.Data["images"] = images
	} else {
		c.Data["images"] = []models.Image{{Url: "../static/img/DSC_0845.jpg"}, {Url: "../static/img/DSC_0861.jpg"}, {Url: "../static/img/DSC_0870.jpg"}, {Url: "../static/img/DSC_0871.jpg"}, {Url: "../static/img/DSC_0888.jpg"}}

	}
	_, err = db.QueryTable(new(models.Article)).Filter("Type", 1).OrderBy("-CreateTime").Limit(8).All(&articles)
	if err == nil {
		c.Data["articles"] = articles
	}
	var DTS []*models.Article
	_, err = db.QueryTable(new(models.Article)).OrderBy("-CreateTime").Limit(10).All(&DTS)

	//_, err = db.Raw("select ID,Title,Content,CreateTime,Type from Article order by CreateTime desc limit ?", 10).QueryRows(&DTS)
	if err == nil {
		c.Data["DTS"] = DTS
	} else {
		c.Data["DTS"] = nil
	}
	var MSG []*models.Message
	_, err = db.QueryTable(new(models.Message)).OrderBy("-CreateTime").Limit(10).All(&MSG)
	if err == nil {
		c.Data["MSG"] = MSG
	} else {
		c.Data["MSG"] = nil
	}
	c.Data["TianQi"] = "汕头，天晴，心情阴郁"
	c.TplName = "Home.tpl"
}
