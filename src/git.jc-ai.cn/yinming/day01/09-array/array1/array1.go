package array1

import (
	"fmt"
	"sort"
)

// ArrayPrint is a test about the array
func ArrayPrint() {
	// 定义一个长度为3元素类型为int的数组a
	var a [3]int
	fmt.Printf("Type: %T\n", a)

	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]

	fmt.Println("Array slice:", numArray[1:]) // [2,0]
}

// SlicePrint is a func
func SlicePrint() {
	var a = make([]int, 5, 10)
	b := fmt.Sprintf("%v", a[0])
	fmt.Printf("%T\n", b)
	fmt.Println(b)
}

// ModifyArray is a function about modify array
func ModifyArray(x [3][2]int) {
	fmt.Println(x)
	x[2][0] = 100
	fmt.Println(x)
}

// FindIndex is a function
func FindIndex(x []int) {
	var resultList [][]int
	const num = 8

	length := len(x)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if x[i]+x[j] == num {
				resultList = append(resultList, []int{i, j})
			}
		}
	}
	fmt.Println(resultList)
}

// SortArray is a func
func SortArray() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}
