package articles

import (
	"blogotech/mongo"
	"context"
	"fmt"
	"log"
)

// DeleteArticle bla blach
func DeleteArticle(ctx context.Context) (article *Article, err error) {

	article, err = GetArticle(ctx)
	if err != nil {
		return
	}

	err = mongo.GetSession(ctx).DB("").C(collection).RemoveId(article.ID)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("Unexpected Persistence Write Error")
	}

	return
}
