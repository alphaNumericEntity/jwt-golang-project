package main

import (
	"github.com/alphanumericentity/jwt-example/controllers"
	"github.com/alphanumericentity/jwt-example/initializers"
	"github.com/alphanumericentity/jwt-example/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	//initialize()
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.Authenticate, controllers.Validate)
	r.Run()

}
