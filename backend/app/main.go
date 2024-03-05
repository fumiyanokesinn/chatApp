package main

import (
	"log"

	"github.com/fumiyanokesinn/chatApp/api"
)

func main() {
	router := api.SetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
