package nats

import (
	m "T4/Models"
	"log"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestPublisher(t *testing.T) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalln(err)
	}
	testData := m.Data{
		Tempreature: "12.12",
		Humility:    "45.23",
	}
	Publisher(testData, nc)
}
