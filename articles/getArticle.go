package articles

import (
	"blogotech/helpers"
	"context"
)

// GetArticle bla blach
func GetArticle(ctx context.Context) string {

	articleID := helpers.GetParams(ctx)["articleId"]

	return "this is the article " + articleID
}
