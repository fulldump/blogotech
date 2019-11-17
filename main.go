package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fulldump/box"
	"github.com/fulldump/goconfig"

	"blogotech/api"
	"blogotech/config"
	"blogotech/mongo"
)

func main() {

	// Default config
	c := config.Config{
		HTTPAddr: ":8000",
		Statics:  "statics",
	}

	// Populate configuration
	goconfig.Read(&c)

	// Connect to mongo
	m, err := mongo.NewSession("mongodb://localhost:27017/blogotech")
	if err != nil {
		fmt.Printf("fail connecting to mongo: %s", err.Error())
		os.Exit(1)
	}

	// Build box API
	b := api.BuildAPI(m, c.Statics)

	// Setup server
	s := &http.Server{
		Addr:    c.HTTPAddr,
		Handler: box.Box2Http(b),
	}

	// Run server
	fmt.Println("Listening to ", s.Addr)
	s.ListenAndServe()
}
