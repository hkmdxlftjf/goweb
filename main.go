package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/helloworld", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.JSON(http.StatusOK, gin.H{
			"message": "hello " + hostname,
		})
	})
	r.POST("/post", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.JSON(http.StatusOK, gin.H{
			"message": "hello post " + hostname,
		})
	})
	r.GET("/header", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.Request.Header)
	})
	r.GET("/500", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, c.Request.Header)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
