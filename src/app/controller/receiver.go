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
		// Redisに接続
		redis := infrastructure.NewRedis()
		// Receive()処理終了後にRedisとの接続を切断する
		defer redis.CloseRedis()

		// クエリストリングパラメーターを取得
		key := context.Query("key")

		// responseInformationをUserInformationで初期化する
		responseInformation := _interface.UserInformation{}

		// Redisからデータを取得する
		if payload, err := redis.Get(key); err != nil {
			fmt.Println("Failed to get data from Redis. :", err)
		} else {
			// Redisから取得したpayloadをGo Object(構造体)に変換する
			if err := json.Unmarshal(payload, &responseInformation); err != nil {
				fmt.Println("Could not Unmarshal the retrieved json. :", err)
			}
			context.JSON(http.StatusOK, responseInformation)
		}
	}
}