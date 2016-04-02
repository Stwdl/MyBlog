package models

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	dbtype := beego.AppConfig.String("dbtype")
	log.Printf("数据库类型为：%s", dbtype)
	if dbtype == "mysql" {
		dbhost := beego.AppConfig.String("dbhost")         //数据库主机地址
		dbport := beego.AppConfig.String("dbport")         //端口
		dbuser := beego.AppConfig.String("dbuser")         //用户名
		dbpassword := beego.AppConfig.String("dbpassword") //密码
		dbname := beego.AppConfig.String("dbname")         //主机名
		if dbport == "" {
			dbport = "3306"
		}
		dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8" //他妈的竟然是手动指定的地址

		//注册实体数据库框架
		orm.RegisterDataBase("default", "mysql", dburl)
	} else if dbtype == "sqlite3" {
		sqlitename := beego.AppConfig.String("sqlitename")
		orm.RegisterDataBase("default", "sqlite3", sqlitename)
	}

	//注册数据库实体
	orm.RegisterModel(new(User), new(Article), new(Image), new(Album), new(AlbumType), new(Music), new(MusicAlbum), new(Message))
	//创建数据库表
	log.Printf("开始创建数据库")
	orm.RunSyncdb("default", false, true)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func MD5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%s", hash.Sum(nil))
}
