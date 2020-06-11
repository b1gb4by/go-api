package controller

import (
	"app/infrastructure"
	_interface "app/interface"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Receive() gin.HandlerFunc {
	return func(context *gin.Context) {

		redis := infrastructure.NewRedis()
		defer redis.CloseRedis()

		key := context.Query("key")

		responseInformation := _interface.UserInformation{}

		if payload, err := redis.Get(key); err != nil {
			fmt.Println("Failed to get data from Redis. :", err)
		} else {
			if err := json.Unmarshal(payload, &responseInformation); err != nil {
				fmt.Println("Could not Unmarshal the retrieved json. :", err)
			}
			context.JSON(http.StatusOK, responseInformation)
		}
	}
}
