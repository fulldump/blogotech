package articles

import (
	"blogotech/helpers"
	"blogotech/testutils"
	"context"
	"testing"

	"github.com/fulldump/box"
)

func TestGetArticle(t *testing.T) {

	ctx := context.Background()

	ctx = helpers.SetBoxContext(ctx, &box.C{
		Parameters: map[string]string{
			"articleId": "77",
		},
	})

	result := GetArticle(ctx)

	testutils.AssertEqual(t, result, "this is the article 77")
}
