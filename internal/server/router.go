package server

import (
	"crud-api/internal/notes"
	"net/http"

	"github.com/gin-gonic/gin" // Importing the Gin web framework for handling HTTP requests and routing.
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// when going to path /abc-->
func NewRouter(database *mongo.Database) *gin.Engine {

	router := gin.Default() // Create a new Gin router with default middleware (logger and recovery).

	router.GET("/health", func(c *gin.Context) { // Define a GET route for the /health endpoint.
		c.JSON(http.StatusOK, gin.H{ // Respond with a JSON object containing a "status" key with the value "ok".
			"status": "healthy",
			"ok":     true,
		})
	})
	notes.RegisterRoutes(router, database) // Register the routes for the notes feature by calling the RegisterRoutes function from the notes package, passing in the router and database connection.
	return router // Return the configured router to be used in the main function for starting the server.

}
