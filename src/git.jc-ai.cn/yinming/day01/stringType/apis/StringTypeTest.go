package apis

import "fmt"

// TraversalString is a func
func TraversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// ModifyString modify a String
func ModifyString() {
	s1 := "big"
	byteS1 := []byte(s1)
	fmt.Println(byteS1)
	byteS1[0] = 'p'
	fmt.Println(byteS1)
	fmt.Println(string(byteS1))
}

// PrintVariableType is a func
func PrintVariableType() {
	// int
	vint := 36
	fmt.Printf("%T\n", vint)

	// string
	vstring := "string"
	fmt.Printf("%T\n", vstring)

	// float
	vfloat := 3.14
	fmt.Printf("%T\n", vfloat)

	// bool
	vbool := true
	fmt.Printf("%T\n", vbool)

}

// StatisticalCharacters 统计汉字的数量
func StatisticalCharacters() {
	characters := "hello沙河小王子"

	// result 20
	fmt.Println(len(characters))
	// result uint8(UTF-8)
	fmt.Printf("%T\n", characters[0])

	fmt.Printf("%T\n", rune(characters[0]))
	// result 10
	fmt.Println(len([]rune(characters)))

	// result 2
	fmt.Println(len(string(rune(characters[7]))))

	charactersCount := 0
	for _, v := range characters {
		if len(string(v)) > 1 {
			charactersCount++
		}
	}
	fmt.Println("汉字数量为:", charactersCount)
}
