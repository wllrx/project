package main

import (
	"fmt"
	"math"
)

/**
if的简单语句
同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
该语句声明的变量作用域仅在 if 之内
*/
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	//returns x**y, the base-x exponential of y.  返回x**y x是基数,y是指数
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
