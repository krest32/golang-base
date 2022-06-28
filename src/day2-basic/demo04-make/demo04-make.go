package main

import "fmt"

func main()  {
	// 第一个长度，第二个为容量，可以不写第二个参数，默认为10
	s1 :=make([]int ,5,10)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
}
