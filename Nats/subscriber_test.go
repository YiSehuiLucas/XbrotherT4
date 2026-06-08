package nats

import (
	"log"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestSubscriber(t *testing.T) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalln(err)
	}
	Subscriber(nc)
}
