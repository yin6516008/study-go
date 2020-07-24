package main

import (
	"fmt"
	"map10/map1"
)

func main() {
	fmt.Println("--------------SortMap---------------")
	map1.SortMap()

	fmt.Println("--------------MapSlice---------------")
	map1.MapSlice()

	fmt.Println("--------------MapOfValueIsSlice---------------")
	map1.MapOfValueIsSlice()

	fmt.Println("--------------CountString---------------")
	s := "hello my name is hello"
	map1.CountString(s)

	fmt.Println("--------------Show---------------")
	map1.Show()
}
