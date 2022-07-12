package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/**
聊天室
*/
func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
