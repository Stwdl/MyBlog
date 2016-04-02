package models

type User struct {
	Id       int64  `orm:"column(ID)"`
	Name     string `orm:"size(100);column(Name)"`
	Email    string `orm:"column(Email)"`
	LoginID  string `orm:"column(LoginID)"`
	Password string `orm:"column(Password)"`
	Mobile   string `orm:"column(Mobile)"`

	IsAdmin bool `orm:"default(false);column(IsAdmin)"`
}

//定制表名
func (a *User) TableName() string {
	return "User"
}
