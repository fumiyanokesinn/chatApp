package main

import (
	"log"

	app "github.com/fumiyanokesinn/chatApp/api"
)

func main() {
	router := app.SetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
