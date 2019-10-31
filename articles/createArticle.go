package articles

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"blogotech/mongo"
)

type createArticleRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// CreateArticle bla blach
func CreateArticle(ctx context.Context, input *createArticleRequest) (article *Article, err error) {

	article = &Article{
		ID:              uuid.New().String(), // TODO: generate uuid here
		CreateTimestamp: time.Now(),
		Title:           input.Title,
		Body:            input.Body,
	}

	err = mongo.GetSession(ctx).DB("").C(collection).Insert(article)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("Unexpected Persistence Write Error")
	}

	return
}
