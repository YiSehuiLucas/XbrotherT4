package nats

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func Publisher(T, H string) {
	// connect to nats
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	// data := []byte(T)
	// data = append(data, byte(H))

	err = nc.Publish("chat.tempreature", []byte(T))
	if err != nil {
		log.Fatalln("tempreature send failed")
	}
	err = nc.Publish("chat.humility", []byte(H))
	if err != nil {
		log.Fatalln("humility send failed")
	}
	fmt.Println("message send successfully")
}
