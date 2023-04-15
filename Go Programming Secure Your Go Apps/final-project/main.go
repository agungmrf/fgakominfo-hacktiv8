package main

import (
	"final-project/database"
)

func main() {
	database.StartDB()

	r := routers.StartApp()
	r.Run(":8080")
}
