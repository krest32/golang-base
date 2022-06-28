package main

import "fmt"

func main()  {
	var(
		name string
		age int
		class string
	)
	//获取多个参数
	fmt.Scanln(&name,&age,&class)
	fmt.Println(name,age,class)
}
