package server

import (
	"net/http"

	"github.com/gin-gonic/gin" // Importing the Gin web framework for handling HTTP requests and routing.
)

// when going to path /abc-->
func NewRouter() *gin.Engine {

	router := gin.Default() // Create a new Gin router with default middleware (logger and recovery).

	router.GET("/health", func(c *gin.Context) { // Define a GET route for the /health endpoint.
		c.JSON(http.StatusOK, gin.H{ // Respond with a JSON object containing a "status" key with the value "ok".
			"status": "healthy",
			"ok":     true,
		})
	})
	return router // Return the configured router to be used in the main function for starting the server.

}
