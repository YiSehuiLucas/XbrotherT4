package redisSelf

import (
	"context"
	"encoding/json"
	"log"
	"time"

	m "T4/Models"

	"github.com/redis/go-redis/v9"
)

func SetR(d m.Data) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	data, err := json.Marshal(d)
	if err != nil {
		log.Fatalln(err)
	}
	// 设置超时时间为 1s
	err = rdb.Set(context.Background(), "data", data, 10*time.Second).Err()
	if err != nil {
		log.Println("redis write error:", err)
	}

}

func ReadR() m.Data {
	var data m.Data

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	val, err := rdb.Get(context.Background(), "data").Result()
	if err != nil {
		log.Fatal("redis read error:", err)
	}

	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("data", data)
	// fmt.Printf("%T", data)

	return data
}
