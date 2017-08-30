package controllers

import (
	"encoding/json"
	"fmt"
	"my_go_web/models"
	"time"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) View() {
	articleID, _ := c.GetInt(":id")
	//浏览计数
	vote := models.Vote{}
	models.Gorm.Where("article_id = ?", articleID).Find(&vote)
	vote.Visit++
	models.Gorm.Model(&vote).Update("visit")

	var article models.Article
	cacheKey := fmt.Sprint("mojotv.article_detail.", articleID)
	//fmt.Println(cacheKey)
	if x, found := models.CacheManager.Get(cacheKey); found {
		//x 就是 []byte
		buffer := x.([]byte)
		json.Unmarshal(buffer, &article)
	} else {
		if models.Gorm.Preload("Tags").Preload("Images").Preload("Shows").First(&article, articleID).RecordNotFound() {
			c.Abort("404")
		}
		buffer, _ := json.Marshal(article)
		var expireTime time.Duration
		if article.UpdatedAt.After(time.Now().Add(-time.Minute * 30)) {
			expireTime = time.Minute * 2
		} else if article.UpdatedAt.After(time.Now().Add(-time.Hour * 2)) {
			expireTime = time.Minute * 15
		} else if article.UpdatedAt.After(time.Now().Add(-time.Hour * 6)) {
			expireTime = time.Minute * 30
		} else if article.UpdatedAt.After(time.Now().Add(-time.Hour * 12)) {
			expireTime = time.Hour * 1
		} else if article.UpdatedAt.After(time.Now().Add(-time.Hour * 24)) {
			expireTime = time.Hour * 2
		} else if article.UpdatedAt.After(time.Now().Add(-time.Hour * 48)) {
			expireTime = time.Hour * 6
		} else if article.UpdatedAt.After(time.Now().Add(-time.Hour * 24 * 3)) {
			expireTime = time.Hour * 24
		} else if article.UpdatedAt.After(time.Now().Add(-time.Hour * 24 * 7)) {
			expireTime = time.Hour * 48
		} else {
			expireTime = models.C_EXPIRE_TIME_FOREVER
		}
		models.CacheManager.Set(cacheKey, buffer, expireTime)
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
	//tagId, _ := c.GetInt("tagId")
	articles := models.GetBatchArticles(offset, 6)
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
