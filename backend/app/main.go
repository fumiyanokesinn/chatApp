package main

import (
	"log"

	"github.com/fumiyanokesinn/chatApp/api"
	"github.com/fumiyanokesinn/chatApp/config"
)

func main() {
	config.GetEnv()
	router := api.SetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
