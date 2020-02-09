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
		// Redisに接続する
		redis := infrastructure.NewRedis()
		// Send()処理終了後にRedisとの接続を切断する
		defer redis.CloseRedis()

		// requestInformationをUserInformationで初期化する
		requestInformation := _interface.UserInformation{}

		// 構造体をBINDする
		err := context.Bind(&requestInformation)
		if err != nil{
			context.JSON(http.StatusBadRequest, gin.H{"Status": "BadRequest"})
		}

		// Redisで使用するキーの作成
		key := requestInformation.ID + ":" + requestInformation.Name

		// 作成した構造体requestInformationをJSONに変換する
		payload, err := json.Marshal(requestInformation)
		if err != nil {
			fmt.Println("JSON Marshal Error : ", err)
			return
		}

		// key, payloadを引数にredisに追加する
		if err := redis.Set(key, payload); err != nil {
			fmt.Println("Failed to store data in Redis. ", err)
		} else {
			context.JSON(http.StatusOK, gin.H{"Status": "Successfully added to redis. "})
		}
	}
}
