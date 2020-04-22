package routers

import (
	"github.com/astaxie/beego"
	"ckblog/controllers"
	"github.com/astaxie/beego/context"
	"fmt"
	"ckblog/comment"
)

func filterRoutes(nowRequestUrl string) bool {
	routes := [] string{
		"/backend/user/login",
		"/backend/user/loginAct",
	}
	for _, v := range routes {
		if v == nowRequestUrl {
			return true
		}
	}
	return false
}

//过滤未登陆的用户
var FilterUser = func(ctx *context.Context) {
	ExistUserSession := ctx.Input.Session(comment.SESSION_NAME)
	fmt.Println("ExistUserSession", ExistUserSession)
	nowRequestUrl := ctx.Request.RequestURI
	filterRoute := filterRoutes(nowRequestUrl)
	if !filterRoute {
		if ExistUserSession == nil {
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
func userLoginRouter() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	//注解路由
	beego.Include(
		&controllers.UserController{},
		&controllers.SiteController{},
	)
}
