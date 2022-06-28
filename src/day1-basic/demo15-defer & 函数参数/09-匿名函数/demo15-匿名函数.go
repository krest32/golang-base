package main

import (
	"fmt"
	"math"
)

// 匿名函数
func main() {

	// 局部内部方法
	getSqt := func(a float64) float64 {
		return math.Sqrt(a)
	}

	// 调用局部内部类
	fmt.Println(getSqt(6.30))
}