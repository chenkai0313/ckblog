package service

import (
	"ckblog/models"
	"strconv"
)

type ArticleService struct{}

type ArticlesLists []ArticlesList
type ArticlesList struct {
	Id               int
	CategoryId       int
	CategoryIdString string
	UserId           int
	Sort             int
	IsDisplay        int
	IsDisplayString  string
	Title            string
	Content          string
	CreatedTime      string
	UpdatedTime      string
}

func CategoryList() map[int]string {
	category := make(map[int]string)
	category[0] = "全部"
	category[1] = "Golang"
	category[2] = "PHP"
	category[3] = "Linux"
	return category
}

func IsDisplayList() map[int]string {
	isDisplay := make(map[int]string)
	isDisplay[0] = "全部"
	isDisplay[1] = "发布"
	isDisplay[2] = "不发布"
	isDisplay[3] = "草稿"
	return isDisplay
}

func stringInterception(str string,length int)  string{
	if len(str) <length{
		return  str
	}else {
		return  str[0 : length]+"。。。。"
	}
}

func (article ArticleService) GetArticleList(params map[string]string, pageSize int, pageNow int) []ArticlesList {
	var articlesList []ArticlesList
	articles := models.GetArticlesListByParams(params, pageSize, pageNow)
	categoryList := CategoryList()
	isDisplayList := IsDisplayList()

	for _, v := range articles {
		rl := ArticlesList{
			Id:               v.Id,
			CategoryId:       v.CategoryId,
			CategoryIdString: categoryList[v.CategoryId],
			UserId:           v.UserId,
			Sort:             v.Sort,
			IsDisplay:        v.IsDisplay,
			IsDisplayString:  isDisplayList[v.IsDisplay],
			Title:            stringInterception(v.Title,60),
			Content:          v.Content,
			CreatedTime:      v.CreatedTime,
			UpdatedTime:      v.UpdatedTime,
		}
		articlesList=append(articlesList,rl)
	}

	return articlesList
}

func (article ArticleService) GetArticleListCount(params map[string]string) int {

	articlesCounts := models.GetArticlesCountByParams(params)
	articlesCountsInt, _ := strconv.Atoi(strconv.FormatInt(articlesCounts, 10))
	return articlesCountsInt
}
