package controllers

import (
	"github.com/astaxie/beego"
	"ckblog/comment"
	"ckblog/models"
)

type SiteController struct {
	beego.Controller

}

//@router /backend/site/index
func (c *SiteController) Index() {
	userSession := c.GetSession(comment.SESSION_NAME)
	userInfo:=userSession.(models.User)
	c.Data["userInfo"]=userInfo
	c.TplName = "index.html"
}

//@router /backend/site/site
func (c *SiteController) Site() {
	c.TplName = "site/site.html"
}

//@router /backend/site/test
func (c *SiteController) Test() {
	c.TplName = "site/site.html"
}