package main

import (
	_ "ckblog/routers"
	"ckblog/comment"
	"github.com/astaxie/beego"
	"ckblog/models"
	"ckblog/controllers"
)




func main() {
	//init session
	comment.InitSession()
	//init mysql
	models.InitMysql()

	beego.SetStaticPath("/static","static")


     //var logger comment.Logger
	//logger.Infof("server end ","1.0","1234")
	// 如果是开发模式，则显示命令信息
	//isDev := !(beego.AppConfig.String("runmode") != "dev")
	//if isDev {
	//	orm.Debug = isDev
	//}

	//error page
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}

