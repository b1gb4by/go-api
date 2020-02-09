package main

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Sender
	router.POST("/sample_project/sender", controller.Send())

	// Receiver
	router.GET("/sample_project/receiver", controller.Receive())

	// PORT環境変数が未指定の場合、デフォルトで8080で待受する
	router.Run()
}
