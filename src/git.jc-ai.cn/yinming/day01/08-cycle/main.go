package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		switch i {
		case 5:
			goto five
		case 6:
			goto six
		}
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
five:
	fmt.Println("five")
six:
	fmt.Println("six")

}
