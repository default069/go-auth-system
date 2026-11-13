package cmd

import (
	"auto-system/internal/db"
	"auto-system/internal/handlers"
	"auto-system/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	database := db.InitPostgres()
	if err := database.AutoMigrate(&models.User{});err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	uh := handlers.NewUserHandler(database)

	api := r.Group("/api/v1")
	{
		api.POST("/signup", uh.SingUp)
	}
	log.Fatal(r.Run(":8080"))
}