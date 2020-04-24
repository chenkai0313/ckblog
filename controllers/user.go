package controllers

import (
	"github.com/astaxie/beego"
	"ckblog/models"
	"ckblog/validate"
	"fmt"
	"ckblog/service"
	"ckblog/comment"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/cache"
)

type UserController struct {
	beego.Controller
}

var cpt *captcha.Captcha
func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}
var userService=service.UserService{}
var userValidate=validate.UserValidate{}


//@router /backend/user/register
func (c *UserController) Register() {
	var newUser models.User
	newUser.Email = "as123d@qq.com"
	newUser.Password = "123456"
	newUser.UserName = "admin123"

	if err := userValidate.RegisterValite(newUser); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}


	if boolRes, err, userinfo := userService.Register(newUser); !boolRes {
		fmt.Println("登陆错误", err)
	} else {
		fmt.Println("注册成功", userinfo)
	}

	c.TplName = "index.html"
	return
}

//@router /backend/user/login
func (c *UserController) Login() {
	c.TplName = "login/login.html"
	return
}

//@router /backend/user/loginAct [post]
func (c *UserController) LoginAct() {
		flash := beego.NewFlash()
		userName := c.GetString("userName")
		password := c.GetString("password")
		captcha := c.GetString("captcha")
		c.Data["userName"] = userName
		c.Data["password"] = password
		c.Data["captcha"] = captcha

		if !cpt.VerifyReq(c.Ctx.Request) {
			flash.Error("Verification code error")
			flash.Store(&c.Controller)
			c.TplName = "login/login.html"
			return
		}

		//determine if the session exists
		us := c.GetSession(comment.SESSION_NAME)
		if us == nil {
			if resBool,user:=userService.Login(userName,password);resBool{
				//set session
				c.SetSession(comment.SESSION_NAME,user)
				c.Redirect("/backend/site/index",302)
				c.Data["uid"]=user
				return
			}else {
				flash.Error("account or password incorrect")
				flash.Store(&c.Controller)
				c.TplName = "login/login.html"
				return
			}
		}else {
			c.Redirect("/backend/site/index",302)
			return
		}
}


//@router /backend/user/out
func (c *UserController) Out() {
	c.DelSession(comment.SESSION_NAME)
	c.Redirect("/backend/user/login",302)
	return

}
