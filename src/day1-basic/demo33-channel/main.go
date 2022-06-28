package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	// 定义数组
	s := []int{7, 2, 8, -9, 4, 0}
	// 定义整数通道
	c := make(chan int)
	// 开启线程，计算sum(s[:3])，从0-3，包含3以内
	fmt.Println(s[:len(s)/2])
	go sum(s[:len(s)/2], c)
	// 开启线程，计算sum(s[3:])，从3开始，不包含3
	fmt.Println(s[len(s)/2:])
	go sum(s[len(s)/2:], c)

	// 从通道 c 中接收，相当于是栈，先计算的，结果会先出来
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}