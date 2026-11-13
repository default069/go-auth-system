package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func InitPostgres() *gorm.DB{
	dsn := "postgres://postgres:1111@localhost:5432/authUsers?sslmode=disable"

	database , err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil{
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("connected")

	return database
}