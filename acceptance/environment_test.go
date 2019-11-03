package acceptance

import (
	"blogotech/api"
	"blogotech/mongo"

	"github.com/fulldump/apitest"
	"github.com/fulldump/box"
	"github.com/google/uuid"
)

func Environment(callbacks ...func(a *apitest.Apitest)) {

	// Connect to mongo
	mongoURI := "mongodb://localhost:27017/blogotech-acceptance-" + uuid.New().String()
	m, _ := mongo.NewSession(mongoURI)
	m.DB("").DropDatabase() // free resources

	// Build API
	b := api.BuildAPI(m) // Lets move api to its own package...

	a := apitest.NewWithHandler(box.Box2Http(b))
	defer a.Destroy() // free resources

	// Execute all callbacks
	for _, f := range callbacks {
		f(a)
	}

}
