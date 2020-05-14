package main

import "fmt"

func main() {
	// 十进制
	i1 := 10
	fmt.Printf("%d\n", i1) // 十进制，占位符%d
	fmt.Printf("%o\n", i1) // 将十进制转换为八进制
	fmt.Printf("%x\n", i1) // 将十进制转换为十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0x1234e
	fmt.Printf("%d\n", i3)

	fmt.Printf("%T\n", i3) // 查看变量的类型
}
