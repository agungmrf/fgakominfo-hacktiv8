package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type AppConfig struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func ConnectDB() (*gorm.DB, error) {
	dbConfig := &DBConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "Digital2023",
		Name:     "book",
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username,
		dbConfig.Password, dbConfig.Name)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
