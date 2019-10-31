package main

import (
	"fmt"
	"net/http"

	"github.com/fulldump/box"
	"github.com/fulldump/goconfig"
)

func main() {

	c := struct {
		HttpAddr string
	}{
		HttpAddr: ":8000", // Default http address
	}

	goconfig.Read(&c)

	b := box.NewBox()

	b.Resource("/hello").
		WithActions(
			box.Get(func() string {
				return "Hello world!"
			}),
		)

	s := &http.Server{
		Addr:    c.HttpAddr,
		Handler: box.Box2Http(b),
	}

	fmt.Println("Listening to ", s.Addr)

	s.ListenAndServe()
}
