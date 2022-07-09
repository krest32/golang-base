package main

import (
	"fmt"
	"time"
)

func main() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new day5-goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Printf("main day5-goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}