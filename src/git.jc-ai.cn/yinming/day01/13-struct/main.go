package main

import (
	"fmt"
	"struct13/struct13"
)

func main() {
	struct13.TestFunc()

	person := struct13.StructFunc("yin", 18, "wuhan")
	fmt.Printf("%#v\n", person)
	person.Dream()
}
