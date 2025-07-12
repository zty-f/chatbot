package main

import (
	"log"
	"net/http"

	"seniors-chatbot/api"
)

func main() {
	http.HandleFunc("/ws/chat", api.ChatWebSocket)
	log.Println("服务已启动，监听端口: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
