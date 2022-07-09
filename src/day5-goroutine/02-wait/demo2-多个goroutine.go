package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	// wg 相当于一个同步的工具类，等待所有线程执行完毕
	wg.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("所有协程执行完毕")
}