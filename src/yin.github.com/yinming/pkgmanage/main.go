package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "hello_world_ff"
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	fmt.Printf("%T\n", s[0])
	fmt.Println(string(unicode.ToLower(rune(s[0]))) + s[1:])
	// s = strings.ToLower() + s[1:]
	fmt.Println(s)
}
