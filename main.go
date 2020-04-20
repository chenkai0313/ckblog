package main

import (
	_ "ckblog/routers"
	"ckblog/comment"
	"github.com/astaxie/beego"
	"ckblog/models"
)





func main() {
	//初始化session
	comment.InitSession()
	// 初始化数据库
	models.InitMysql()

	beego.SetStaticPath("/static","static")


	// 如果是开发模式，则显示命令信息
	//isDev := !(beego.AppConfig.String("runmode") != "dev")
	//if isDev {
	//	orm.Debug = isDev
	//}

	beego.Run()
}

