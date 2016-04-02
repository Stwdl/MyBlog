package models

import (
	"time"
)

///图片
type Image struct {
	ID int64 `orm:"column(ID)"`
	///名字
	Name string `orm:"column(Name)"`
	///图片地址
	Url string `orm:"column(Url)"`
	///描述
	Descript string `orm:"column(Descript)"`
	///拍照地点？
	Where string `orm:"column(Where)"`
	///上传时间
	Createtime time.Time `orm:"auto_now_add;type(datetime);column(CreateTime)"`
	//拍照的时间？通过后台解析图片还是？
	Phototime time.Time `orm:"date(datetime);column(Phototime)"`
	//图片是否在本地
	IsLocation bool `orm:"column(IsLocation)"`
	//是否首页展示
	IsShowHome bool `orm:"column(IsShowHome)"`
	///所属相册ID
	AlbumID int `orm:"column(AlbumID)"`
}

func (a *Image) TableName() string {
	return "Image"
}

///相册
type Album struct {
	ID int `orm:"column(ID)"`
	//相册名称
	Name string `orm:"column(AlbumID)"`
	//描述
	Descript string `orm:"column(Name)"`
	//分类ID
	TypeID int `orm:"column(TypeID)"`
	//是否公开
	IsOpen bool `orm:"column(IsOpen)"`
}

func (a *Album) TableName() string {
	return "Album"
}

///相册分类
type AlbumType struct {
	ID   int    `orm:"column(ID)"`
	Name string `orm:"column(Name)"`
}

func (a *AlbumType) TableName() string {
	return "AlbumType"
}
