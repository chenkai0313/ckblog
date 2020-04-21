package routers

import (
	"github.com/astaxie/beego"
	"ckblog/controllers"
	"github.com/astaxie/beego/context"
	"fmt"
)

//过滤未登陆的用户
var FilterUser = func(ctx *context.Context) {
	data, ok := ctx.Input.Session("uid").(string)
	fmt.Println("data",data)
	if ctx.Request.RequestURI != "/backend/user/login" && ctx.Request.RequestURI !="/backend/user/loginAct"{
		if !ok   {
			ctx.Redirect(302, "/backend/user/login")
		}
	}

}

func init() {
	//固定路由
	//beego.Router("/",  &controllers.UserController{},"*:Login")
	//beego.Router("/backend/user/login", &controllers.UserController{},"*:Login")
    //beego.Router("/backend/register", &controllers.UserController{},"*:Register")
	userLoginRouter()
}

//需要登陆才能访问的路由
func userLoginRouter()  {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	//注解路由
	beego.Include(
		&controllers.UserController{},
		&controllers.SiteController{},
	)
}