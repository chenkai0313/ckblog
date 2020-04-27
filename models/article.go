package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"strconv"
	"ckblog/untils"
	"fmt"
)

const Article_TABLE_NAME = "article"

type Article struct {
	Id          int
	CategoryId  int
	UserId      int
	Sort        int
	IsDisplay   int
	Title       string
	Content     string
	CreatedTime string
	UpdatedTime string
}

func GetArticleById(id int) Article {
	var articleInfo Article
	o := orm.NewOrm()
	qs := o.QueryTable(Article_TABLE_NAME)
	err := qs.Filter("id", id).One(&articleInfo)
	if err != nil {
		logs.Error("get article info by id fail: ", err)
	}
	return articleInfo
}

func UpdateArticle(article Article) (bool, error,Article) {
	var articleInfo Article
	o := orm.NewOrm()
	nowTime := untils.GetBasicDateTime()
	article.UpdatedTime = nowTime
	artId, err := o.Update(&article)
	if err != nil {
		logs.Error("update article fail: ", err)
		return false, err,articleInfo
	}
	artIdInt, _ := strconv.Atoi(strconv.FormatInt(artId, 10))
	articleInfo = GetArticleById(artIdInt)
	return true, nil,articleInfo
}

func InsertArticle(article Article) (bool, error,Article) {
	var articleInfo Article
	o := orm.NewOrm()
	nowTime := untils.GetBasicDateTime()
	article.UpdatedTime = nowTime
	artId, err := o.Insert(&article)
	if err != nil {
		logs.Error("insert article fail: ", err)
		return false, err,articleInfo
	}
	artIdInt, _ := strconv.Atoi(strconv.FormatInt(artId, 10))
	articleInfo = GetArticleById(artIdInt)
	return true, nil,articleInfo
}

func GetArticlesByParams(params map[string]string) []Article {
	var articles []Article
	o := orm.NewOrm()
	qs := o.QueryTable(Article_TABLE_NAME)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	if _, err := qs.Limit(-1).All(&articles); err != nil {
		logs.Error("get article settle list fail: ", err)
	}
	return articles
}

func GetArticlesListByParams(params map[string]string, pageSize int, pageNow int) []Article {
	var articles []Article
	o := orm.NewOrm()
	qs := o.QueryTable(Article_TABLE_NAME)
	for k, v := range params {
		if len(v) > 0 {
			if k == "title" && v != "" {
				qs = qs.Filter("title__icontains", v)
			}
			if k == "is_display" && v != "0" {
				qs = qs.Filter(k, v)
			}
			if k == "category_id" && v != "0" {
				qs = qs.Filter(k, v)
			}
			if k == "sort" && v != "0" {
				if v == "1" {
					qs = qs.OrderBy("sort")
				}
				if v == "2" {
					qs = qs.OrderBy("-sort")
				}
			}
		}
	}
	if _, err := qs.Limit(pageSize, (pageNow-1)*pageSize).OrderBy("-id").All(&articles); err != nil {
		logs.Error("get article settle list fail: ", err)
	}
	return articles
}

func GetArticlesCountByParams(params map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(Article_TABLE_NAME)

	for k, v := range params {
		if len(v) > 0 {
			if k == "title" && v != "" {
				qs = qs.Filter("title__icontains", v)
			}
			if k == "is_display" && v != "0" {
				qs = qs.Filter(k, v)
			}
			if k == "category_id" && v != "0" {
				qs = qs.Filter(k, v)
			}
			if k == "sort" && v != "0" {
				if v == "1" {
					qs = qs.OrderBy("sort")
				}
				if v == "2" {
					qs = qs.OrderBy("-sort")
				}
			}
		}
	}
	var counts int64
	if counts, err := qs.Count(); err != nil {
		logs.Error("get article settle list fail: ", err)
	} else {
		return counts
	}
	return counts
}

func DelArticleById(id int) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(Article_TABLE_NAME)
	qs = qs.Filter("id", id)
	if _, err := qs.Delete(); err != nil {
		fmt.Println("deletedeletedeletedelete error",err)
		logs.Error("delete article by id  fail ", err)
		return false
	}
	return true
}
