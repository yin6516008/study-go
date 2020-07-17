package main

import (
	"fmt"
	"math"
)

func main() {
	f1 := 1.343
	fmt.Printf("%T\n", f1) // 默认go语言中，小数都是float64类型的
	f2 := float32(1.23455) // 如果需要定义float32，必须显示声明的定义
	fmt.Printf("%T\n", f2)

	fmt.Printf("%f\n", math.MaxFloat32)
}
