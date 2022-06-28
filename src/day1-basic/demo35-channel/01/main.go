package main

import "fmt"

func main()  {
	x := 1
	u := 2
	// x = u, u = x+u, 同时进行运算
	x , u = u, x+u
	fmt.Println(x,u)

}
