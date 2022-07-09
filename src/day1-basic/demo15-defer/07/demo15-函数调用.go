package main

import "fmt"

// 函数内部调用
func max(num1, num2 int) int  {
	var result int

	if(num1>num2){
		result = num1
	}else{
		result = num2
	}
	return result
}

// 可以直接在字符串内调用函数
func  main()  {
	var a,b int = 1,2
	fmt.Printf("%d 是更大的数字\n",max(a,b))
	fmt.Printf("返回的字符串：%s", str())
}

// 函数无参
func str() string  {
	return "aaa"
}