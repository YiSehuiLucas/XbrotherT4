package datagen

import (
	"math/rand"
	"strconv"

	m "T4/Models"
)

// 定义模拟温湿度最大最小值
const min float32 = -40.0
const max float32 = 50.0


// 模拟生成 温湿度数据
func DataGen() m.Data {
	var d m.Data
	t := min + rand.Float32()*(max-min)
	h := min + rand.Float32()*(max-min)
	d.Tempreature = strconv.FormatFloat(float64(t), 'f', 2, 32)
	d.Humility = strconv.FormatFloat(float64(h), 'f', 2, 32)
	return d
}
