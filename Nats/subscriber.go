package nats

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// 从publisher接收数据
// 数据发送给 redis
func Subscriber() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalln("connect to nats failed", err)
	}
	defer nc.Close()

	fmt.Println("start listening")
	// subscribe message
	_, err = nc.Subscribe("chat.tempreature", func(msg *nats.Msg) {
		fmt.Println("receiced message:", string(msg.Data))
	})

	_, err = nc.Subscribe("chat.humility", func(msg *nats.Msg) {
		fmt.Println("receiced message:", string(msg.Data))
	})

	if err != nil {
		log.Fatalln(err)
	}

	select {}
}
