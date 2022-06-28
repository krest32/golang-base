package main

import (
	"fmt"
	"sync"
	"time"
)

var wb sync.WaitGroup
var exit bool

func worker2()  {
	for{
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wb.Done()
}

func main(){
	wb.Add(1)
	go worker2()
	// 避免线程执行额过快
	time.Sleep(time.Second*3)
	exit = true
	wb.Wait()
	fmt.Println("over")
}

