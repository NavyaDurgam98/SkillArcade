package main

import (
	"BACKEND/Data"
	"BACKEND/controllers"
	"net/http" // handles http requests and responses

	"github.com/gin-gonic/gin" // for using gin framework
)

func main() {

	//Intialize DB connection
	Data.ConnectToDB()

	// initializes a new Gin router for handling incoming API's
	r := gin.Default()

	controllers.UserLoginRouter(r)
	controllers.UserRegisterRouter(r)
	controllers.ForgotRouter(r)
	controllers.ResetRouter(r)

	// Sample routes
	r.GET("/", func(c *gin.Context) { // func(c *gin.Context) request handler function, c is pointer to gin.Context which provides varoius methods to handle http,query,json etc.
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to my API!"})
	})

	// Run server on port 8080
	r.Run()
}
