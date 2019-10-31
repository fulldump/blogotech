package articles

import (
	"blogotech/helpers"
	"blogotech/mongo"
	"context"
	"fmt"
	"log"
)

// GetArticle bla blach
func GetArticle(ctx context.Context) (article *Article, err error) {

	articleID := helpers.GetParams(ctx)["articleId"]

	err = mongo.GetSession(ctx).DB("").C(collection).FindId(articleID).One(&article)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("Unexpected Persistence Read Error")
	}

	return
}
