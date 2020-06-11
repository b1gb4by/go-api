package controller

import (
	"app/infrastructure"
	"app/interface"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send() gin.HandlerFunc {
	return func(context *gin.Context) {
		redis := infrastructure.NewRedis()
		defer redis.CloseRedis()

		requestInformation := _interface.UserInformation{}

		err := context.Bind(&requestInformation)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Status": "BadRequest"})
		}

		key := requestInformation.ID + ":" + requestInformation.Name

		payload, err := json.Marshal(requestInformation)
		if err != nil {
			fmt.Println("JSON Marshal Error : ", err)
			return
		}

		if err := redis.Set(key, payload); err != nil {
			fmt.Println("Failed to store data in Redis. ", err)
		} else {
			context.JSON(http.StatusOK, gin.H{"Status": "Successfully added to redis. "})
		}
	}
}
