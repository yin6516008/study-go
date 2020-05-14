package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串用"(双引号)包裹
	s := "hello"

	// 字符用'(单引号)包裹，字符是一个单独的字符、汉字、符号
	c1 := 'h'
	c2 := '1'
	c3 := '王'
	fmt.Println(c1, c2, c3, s)

	// 转义
	path := "D:\\Go\\src\\github.com\\day01"
	fmt.Println(path)

	// 多行字符串
	s1 := `
		I'm Ok,
		I'm a mutil line string
	`
	fmt.Println(s1)

	// 字符串长度
	fmt.Println(len(s1))

	// 字符串拼接
	one := "hello"
	two := "world"

	three := one + two
	four := fmt.Sprintf("%s not %s", one, two)

	// 字符串分割
	ret := strings.Split(three, "not")

	// 包含
	fmt.Println(strings.Contains(four, "hello"))

	// 前缀
	fmt.Println(strings.HasPrefix(four, "hello"))

	// 后缀
	fmt.Println(strings.HasPrefix(four, "world"))

	// 索引
	fmt.Println(strings.Index(four, "o"))
	fmt.Println(strings.LastIndex(four, "o"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))

	// 字符串修改
	name := "Tom"
	nameRune := []rune(name)
	nameRune[0] = 'J'
	fmt.Println(string(nameRune))

	// 字符串和字符
	h1 := "红" // string
	h2 := '红' // rune(int32)
	fmt.Printf("h1: %T, h2: %T", h1, h2)
}
