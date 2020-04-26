package controllers

import (
	"github.com/astaxie/beego"
	"ckblog/service"
	"ckblog/untils"
	"html/template"
	"ckblog/comment"
	"ckblog/validate"
	"strconv"
	"ckblog/models"
)

type ArticleController struct {
	beego.Controller
}

var articlesService= service.ArticleService{}
var response = comment.Response{}
var articleValidate = validate.ArticleValidate{}


//@router /backend/article/backendIndex
func (c *ArticleController) BackendIndex() {

	flashGet:=beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	if n,ok:=flashGet.Data["notice"];ok{
		flash.Notice(n)
	}else if n,ok=flashGet.Data["error"];ok{
		flash.Error(n)
	}else{
		flash.Warning(n)
	}

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
	CategoryList:=articlesService.CategoryList()
	delete(CategoryList,0)
	c.Data["CategoryType"]=CategoryList
	IsDisplayList:=articlesService.IsDisplayList()
	delete(IsDisplayList,0)
	c.Data["IsDisplayType"]=IsDisplayList
	c.TplName = "article/list.html"
}

//@router /backend/article/backendArticleDel
func (c *ArticleController) BackendArticleDel() {
	flash := beego.NewFlash()
	id, err := c.GetInt("id")
	if err != nil {
		flash.Error("Params error")
	}
	if !articlesService.DelArtilceById(id){
		flash.Error("Delete fail")
	}else {
		flash.Notice("success")
	}

	flash.Store(&c.Controller)
	c.Redirect("/backend/article/backendIndex",302)

	return
}


//@router /backend/article/backendArticleAdd [post]
func (c *ArticleController) BackendArticleAdd() {
	flash := beego.NewFlash()
	var article models.Article
	article.Title=c.GetString("post_title")
	article.Content=c.GetString("post_content")
	article.CreatedTime=c.GetString("post_created_time")
	PostCategoryIdInt, _ := strconv.Atoi(c.GetString("post_category_id"))
	article.CategoryId=PostCategoryIdInt
	PostIsDisplayInt, _ := strconv.Atoi(c.GetString("post_is_display"))
	article.IsDisplay=PostIsDisplayInt
	PostSortInt, _ := strconv.Atoi(c.GetString("post_sort"))
	article.Sort=PostSortInt

	if err := articleValidate.AddArticle(article); err != nil {
		flash.Error(err.Error())
		c.Redirect("/backend/article/backendIndex",302)
		return
	}
	if resBool,err:=articlesService.AddArticle(article);!resBool{
		flash.Error(err.Error())
		c.Redirect("/backend/article/backendIndex",302)
		return
	}
	flash.Notice("success")
	flash.Store(&c.Controller)
	c.Redirect("/backend/article/backendIndex",302)
	return

}

////@router /backend/article/backendArticleAdd [post]
//func (c *ArticleController) BackendArticleEdit() {
//	c.TplName = "article/edit.html"
//}