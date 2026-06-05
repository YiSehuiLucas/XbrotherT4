package datagen

import (
	"fmt"
	"testing"
)

func TestDataGen(t *testing.T) {
	T, H := DataGen()
	fmt.Printf("%.2f, %.2f", T, H)
}
