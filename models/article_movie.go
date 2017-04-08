package models

type ArticleMovie struct {
	ArticleId *Articles `orm:"column(article_id);rel(fk)"`
	MovieId   *Movies   `orm:"column(movie_id);rel(fk)"`
}
