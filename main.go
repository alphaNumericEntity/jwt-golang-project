package main

import (
	"fmt"
	"net/http"

	"github.com/alphanumericentity/jwt-example/initializers"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("hello woprld")
}
