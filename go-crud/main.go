package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostCreate)
	router.GET("/posts", controllers.PostIndex)
	router.GET("/posts/:id", controllers.PostShow)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.DELETE("posts/:id", controllers.PostDelete)
	router.Run() // listen and serve on 0.0.0.0:8080
}
