package main

/**
类型转换
*/
import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//returns the square root of x.返回x的平方根
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(z)
}
