package delivery

import (
	"go-crud-api/internal/delivery/dependencies"
	"go-crud-api/internal/interfaces/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) {
	container := dependencies.Setup()

	// Configura CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Injeta e registra rotas
	_ = container.Invoke(func(taskHandler *handlers.TaskHandler) error {
		router.POST("/tasks", taskHandler.CreateTask)
		router.GET("/tasks", taskHandler.GetTasks)
		router.PUT("/tasks/:id", taskHandler.UpdateTask)
		router.DELETE("/tasks/:id", taskHandler.DeleteTask)
		return nil
	})
}
