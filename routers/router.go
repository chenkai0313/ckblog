package routers

import (
	"github.com/astaxie/beego"
	"ckblog/controllers"
)

//过滤未登陆的用户
//var FilterUser = func(ctx *context.Context){
//	_, ok := ctx.Input.Session("uid").(string)
//	ok2 := strings.Contains(ctx.Request.RequestURI, "/backend/login")
//	ok3 := strings.Contains(ctx.Request.RequestURI, "/backend/register")
//	if !ok && !ok2 && !ok3{
//		ctx.Redirect(302, "/backend/login")
//	}
//}

func init() {
	//固定路由
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/backend/user/login", &controllers.UserController{},"*:Login")
    //beego.Router("/backend/register", &controllers.UserController{},"*:Register")
	userLoginRouter()
}

//需要登陆才能访问的路由
func userLoginRouter()  {
	//beego.InsertFilter("/backend/*", beego.BeforeRouter, FilterUser)
	//注解路由
	beego.Include(
		&controllers.UserController{},
		&controllers.SiteController{},
	)
}