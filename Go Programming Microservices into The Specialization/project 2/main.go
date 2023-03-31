package main

import (
	"fmt"
	"log"
	"net/http"
	"project-2/model"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"project-2/controller"
	"project-2/repository"
	"project-2/routes"
)

func main() {
	dsn := "host=localhost user=postgres password=Digital2023 dbname=books port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connection successful")

	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migration successful")

	bookRepository := repository.NewBookRepository(db)
	bookController := controller.NewBookController(bookRepository)
	bookRouter := router.NewBookRouter(bookController)

	r := mux.NewRouter()
	bookRouter.SetupRoutes(r.PathPrefix("/books").Subrouter())

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
