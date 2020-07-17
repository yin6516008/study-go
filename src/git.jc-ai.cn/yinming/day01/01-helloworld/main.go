package main

import "fmt"

// 声明全局变量
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "yin"
	age = 18
	isOk = true

	// 类型推导
	gener := "man"

	// 匿名变量
	_ = name

	fmt.Printf("name: %s\n", name)
	fmt.Println(age)
	fmt.Println(isOk)
	fmt.Println(gener)
	
}
