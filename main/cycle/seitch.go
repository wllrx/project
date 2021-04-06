package main

import (
	"fmt"
	"runtime"
)

/**
switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。
实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。
Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
*/
func main() {

	fmt.Println("go runs on")
	//正在运行的操作系统
	switch os := runtime.GOOS; os {
	//在windows系统的执行时,断点打的此处会直接执行完毕,跳过断点处
	case "darwin":
		fmt.Println("OX X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
