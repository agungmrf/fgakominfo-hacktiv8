package main

import (
	"final-project/database"
	"final-project/routers"
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	database.StartDB()
	r := routers.StartApp()
	r.Run(fmt.Sprintf(":%s", port))
}
