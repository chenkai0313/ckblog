package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
	"blog/validate"
	"fmt"
	"blog/service"
)

type UserController struct {
	beego.Controller

}

var userService service.UserService

//@router /backend/user/register
func (c *UserController) Register() {
	var newUser models.User
	newUser.Email="as123d@qq.com"
	newUser.Password="123456"
	newUser.UserName="admin123"
	var userValidate validate.UserValidate

	if err:=userValidate.RegisterValite(newUser);err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("success")
	}

	if boolRes,err,userinfo:=userService.Register(newUser);!boolRes{
		fmt.Println("登陆错误",err)
	}else {
		fmt.Println("注册成功",userinfo)
	}

	c.TplName = "index.html"
}




//@router /backend/user/login
func (c *UserController) Login() {
	//c.SetSession("uid", "当前登陆人是 ck")

	c.TplName = "login/login.html"
}




//@router /backend/user/out
func (c *UserController) Out() {

	c.DelSession("uid")

	c.TplName = "login/out.html"
}