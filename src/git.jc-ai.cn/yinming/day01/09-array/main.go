package main

import (
	"array09/array1"
	"fmt"
)

func main() {
	fmt.Println("-----------ArrayPrint-----------")
	array1.ArrayPrint()

	fmt.Println("-----------ModifyArray------------")
	a := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	array1.ModifyArray(a)

	fmt.Println("-----------SumArray------------")
	b := [...]int{
		1,
		3,
		4,
		7,
		10,
	}
	array1.SumArray(b[:])

	fmt.Println("-----------FindIndex------------")
	x := [5]int{1, 3, 5, 7, 8}
	array1.FindIndex(x[:])

	fmt.Println("-----------SlicePrint------------")
	array1.SlicePrint()

	fmt.Println("-----------SortArray------------")
	array1.SortArray()
}
