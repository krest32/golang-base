package main

import "fmt"

// 数组操作
func main() {
	var numbers = make([]int,3,5)
	printSlice(numbers)

}

func printSlice(x []int){
	// len 代表长度
	// cap 代表最大容量
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}