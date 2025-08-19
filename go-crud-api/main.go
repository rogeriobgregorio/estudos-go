package main

// This is a simple HTTP server that responds with "Hello, World!" on the root path.
/*
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
*/

// Using Gin framework for a more structured approach
import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var collection *mongo.Collection

func connectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("⛔ MongoDB connection error!")
		log.Fatal(err)
	}
	
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection = client.Database("taskdbs").Collection("tasks")
	log.Println("✅ Connected to MongoDB!")
}

func main() {
	connectDB()
	router := gin.Default()

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}