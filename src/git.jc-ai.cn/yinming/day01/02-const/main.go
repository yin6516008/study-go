package main

// 定义一个常量，常量无法被修改
const pi = 3.1415926

// 声明多个常量
const (
	STATUS_OK        = 200
	STATUS_NOT_FOUND = 404
)

// 在批量声明常量的时候，如果没设置值，那么这个常量和上一个常量同值
const (
	n1 = 100 // 100
	n2       // 100
	n3       // 100
)

// iota是一个常量计数器，默认为0，每新增一行iota + 1
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2        // 1
	_         // 2
	b3        // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3        // 2
	c4        // 3
)

const (
	d1, d2 = iota + 1, iota + 2 // d1:1, d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2, d4:3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 将1左移10位等于 10000000000(二进制)，换算成十进制等于1024
	MB = 1 << (10 * iota) // 将1左移20位
	GB = 1 << (10 * iota) // 将1左移30位
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

}
