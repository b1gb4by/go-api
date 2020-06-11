package main

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/sample_project/sender", controller.Send())

	router.GET("/sample_project/receiver", controller.Receive())

	router.Run()
}
