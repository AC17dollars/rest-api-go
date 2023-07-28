package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getRoutes() {
	router.GET("/todos", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, database_data)
	})
	router.GET("/todo", func(c *gin.Context) {
		id, _ := c.GetQuery("id")
		intId, _ := strconv.ParseInt(id, 10, 32)
		if id == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Enter a valid 'id' query or use /todos to get all todos",
			})
			return
		}
		for i := 0; i < len(database_data); i++ {

			if database_data[i].Id == int(intId) {
				c.IndentedJSON(http.StatusOK, database_data[i])
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Can't find todo id: " + id + ". The todo doesn't exist.",
		})
	})
}
