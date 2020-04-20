package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"strconv"
	"ckblog/untils"
	"fmt"
)

const USER_TABLE_NAME = "user"

type User struct {
	Id          int
	UserName    string
	Email       string
	Password    string
	Token       string
	CreatedTime string
	UpdatedTime string
}

func InsertUser(user User) (bool, User) {
	var userInfo User
	o := orm.NewOrm()
	nowTime:=untils.GetBasicDateTime()
	user.CreatedTime=nowTime
	user.UpdatedTime=nowTime
	userId, err := o.Insert(&user)
	if err != nil {
		logs.Error("insert user fail: ", err)
		return false, userInfo
	}
	userIdInt, _ := strconv.Atoi(strconv.FormatInt(userId, 10))
	userInfo = GetUserByUserId(userIdInt)
	return true, userInfo
}

func GetUserByUserName(userName string) User {
	o := orm.NewOrm()
	qs := o.QueryTable(USER_TABLE_NAME)
	var userInfo User
	err := qs.Filter("user_name", userName).One(&userInfo)
	if err != nil {
		logs.Error("get user info by user_name fail: ", err)
	}

	fmt.Println("user",userInfo)
	return userInfo
}

func GetUserByEmail(email string) User {
	o := orm.NewOrm()
	qs := o.QueryTable(USER_TABLE_NAME)
	var userInfo User
	err := qs.Filter("email", email).One(&userInfo)
	if err != nil {
		logs.Error("get user info by user_name fail: ", err)
	}
	return userInfo
}

func GetUserByUserId(userId int) User {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable(USER_TABLE_NAME)
	var userInfo User
	err := qs.Filter("id", userId).One(&userInfo)
	if err != nil {
		logs.Error("get user info by id fail: ", err)
	}
	return user
}

func UpdateUserById(user User) (bool, User) {
	var userInfo User
	user.UpdatedTime=untils.GetBasicDateTime()
	o := orm.NewOrm()
	if userId, err := o.Update(&user); err != nil {
		logs.Error("update user info fail: ", err)
		return false, userInfo
	} else {
		userIdInt, _ := strconv.Atoi(strconv.FormatInt(userId, 10))
		userInfo = GetUserByUserId(userIdInt)
		return false, userInfo
	}
}
