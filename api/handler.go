package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"seniors-chatbot/model"
	"seniors-chatbot/service"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ChatWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var userMsg model.Message
		json.Unmarshal(msg, &userMsg)

		// 启动流式AI回复
		out := make(chan string)
		errChan := make(chan error)
		go service.GetAIReplyStreamed(userMsg.Content, out, errChan)

		var aiReply string
		streaming := true
		isFirstChunk := true
		for streaming {
			select {
			case chunk, ok := <-out:
				if !ok {
					out = nil
					continue
				}
				aiReply += chunk
				// 适老化处理：只在第一个分片添加前缀
				var contentToSend string
				if isFirstChunk {
					contentToSend = service.SeniorFriendly(chunk)
					isFirstChunk = false
				} else {
					contentToSend = chunk
				}
				// 返回给前端
				resp := model.Message{Sender: "robot", Content: contentToSend}
				respBytes, _ := json.Marshal(resp)
				conn.WriteMessage(websocket.TextMessage, respBytes)
			case err, ok := <-errChan:
				if ok && err != nil {
					fmt.Printf("AI回复失败: %v\n", err)
				}
				errChan = nil
			}
			if out == nil && errChan == nil {
				streaming = false
			}
		}
	}
}
