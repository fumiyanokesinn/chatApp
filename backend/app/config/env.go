package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file:\n%s", err))
	}
}

func TestGetEnv() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file:\n%s", err))
	}
}
