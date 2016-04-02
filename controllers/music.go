package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MusicController struct {
	beego.Controller
}

func (m *MusicController) Get() {
	db := orm.NewOrm()
	var musics []*models.Music
	num, err := db.QueryTable(new(models.Music)).Filter("IsOpen", true).OrderBy("-LoveCount").Limit(10).All(&musics)
	if err == nil && num > 0 {
		m.Data["musics"] = musics
	} else {
		m.Data["musics"] = []models.Music{{Name: "谁能明白我", Singer: "林子祥", Url: "http://win.web.rc03.sycdn.kuwo.cn/0f96f9df2b150c9a81d465979328ae02/56fffc81/resource/a2/65/19/1257103968.aac"}}
	}
	m.TplName = "music.tpl"
}
