package redisSelf

import (
	m "T4/Models"
	"testing"
)

func TestSetR(t *testing.T) {
	testData := m.Data{
		Tempreature: "12.12",
		Humility:    "45.23",
	}

	SetR(testData)
}


func TestReadR(t *testing.T) {
	ReadR()
}