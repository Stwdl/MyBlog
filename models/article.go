package models

import (
	"time"
)

///文章
type Article struct {
	//index:添加索引
	ID int64 `orm:"index;column(ID)"`
	//自动添加当前时间type指定带不带时间不带的话用date
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(CreateTime)"`
	//size指定字符串长度
	Title string `orm:"size(200);column(Title)"`
	//文章内容
	Content string
	//作者；default设置默认值;column指定映射到数据库的列名
	Autor string `orm:"default(醉翁二郎);column(Autor)"`

	///访问数：总长度，精度orm:"digits(12);decimals(4)"
	ViewCount int64 `orm:"column(ViewCount)"`
	//类型 1：文章 2：说说
	Type int32 `orm:"column(Type)"`
}

func (a *Article) TableName() string {
	return "Article"
}

///专栏类型
type ArticleType struct {
	ID         int32     `orm:"column(ID)"`
	Name       string    `orm:"column(Name)"`
	Createtime time.Time `orm:"auto_now_add;type(datetime);column(CreateTime)"`
}

func (a *ArticleType) TableName() string {
	return "ArticleType"
}
