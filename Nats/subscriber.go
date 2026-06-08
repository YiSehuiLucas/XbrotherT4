package nats

import (
	m "T4/Models"
	redisSelf "T4/Redis"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// 从publisher接收数据
// 数据发送给 redis
func Subscriber(nc *nats.Conn) {
	// Decode
	// subscribe message
	var data m.Data
	_, err := nc.Subscribe("senior.data", func(msg *nats.Msg) {
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("received : ", data)
		// 写入 redis
		redisSelf.SetR(data)
		fmt.Println("write into redis")

	})
	if err != nil {
		log.Fatalln(err)
	}

	// select {}
}
