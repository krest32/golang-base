package main

import (
	"fmt"
	"io"
	"os"
)

// 循环读取文件的内容
func main() {
	file, err := os.Open("E:\\project\\GoPath\\src\\Go-Guide\\src\\day6-lib\\IO\\requestBody.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}