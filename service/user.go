package service

import (
	"ckblog/untils"
	"ckblog/models"
	"errors"
	"strings"
)


type UserService  struct{}


//注册
func(login *UserService)Register(user models.User)(bool,error,models.User){
	//判断用户是否存在
	var ExistUser models.User
	ExistUser =models.GetUserByUserName(user.UserName)
	if ExistUser.Id != 0 {
		return  false,errors.New("user name exist"),models.User{}
	}

	ExistUser =models.GetUserByEmail(user.Email)
	if ExistUser.Id !=0 {

		return  false,errors.New("user email exist"),models.User{}
	}

	var encrypt =untils.Encrypt{}
	userPassword:=encrypt.EncodeMd5([]byte(user.Password))

	randString:=untils.RandomString(8)
	rangToken:=encrypt.EncodeMd5([]byte(randString))

	newUser:=models.User{UserName:user.UserName,Email:user.Email,Password:userPassword,Token:rangToken}

	if resBol,userInfo:=models.InsertUser(newUser); !resBol{
		return true,nil,userInfo
	} else {
		return true,nil,userInfo
	}
}

//登陆
func (login *UserService)Login(userName,password string)(bool,models.User){
	var ExistUser models.User
	ExistUser =models.GetUserByUserName(userName)
	if ExistUser.Id ==0 {
		return  false,ExistUser
	}
	var encrypt =untils.Encrypt{}
	pwdMd5 := encrypt.EncodeMd5([]byte(password))
	if strings.Compare(pwdMd5, ExistUser.Password) != 0 {
		return  false,ExistUser
	}else {
		return true,ExistUser
	}

	return false,ExistUser
}
