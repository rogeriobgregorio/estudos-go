package main

import (
	_ "go-crud-api/docs"
	"go-crud-api/internal/delivery"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title   		Task Management API
// @version 		1.0
// @description API for managing tasks using Go, Gin, and MongoDB.
// @host 				localhost:8080
// @BasePath 		/
func main() {
	r := gin.Default()

	// Rota do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Passa o router para a função que registra as rotas da API
	delivery.Start(r)

	// Inicia o servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
