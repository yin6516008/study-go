package main

import "fmt"

func main() {
	n := 100
	fmt.Printf("%T\n", n) // 类型
	fmt.Printf("%v\n", n) // 值
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%d\n", n) // 整型
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%x\n", n) // 十六进制

	s := "star"
	fmt.Printf("%s\n", s) // 字符串
}
