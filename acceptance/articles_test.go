package acceptance

import (
	"net/http"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/box"

	"blogotech/api"
	"blogotech/mongo"
	"blogotech/testutils"
)

func TestArticles(t *testing.T) {

	// Fake config:
	//c := config.Config{} // lets move config to its own package

	// Connect to mongo
	m, err := mongo.NewSession("mongodb://localhost:27017/blogotech-acceptance")
	testutils.AssertNil(t, err)
	m.DB("").DropDatabase() // free resources

	// Build API
	b := api.BuildAPI(m) // Lets move api to its own package...

	s := apitest.NewWithHandler(box.Box2Http(b))
	defer s.Destroy() // free resources

	// At this point we need to do a standard http request...

	// now... a simple use case:
	// #1 Create a new article
	res1 := s.Request("POST", "/api/v0/articles").WithBodyString(`
		{
			"title": "my-title",
			"body": "my-body"
		}
	`).Do()

	body1 := *res1.BodyJsonMap()
	testutils.AssertEqual(t, body1["title"], "my-title")
	testutils.AssertEqual(t, body1["body"], "my-body")

	articleID := body1["id"].(string)

	// #2 Get that article by id
	res2 := s.Request("GET", "/api/v0/articles/"+articleID).Do()

	body2 := *res2.BodyJsonMap()
	testutils.AssertEqual(t, body2["id"], articleID)
	testutils.AssertEqual(t, body2["title"], "my-title")
	testutils.AssertEqual(t, body2["body"], "my-body")

	// #3 Delete article
	res3 := s.Request("DELETE", "/api/v0/articles/"+articleID).Do()
	testutils.AssertEqual(t, res3.StatusCode, http.StatusOK)

	// #4 Try to get that article
	res4 := s.Request("GET", "/api/v0/articles/"+articleID).Do()
	testutils.AssertEqual(t, res4.StatusCode, http.StatusNotFound)

}
