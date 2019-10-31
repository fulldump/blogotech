package main

import (
	"github.com/fulldump/box"
)

func main() {

	b := box.NewBox()

	b.Resource("/hello").
		WithActions(
			box.Get(func() string {
				return "Hello world!"
			}),
		)

	b.Serve()
}
