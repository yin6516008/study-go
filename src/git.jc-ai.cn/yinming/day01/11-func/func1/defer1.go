package func1

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Printf("f1: %d\n", x)
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// DeferTest is a func
func DeferTest() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
