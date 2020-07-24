package map1

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// SortMap is func
func SortMap() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// MapSlice is a func
func MapSlice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// MapOfValueIsSlice is a func
func MapOfValueIsSlice() {
	var mapSlice = make(map[string][]string, 10)

	key := "China"

	value, ok := mapSlice[key]

	if !ok {
		value = make([]string, 2, 10)
	}

	value = append(value, "BeiJing", "ShangHai")
	mapSlice[key] = value
	fmt.Println(mapSlice)
}

// CountString is a func
func CountString(s string) {
	sSlice := strings.Split(s, " ")

	countMap := make(map[string]int, 10)
	for _, key := range sSlice {
		_, v := countMap[key]
		if !v {
			countMap[key] = 1
		} else {
			countMap[key]++
		}
	}
	fmt.Println(countMap)
}

// Show is a func
func Show() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", m["q1mi"])
}
