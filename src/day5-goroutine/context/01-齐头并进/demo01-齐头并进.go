package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker() {
	// 执行打印任务
	for i := 0; i <= 10; i++{
		fmt.Println("worker:", i)
		// 等待1秒钟
		time.Sleep(time.Second)
	}
	// 接受外部命令进行退出
	wg.Done()
}

func main(){
	wg.Add(1)
	// 开启线程执行
	go worker()
	// 等待线程执行完毕
	wg.Wait()
	fmt.Println("over")
}
