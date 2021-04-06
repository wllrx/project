package main

import "fmt"

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//会int溢出,int类型最大存储一个64位的整数有时会更小
	//fmt.Println(needInt(Big))
}

const (
	//将1左移100位来创建一个非常大的数字
	//即这个数的二进制是1后面加100个0
	Big = 1 << 100
	//再往右移99位,即Small = 1 << 1,或者说Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
