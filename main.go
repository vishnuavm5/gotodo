package main

import (
	"hellocheck/api"
	"log"
)

func main() {

	app := api.New()
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}

}
