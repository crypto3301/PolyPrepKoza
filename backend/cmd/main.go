package main

import (
	"fmt"
	"polyprep/config"
	"polyprep/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)

	fmt.Println("Сервер запущен на порту", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
