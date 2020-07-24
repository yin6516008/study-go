package pointer1

import "fmt"


// PrintPointer is a func
func PrintPointer() {
	a := 10
	b := &a
	fmt.Printf("%d, %p\n", a, &a)
	fmt.Printf("pointer: %p \ntype: %T\n", b, b)
	fmt.Println(&b)
}

