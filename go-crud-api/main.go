package main

// @title   		Task Management API
// @version 		1.0
// @description API for managing tasks using Go, Gin, and MongoDB.
// @host 				localhost:8080
// @BasePath 		/

import (
	"context"
	"go-crud-api/models"
	"log"
	"net/http"
	"time"

	_ "go-crud-api/docs" // This line is necessary for go-swagger to find the docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func connectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛔ MongoDB connection error:", err)
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection = client.Database("taskdbs").Collection("tasks")
	log.Println("✅ Connected to MongoDB!")
}

// @Summary 		Create a new task
// @Description Create a new task
// @Tags 				tasks
// @Accept 			json
// @Produce 		json
// @Param 			task 		body 			models.Task 						true "Task data"
// @Sucess 			201 		{object} 	map[string]interface{}
// @Failure 		400 		{object} 	map[string]string
// @Failure 		500 		object} 	map[string]string
// @Router 			/tasks 	[post]
func createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// @Summary 		List all tasks
// @Description List all tasks
// @Tags 				tasks
// @Produce 		json
// @Sucess 			200 		{array} 	models.Task
// @Failure 		500 		{object} 	map[string]string
// @Router 			/tasks 	[get]
func listTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// @Summary 		Update a task
// @Description Update a task by ID
// @Tags 				tasks
// @Accept 			json
// @Produce 		json
// @Param 			id 		path 			string 						true "Task ID"
// @Param 			task 	body 			models.Task 			true "Updated task data"
// @Sucess 			200 	{object} 	map[string]string
// @Failure 		400 	{object} 	map[string]string
// @Failure 		404 	{object} 	map[string]string
// @Failure 		500 	{object} 	map[string]string
// @Router 			/tasks/{id} 		[put]
func updateTask(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": task})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// @Summary 		Delete a task
// @Description Delete a task by ID
// @Tags 				tasks
// @Produce 		json
// @Param 			id 		path 			string 						true "Task ID"
// @Sucess 			200 	{object} 	map[string]string
// @Failure 		400 	{object} 	map[string]string
// @Failure 		404 	{object} 	map[string]string
// @Failure 		500 	{object} 	map[string]string
// @Router 			/tasks/{id} 		[delete]
func deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func main() {
	connectDB()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// Define a route for listing tasks
	router.GET("/tasks", listTasks)

	// Define a route for creating tasks
	router.POST("/tasks", createTask)

	// Define a route for updating tasks
	router.PUT("/tasks/:id", updateTask)

	// Define a route for deleting tasks
	router.DELETE("/tasks/:id", deleteTask)

	// Start the server on port 8080 and handle errors
	if err := router.Run(":8080"); err != nil {
		log.Fatal("⛔ Error starting server:", err)
	}
}
