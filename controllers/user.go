package controllers

import (
	"github.com/astaxie/beego"
	"ckblog/models"
	"ckblog/validate"
	"fmt"
	"ckblog/service"
)

type UserController struct {
	beego.Controller
}

var userService service.UserService

//@router /backend/user/register
func (c *UserController) Register() {
	var newUser models.User
	newUser.Email = "as123d@qq.com"
	newUser.Password = "123456"
	newUser.UserName = "admin123"
	var userValidate validate.UserValidate

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
	flash := beego.NewFlash()
	c.Data["Website"] = "beego.me"
	if c.Ctx.Input.IsPost() {
		userName := c.GetString("userName")
		password := c.GetString("password")
		//判断session是否存在
		us := c.GetSession("login_session")
		if us == nil {
			if userService.Login(userName,password){
				//设置session
				c.SetSession(userName+"session"+password, "session")
				c.Redirect("/backend/site/index",302)
				return
			}else {
				flash.Error("帐号密码错误")
				flash.Store(&c.Controller)
				c.Data["userName"] = userName
				c.Data["password"] = password

				c.TplName = "login/login.html"
				return
			}
		}else {
			c.Redirect("/backend/site/index",302)
			return
		}
	}
	c.TplName = "login/login.html"
	return
}


//@router /backend/user/out
func (c *UserController) Out() {

	c.DelSession("uid")

	c.TplName = "login/out.html"
}
