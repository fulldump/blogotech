package articles

import (
	"blogotech/mongo"
	"blogotech/testutils"
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestCreateArticle(t *testing.T) {

	t.Parallel()

	s, _ := mongo.NewSession("mongodb://localhost:27017/blogotech-test-" + uuid.New().String())
	defer s.DB("").DropDatabase()

	ctx := context.Background()
	ctx = mongo.SetSession(ctx, s)

	article, err := CreateArticle(ctx, &createArticleRequest{
		Title: "my title",
		Body:  "my body",
	})

	testutils.AssertNil(t, err)
	testutils.AssertEqual(t, article.Title, "my title")
	testutils.AssertEqual(t, article.Body, "my body")

	// TODO: check persistence
}
