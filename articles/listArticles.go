package articles

import (
	"context"
	"fmt"
	"log"

	"blogotech/mongo"
)

// ListArticles bla blach
func ListArticles(ctx context.Context) (output []*Article, err error) {

	output = []*Article{} // Initialize array (avoid empty response)

	err = mongo.GetSession(ctx).DB("").C(collection).Find(nil).All(&output)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("Unexpected Persistence Read Error")
	}

	return
}
