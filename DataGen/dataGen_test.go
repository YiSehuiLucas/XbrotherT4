package datagen

import (
	"fmt"
	"testing"
)

func TestDataGen(t *testing.T) {
	T, H := DataGen()
	fmt.Println(T)
	fmt.Println(H)
}
