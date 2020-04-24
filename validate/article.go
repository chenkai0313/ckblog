package validate

import (
	"github.com/astaxie/beego/validation"
	"errors"
	"ckblog/models"
)

type ArticleValidate struct{}

//type AddArticle struct {
//	CategoryId  int
//	Sort        int
//	IsDisplay   int
//	Title       string
//	Content     string
//	CreatedTime string
//}

func (article *ArticleValidate) DelArticleById(id int) {

}


func (article *ArticleValidate) AddArticle(art models.Article) error{
	valid := validation.Validation{}
	valid.Required(art.CategoryId, "category")
	valid.Required(art.Sort, "sort")
	valid.Required(art.IsDisplay, "display")
	valid.Required(art.Title, "title")
	valid.Required(art.Content, "content")
	valid.Required(art.CreatedTime, "created_time")

	valid.MaxSize(art.Title, 255,"max title")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return  errors.New(err.Key+" "+err.Message)
		}
	}
	return nil
}