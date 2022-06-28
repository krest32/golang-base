package main

import (
	"bufio"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)


type Request struct {
	method  	string
	url 		string
	requestBody string
	contentType string
}
func main()  {
	var responseTxt string
	filePath := "src/day6-lib/Ini/02-指针调用-defer & lock-util/requestBody.txt"
	iniFile := "src/day6-lib/Ini/02-指针调用-defer & lock-util/config.ini"
	writeFilePath := "src/day6-lib/Ini/02-指针调用-defer & lock-util/response.txt"
	res := Util(filePath,iniFile)
	if strings.EqualFold(res.method,"post") {
		responseTxt = post(&res)
	}else if strings.EqualFold(res.method,"get") {
		fmt.Printf("get")
	}
	writeFile(responseTxt,writeFilePath)
}

/**
	流程总文件
 */
func Util(filePath,iniFile string) (request Request) {
	cfg, err := ini.Load(iniFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	content := readFile(filePath)
	res := Request{
		method: cfg.Section("request").Key("method").String(),
		url: cfg.Section("request").Key("url").String(),
		requestBody: content,
		contentType: cfg.Section("request").Key("contentType").String(),
	}
	return res
}

/**
	读取文件内容
 */
func readFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return ""
	}
	return string(content)
}

/**
	发送Post请求
 */
func post(res *Request) (response string)  {
	url := res.url
	contentType := res.contentType
	requestBody := res.requestBody

	resp, err := http.Post(url, contentType, strings.NewReader(requestBody))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return ""
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return ""
	}
	return string(b)
}

/**
	将内容写入文件
 */
func writeFile(content,filePath string)  {
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
}