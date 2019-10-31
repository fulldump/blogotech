package articles

import (
	"context"
	"testing"

	"github.com/fulldump/box"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"

	"blogotech/helpers"
	"blogotech/mongo"
	"blogotech/testutils"
)

func TestGetArticle(t *testing.T) {

	t.Parallel()

	s, _ := mongo.NewSession("mongodb://localhost:27017/blogotech-test-" + uuid.New().String())
	defer s.DB("").DropDatabase()

	// Create article:
	s.DB("").C(collection).Insert(bson.M{
		"_id":   "my id",
		"title": "my title",
		"body":  "my body",
	})

	// Run test
	ctx := context.Background()
	ctx = mongo.SetSession(ctx, s)
	ctx = helpers.SetBoxContext(ctx, &box.C{
		Parameters: map[string]string{
			"articleId": "my id",
		},
	})

	article, err := GetArticle(ctx)

	testutils.AssertNil(t, err)
	testutils.AssertEqual(t, article.Title, "my title")
	testutils.AssertEqual(t, article.Body, "my body")
}
