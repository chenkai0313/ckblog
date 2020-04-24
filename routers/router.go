package routers

import (
	"github.com/astaxie/beego"
	"ckblog/controllers"
	"github.com/astaxie/beego/context"
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

//filter unlogged users
var FilterUser = func(ctx *context.Context) {
	ExistUserSession := ctx.Input.Session(comment.SESSION_NAME)
	nowRequestUrl := ctx.Request.RequestURI
	filterRoute := filterRoutes(nowRequestUrl)
	if !filterRoute {
		if ExistUserSession == nil {
			ctx.Redirect(302, "/backend/user/login")
		}
	}
}

func init() {
	//beego.Router("/",  &controllers.UserController{},"*:Login")
	//beego.Router("/backend/user/login", &controllers.UserController{},"*:Login")
	//beego.Router("/backend/register", &controllers.UserController{},"*:Register")
	userLoginRouter()
}

func userLoginRouter() {
	//beego.InsertFilter("/backend/*", beego.BeforeRouter, FilterUser)
	beego.Include(
		&controllers.UserController{},
		&controllers.SiteController{},
		&controllers.ArticleController{},
	)
}
