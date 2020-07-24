package struct13

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

// TestFunc is a func
func TestFunc() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	fmt.Println(m)

}

// Person is a struct
type Person struct {
	name, city string
	age        int8
}

// StructFunc is a func
func StructFunc(name string, age int8, city string) *Person {
	return &Person{
		name: name,
		age:  age,
		city: city,
	}

}

// Dream is a func of the Person
func (p Person) Dream() {
	fmt.Printf("%s is Dreaming.", p.name)
}
