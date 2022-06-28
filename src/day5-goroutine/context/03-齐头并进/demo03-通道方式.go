package main

import (
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func worker3(exitChan chan struct{}) {
	LOOP:
		for{
			fmt.Println("worker")
			time.Sleep(time.Second)
			select {
			case <-exitChan:
				break LOOP
			default:
			}
		}
		wg3.Done()
}

func main(){
	var exitChan = make(chan struct{})
	wg3.Add(1)
	go worker3(exitChan)
	time.Sleep(time.Second*3)
	// 给 day5-goroutine 发送推出信号
	exitChan<- struct{}{}
	// 关闭通道
	close(exitChan)
	wg3.Wait()
	fmt.Println("over")
}