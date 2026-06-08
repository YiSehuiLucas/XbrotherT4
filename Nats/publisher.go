package nats

import (
	"encoding/json"
	"log"

	m "T4/Models"

	"github.com/nats-io/nats.go"
)


func Connect() *nats.Conn {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalln(err)
	}
	return nc
}

func Publisher(d m.Data,nc *nats.Conn) {
	// Encode
	msg, err := json.Marshal(d)
	if err != nil {
		log.Fatalln(err)
	}

	err = nc.Publish("senior.data", msg)
	if err != nil {
		log.Fatalln("data send failed")
	}

	// fmt.Println("send:", d)
}
