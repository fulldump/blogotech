package main

import (
	"github.com/fulldump/box"
	"github.com/globalsign/mgo"

	"blogotech/articles"
	"blogotech/mongo"
)

// BuildAPI return box Handler
func BuildAPI(s *mgo.Session) *box.B {
	b := box.NewBox()

	v0 := b.Resource("/api/v0")

	v0.WithInterceptors(
		mongo.SessionInterceptor(s),
	)

	v0.Resource("/articles").
		WithActions(
			box.Get(articles.ListArticles),
			box.Post(articles.CreateArticle),
		)

	v0.Resource("/articles/{articleId}").
		WithActions(
			box.Get(articles.GetArticle),
			box.Delete(articles.DeleteArticle),
			box.Patch(articles.UpdateArticle),
		)

	return b
}
