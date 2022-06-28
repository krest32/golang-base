package main

import (
	"fmt"
	"html/template"
	"net/http"
)


type UserInfo struct {
	Name string
	Gender string
	Age int
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	// 解析两个tmpl文件
	tmpl, err := template.ParseFiles("src/day6-lib/template/demo03-slice-变量类型/t.tmpl",
		"src/day6-lib/template/demo03-slice-变量类型/ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", tmplDemo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
