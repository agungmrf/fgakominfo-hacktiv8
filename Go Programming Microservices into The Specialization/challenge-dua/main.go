package main

import (
	"challenge-dua/routes"
	"log"
)

func main() {
	r := routes.SetupRouter()

	if err := r.Run(":4000"); err != nil {
		log.Fatal(err)
	}
}
