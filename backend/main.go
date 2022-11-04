package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	log.Println("Hello World")
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatalf(err)
	}
}
