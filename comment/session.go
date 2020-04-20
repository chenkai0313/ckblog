package comment

import (
	//_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
	"blog/comment/enum"
)


func InitSession(){
	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionName = enum.LocalSessionName
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = enum.SessionExpireTime
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = enum.SessionExpireTime

	//is, _ := pubMethod.PathExists(enum.SessionPath)
	//if !is {
	//	_ = os.Mkdir(enum.SessionPath, os.ModePerm)
	//}
	beego.BConfig.WebConfig.Session.SessionProviderConfig = enum.SessionPath


}