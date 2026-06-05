package datagen

import (
	"math/rand"
)

const min float32 = -40.0
const max float32 = 50.0

// 模拟生成 温湿度数据
func DataGen() (tempreature, humility float32) {
	t := min + rand.Float32()*(max-min)
	h := min + rand.Float32()*(max-min)
	tempreature = t
	humility = h
	return tempreature, humility
}
