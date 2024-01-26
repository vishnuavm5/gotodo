package main

import (
	"hellocheck/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := api.New()
	err = app.Start()

	if err != nil {
		log.Fatal(err)
	}

}
