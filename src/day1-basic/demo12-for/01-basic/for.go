package main

import "fmt"

// 单层for
func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	// 传统for循环
	for a := 0; a < 10; a++ {
		fmt.Printf("传统for循环 a 的值为: %d\n", a)
	}

	// 类似布尔的for循环
	for a < b {
		a++
		fmt.Printf("类似布尔的for循环 a 的值为: %d\n", a)
	}

	// 使用range执行for循环
	for i, x:= range numbers {
		fmt.Printf("使用range执行for循环 第 %d 位 x 的值 = %d\n", i, x)
	}
}