package main

import (
	"fmt"
)

// 通过数组的形式修改字符串
func main() {
	// 定义字符串
	s1 := "big"

	// 定义数组，将字符串转化为数组
	byteArr :=[]byte(s1)

	// 修改数组中的字符
	byteArr[0] = 'p'
	byteArr[2] = 'p'

	fmt.Println(string(byteArr))
}
