package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() {
	link := fmt.Sprintf("%s?charset=utf8", beego.AppConfig.String("mysqlDsn"))
	beego.Info("mysql init.....", link)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", link, 30, 30)
	orm.RegisterModel(new(User))
}