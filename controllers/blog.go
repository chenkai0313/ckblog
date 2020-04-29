package controllers

import (
	"github.com/astaxie/beego"
	"ckblog/comment"
)

type BlogController struct {
	beego.Controller
}


//@router /blog/article/blogCategory
func (c *ArticleController) BlogCategory() {
	categoryList:=articlesService.CategoryList()
	delete(categoryList,0)
	c.Data["json"] = response.JsonFormat(comment.CODE_SUCCESS, categoryList, "success")
	c.ServeJSON()
}


//@router /blog/article/blogIndex
func (c *ArticleController) BlogIndex() {

}
