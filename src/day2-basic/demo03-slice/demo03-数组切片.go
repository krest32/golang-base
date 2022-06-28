package main

import "fmt"

func main() {
	var s1[]int
	var s2[]string
	fmt.Println(s1,s2)

	fmt.Println(s1 == nil)


	s1 = []int{1,2,3}
	s2 = []string{"sha","zhan","tyu"}
	fmt.Println(s1,s2)

	// 求取容量和长度
	fmt.Println(len(s1),len(s2))
	fmt.Println(cap(s1),cap(s2))

	// 可变数组
	a1 :=[...]int{1,2,3,5,7,8}
	s3 := a1[0:4]

	fmt.Println(len(a1),cap(a1))

	fmt.Println(s3)
}
