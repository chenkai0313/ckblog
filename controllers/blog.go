package controllers

import (
	"github.com/astaxie/beego"
	"ckblog/comment"
	"ckblog/service"
)

type BlogController struct {
	beego.Controller
}


//@router /blog/article/blogCategory
func (c *ArticleController) BlogCategory() {
	categoryList := articlesService.CategoryList()
	delete(categoryList, 0)
	c.Data["json"] = response.JsonFormat(comment.CODE_SUCCESS, categoryList, "success")
	c.ServeJSON()
}

type BlogIndex struct {
	Title      string
	PageNow    int
	CategoryId int
	Count      int
	List       []service.ArticlesList
}

//@router /blog/article/blogIndex
func (c *ArticleController) BlogIndex() {

	var pageSize = 10
	pageNow, err := c.GetInt("pageNow")
	if err != nil || pageNow == 0 {
		pageNow = 1
	}

	params := make(map[string]string)
	params["category_id"] = c.GetString("category_id")
	params["is_display"] = "1"
	params["title"] = c.GetString("title")

	articlesCounts := articlesService.GetArticleListCount(params)
	articles := articlesService.GetArticleList(params, pageSize, pageNow)


	responseList:=BlogIndex{
		Title: params["title"],
		PageNow: pageNow,
		CategoryId: pageNow,
		Count: articlesCounts,
		List : articles,
	}


	c.Data["json"] = response.JsonFormat(comment.CODE_SUCCESS, responseList, "success")
	c.ServeJSON()

}
