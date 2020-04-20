package service

import (
	"ckblog/untils"
	"ckblog/models"
	"errors"
)


type UserService  struct{}


//注册
func(login UserService)Register(user models.User)(bool,error,models.User){
	//判断用户是否存在
	var ExistUser models.User
	ExistUser =models.GetUserByUserName(user.UserName)
	if ExistUser != (models.User{}){
		return  false,errors.New("user name exist"),models.User{}
	}

	ExistUser =models.GetUserByEmail(user.Email)
	if ExistUser != (models.User{}){

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
