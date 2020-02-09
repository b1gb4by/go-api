package infrastructure

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

type Redis struct {
	connection redis.Conn
}

// 接続処理
func NewRedis() *Redis {
	// IPポートの設定
	const ipPort = "redis:6379"

	// redisに接続する
	c, err := redis.Dial("tcp", ipPort)
	if err != nil {
		panic(err)
	}

	// 接続情報をConnインタフェースのconnectionに保存
	r := &Redis{
		connection: c,
	}

	return r
}

// 切断処理
func (r *Redis) CloseRedis() {
	// redisとの通信を切断する
	_ = r.connection.Close()
}

func (r *Redis) Set(key string, payload []byte) error {
	// 生成したキーが既に存在するかチェックする
	if r.keyExist(key) == true {
		fmt.Println("Delete the key because it was already registered in redis.")
		fmt.Println("Update an existing key.")
		// 存在する場合、データを更新する
		r.update(key, payload)
	} else {
		// キーをredisに追加する
		if _, err := r.connection.Do("SET", key, payload); err != nil {
			fmt.Println("infrastructure/database/Set() : ", err)
			os.Exit(1)
			return err
		}
	}
	return nil
}

func (r *Redis) Get(key string) ([]byte, error) {
	// キーを使用し、redisからデータを追加する
	payload, err := redis.Bytes(r.connection.Do("GET", key))
	if err != nil {
		fmt.Println("infrastructure/database/Set() : ", err)
		return payload, err
	}
	return payload, err
}

func (r *Redis) keyExist(key string) bool {
	// キーが既にredis内に存在するかチェックする
	result, err := redis.Bool(r.connection.Do("EXISTS", key))
	if err != nil {
		fmt.Println("infrastructure/database/keyExist() : ", err)
	}
	return result
}

func (r *Redis) update(key string, payload []byte) {
	// キーから値を取得後、新たなデータを登録する
	_, err := r.connection.Do("GETSET", key, payload)
	if err != nil {
		fmt.Println("infrastructure/database/update() : ", err)
	}
}