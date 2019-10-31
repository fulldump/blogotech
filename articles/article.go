package articles

import "time"

const collection = "articles"

// Article is a blog post
type Article struct {
	ID              string    `json:"id"                 bson:"_id"`
	CreateTimestamp time.Time `json:"createTimestamp"    bson:"createTimestamp"`
	Title           string    `json:"title"              bson:"title"`
	Body            string    `json:"body"               bson:"body"`
}
