package models

type Music struct {
	ID int64 `orm:"column(ID)"`
	//音乐名
	Name string `orm:"column(Name)"`
	//作者信息
	Singer string `orm:"column(Singer)"`
	///地址
	Url string `orm:"column(Url)"`
	//是否本地音乐
	IsLocation bool `orm:"column(IsLocation)"`
	//是否公开
	IsOpen bool `orm:"column(IsOpen)"`
	//所属音乐专辑
	AlubmID int `orm:"column(AlubmID)"`
	//喜欢计数
	LoveCount int64 `orm:"column(LoveCount)"`
	//讨厌计数
	HateCount int64 `orm:"column(HateCount)"`
	///播放数
	PlayCount int64 `orm:"column(PlayCount)"`
	//采集来源
	Form string `orm:"column(Form)"`
}

func (a *Music) TableName() string {
	return "Music"
}

///音乐专辑
type MusicAlbum struct {
	ID     int    `orm:"column(ID)"`
	Name   string `orm:"column(Name)"`
	IsOpen bool   `orm:"column(IsOpen)"`
	//喜欢计数
	LoveCount int64 `orm:"column(LoveCount)"`
	//讨厌计数
	HateCount int64 `orm:"column(HateCount)"`
	///访问量
	ViewCount int64 `orm:"column(ViewCount)"`
}

func (a *MusicAlbum) TableName() string {
	return "MusicAlbum"
}
