package func1

import (
	"errors"
	"fmt"
)

// Jisuan is a func class
type Jisuan func(int, int) int

// Add is a func
func Add(x int, y int) int {
	return x + y
}

// FuncAsParameter is a func
func FuncAsParameter(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// FuncAsReturnValue is a func
func FuncAsReturnValue(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return Add, nil
	case "-":
		return Add, nil
	default:
		return nil, errors.New("无法识别的操作符")
	}
}

// AnonymousFunc is a anonymous function
func AnonymousFunc() {
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

}

// TestFunc is a func
func TestFunc() {
	var f Jisuan
	f = Add
	fmt.Printf("%T\n", f)
}
