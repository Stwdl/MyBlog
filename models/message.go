package models

import (
	"time"
)

type Message struct {
	ID int64  `orm:"column(ID)"`
	IP string `orm:"column(IP)"`
	///来源 1：访客 2：主人
	TypeID     int       `orm:"column(TypeID)"`
	Name       string    `orm:"column(Name)"`
	Content    string    `orm:"column(Content)"`
	Createtime time.Time `orm:"auto_date_add;type(datetime);column(Createtime)"`
}

func (a *Message) TableName() string {
	return "Message"
}
