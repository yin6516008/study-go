package main

import (
	"fmt"
	"func11/func1"
)

func main() {
	fmt.Println("----------------TestFunc------------------")
	func1.TestFunc()

	fmt.Println("----------------FuncAsParameter------------------")
	result := func1.FuncAsParameter(1, 3, func1.Add)
	fmt.Println(result)

	fmt.Println("----------------FuncAsParameter------------------")
	resultInt, resultFunc := func1.FuncAsReturnValue("f")
	fmt.Printf("func: %T, error: %T\n", resultInt, resultFunc)

	fmt.Println("----------------Anonymous------------------")
	func1.AnonymousFunc()

	fmt.Println("----------------Anonymous------------------")
	func1.DeferTest()

	fmt.Println("----------------DispatchCoin---------------")
	left := func1.DispatchCoin()
	fmt.Println("剩下：", left)

}
