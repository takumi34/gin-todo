package main

import (
	"github.com/gin-gonic/gin"
	"github.com/takumi34/gin-todo/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/", controllers.HelloGET)
	router.Run(":8080")
}
