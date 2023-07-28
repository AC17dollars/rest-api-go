package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	//ping check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//initializing GET routes
	getRoutes()

	router.Run() // listen and serve on 0.0.0.0:8080
}
