package acceptance

import (
	"blogotech/testutils"
	"testing"

	"github.com/fulldump/apitest"
)

func TestStatics(t *testing.T) {

	Environment(func(s *apitest.Apitest) {

		resp := s.Request("GET", "/example.html").Do()
		testutils.AssertEqual(t, resp.BodyString(), "<strong>example</strong>")
	})

}
