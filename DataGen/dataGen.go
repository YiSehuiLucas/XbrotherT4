package datagen

import (
	"math/rand"
	"strconv"
)

// 定义模拟温湿度最大最小值
const min float32 = -40.0
const max float32 = 50.0

// 模拟生成 温湿度数据
func DataGen() (tempreature, humility string) {
	t := min + rand.Float32()*(max-min)
	h := min + rand.Float32()*(max-min)
	tempreature = strconv.FormatFloat(float64(t), 'f', 2, 32)
	humility = strconv.FormatFloat(float64(h), 'f', 2, 32)
	return tempreature, humility
}
