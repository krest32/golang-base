package main

import (
	"fmt"
	"time"
)

// 定义一个方法
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 开启一个携程执行任务
	go say("world")
	say("hello")
}