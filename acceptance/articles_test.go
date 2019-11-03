package acceptance

import (
	"net/http"
	"testing"

	"github.com/fulldump/apitest"

	"blogotech/testutils"
)

func TestArticles(t *testing.T) {

	articleID := ""

	Environment(func(s *apitest.Apitest) {
		// #1 Create a new article
		r := s.Request("POST", "/api/v0/articles").WithBodyString(`
			{
				"title": "my-title",
				"body": "my-body"
			}
		`).Do()

		b := *r.BodyJsonMap()
		testutils.AssertEqual(t, b["title"], "my-title")
		testutils.AssertEqual(t, b["body"], "my-body")

		articleID = b["id"].(string)

	}, func(s *apitest.Apitest) {

		// #2 Get that article by id
		r := s.Request("GET", "/api/v0/articles/"+articleID).Do()

		b := *r.BodyJsonMap()
		testutils.AssertEqual(t, b["id"], articleID)
		testutils.AssertEqual(t, b["title"], "my-title")
		testutils.AssertEqual(t, b["body"], "my-body")

	}, func(s *apitest.Apitest) {

		// #3 Delete article
		r := s.Request("DELETE", "/api/v0/articles/"+articleID).Do()
		testutils.AssertEqual(t, r.StatusCode, http.StatusOK)

	}, func(s *apitest.Apitest) {

		// #4 Try to get that article
		r := s.Request("GET", "/api/v0/articles/"+articleID).Do()
		testutils.AssertEqual(t, r.StatusCode, http.StatusNotFound)

	})

}
