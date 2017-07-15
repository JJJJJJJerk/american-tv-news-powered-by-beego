package controllers

import (
	"encoding/json"
	"fmt"
	"my_go_web/models"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Index() {

	articles := []models.Article{}
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Preload("Vote").Preload("Tags").Preload("Images").Find(&articles)

	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/article", "fa fa-home", "资讯"}}
	c.Data["Articles"] = articles
	c.Data["Keyword"] = "美剧keywords"
	c.Data["Description"] = "美剧描述"
	c.Data["Title"] = "美剧资讯"

	c.Layout = "layout/base.html"
	c.TplName = "article/index.html"
}

func (c *ArticleController) View() {
	articleID, _ := c.GetInt(":id")
	//浏览计数
	vote := models.Vote{}
	models.Gorm.Where("article_id = ?", articleID).Find(&vote)
	vote.Visit++
	models.Gorm.Model(&vote).Update("visit")

	article := models.Article{}
	if models.Gorm.Preload("Tags").Preload("Images").Preload("Shows").First(&article, articleID).RecordNotFound() {
		c.Abort("404")
	}

	//设置head seo参数
	//设置breadcrumb
	//设置side bar
	//设置head navigation bar
	url := fmt.Sprint("/article/", articleID)
	tagUrl := fmt.Sprint("/tag/", article.FirstTagID)
	tagName := fmt.Sprint(article.FirstTagName, article.FirstTagNameEn)
	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {tagUrl, "fa fa-book", tagName}, {url, "fa fa-cloud", article.Title}}
	c.Data["Article"] = article
	c.Data["Vote"] = vote
	c.Data["Title"] = article.Title
	c.Data["Description"] = article.Description

	if json, err := json.Marshal(article.Images); err == nil {
		strrrr := string(json)
		c.Data["JsonImages"] = strrrr
	} else {
		c.Data["JsonImages"] = ""
	}

	c.Layout = "layout/base_view.html"
	c.TplName = "article/view.html"
}

func (c *ArticleController) LoadMore() {
	offset, _ := c.GetInt("offset")
	size := 6
	//tagId, _ := c.GetInt("tagId")
	articles := []models.Article{}
	models.Gorm.Offset(offset).Limit(size).Order("created_at DESC").Preload("Tags").Preload("Vote").Preload("Images").Find(&articles)
	c.JsonRetrun("success", "欢迎访问我们的小站", articles)
}

//评分ajax
func (c *ArticleController) VoteScore() {
	voteID, _ := c.GetInt("voteID")
	score, _ := c.GetFloat("score")
	var vote models.Vote
	models.Gorm.First(&vote, voteID)
	count := float32(vote.VoteCount)
	vote.Score = (vote.Score*count + float32(score)) / (count + 1)
	vote.VoteCount = vote.VoteCount + 1
	models.Gorm.Model(&vote).Update("vote_count", "score")
	c.JsonRetrun("success", "rate score successed", vote)
}
