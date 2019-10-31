package main

import (
	"fmt"
	"net/http"

	"github.com/fulldump/box"
	"github.com/fulldump/goconfig"
)

func main() {

	// Default config
	c := Config{
		HTTPAddr: ":8000",
	}

	// Populate configuration
	goconfig.Read(&c)

	// Build box API
	b := BuildAPI()

	// Setup server
	s := &http.Server{
		Addr:    c.HTTPAddr,
		Handler: box.Box2Http(b),
	}

	// Run server
	fmt.Println("Listening to ", s.Addr)
	s.ListenAndServe()
}
