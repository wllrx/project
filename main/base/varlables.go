package main

import "fmt"

/**
变量的使用
*/
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "abc"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("hello")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
}
