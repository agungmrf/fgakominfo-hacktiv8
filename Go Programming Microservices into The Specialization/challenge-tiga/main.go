package main

import (
	"log"

	"challenge-tiga/config"
	"challenge-tiga/models"
	"challenge-tiga/repository"
	"challenge-tiga/router"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate model Book
	if err := db.AutoMigrate(&models.Book{}).Error; err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	repo := repository.NewBookRepository(db)

	r := router.SetupRouter(repo)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
