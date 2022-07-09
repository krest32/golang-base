package main

import (
	"fmt"
)

func main() {
	var stockade = 123
	var enodata = "2020-12-31"
	// %d 表示整型数字， %s 表示字符串
	var url = "Code=%d, endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockade, enodata)
	fmt.Println(targetUrl)
}