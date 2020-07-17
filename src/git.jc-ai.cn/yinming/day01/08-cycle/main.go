package main

import (
	"fmt"
)

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		switch i {
// 		case 5:
// 			goto five
// 		case 6:
// 			goto six
// 		}
// 		fmt.Println(i)
// 		if i == 5 {
// 			break
// 		}
// 	}

// five:
// 	fmt.Println("five")

// 	fmt.Println("seven")

// six:
// 	fmt.Println("six")

// }

func main() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
