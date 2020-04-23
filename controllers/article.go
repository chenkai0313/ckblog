package controllers

import (
	"github.com/astaxie/beego"
	"ckblog/service"
	"ckblog/untils"
	"html/template"
)

type ArticleController struct {
	beego.Controller
}

var articlesService service.ArticleService

//@router /backend/article/backendIndex
func (c *SiteController) BackendIndex() {

	var pageSize = 10

	pageNow, err := c.GetInt("pageNow")
	if err != nil || pageNow==0{
		pageNow = 1
	}

	params := make(map[string]string)
	params["category_id"]=c.GetString("category_id")
	params["is_display"]=c.GetString("is_display")
	params["sort"]=c.GetString("sort")
	params["title"]=c.GetString("title")

	articlesCounts := articlesService.GetArticleListCount(params)
	articles := articlesService.GetArticleList(params, pageSize, pageNow)
	c.Data["Data"] = articles
	pagination:=untils.Pagination(articlesCounts,pageNow,pageSize,params)
	c.Data["Pagination"]=template.HTML(pagination)


	c.TplName = "article/list.html"
}
