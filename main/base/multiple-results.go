package main

import "fmt"

/**
多值返回
函数可以返回任意数量的返回值。

swap 函数返回了两个字符串。
*/
func main() {
	s, s2 := swap("hello", "word")
	fmt.Println(s, s2)
}
func swap(x, y string) (string, string) {
	return x, y
}
