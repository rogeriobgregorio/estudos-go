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
	"github.com/gin-gonic/gin"
)

func main() {
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