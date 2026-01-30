package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"zelential.com/user-service/handler"
	"zelential.com/user-service/repository"
	"zelential.com/user-service/service"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=zelential sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	jwtKey := []byte("secret_key")
	userHandler := handler.NewUserHandler(userService, jwtKey)

	r := gin.Default()
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, "ok") })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
