package array1

import "fmt"

// SumArray is a sum function
func SumArray(x []int) {
	count := 0
	for _, v := range x {
		count += v
	}
	fmt.Println(count)
}
