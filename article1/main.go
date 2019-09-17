package main

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,world")
	})

	// For each matched request Context will hold the route definition
	r.POST("/user/:name/*action", func(c *gin.Context) {

		//c.FullPath() == "/user/:name/*action" // true
	})

	r.Run(":8080")
}
