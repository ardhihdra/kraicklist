package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"challenge.haraj.com.sa/kraicklist/db"
)

func main() {
	// initialize searcher
	DB := &db.Db{}
	db, err := DB.InitDB()
	if err != nil {
		return
	}

	// define http handlers
	initHandler(db)
	// define port, we need to set it as env for Heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// start server
	fmt.Printf("Server is listening on %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}
