package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"ckblog/comment"
	"github.com/weekface/mgorus"
	"fmt"
)


type BaseController struct {
	beego.Controller
}

var Logger comment.Logger

func init(){
	logger := logrus.New()
	hooker, err := mgorus.NewHooker("mongodb://root:123456@11.11.11.1127017/?authSource=admin", "ckblog", "sys_log")
	if err == nil {
		logger.Hooks.Add(hooker)
	} else {
		fmt.Println("log:", err)
	}
	Logger = comment.NewLogger(logger,logrus.Fields{})
}