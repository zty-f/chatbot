# seniors-chatbot

面向中老年人社交的AI智能聊天后端服务

## 功能简介
- WebSocket 实时聊天接口
- 对接 OpenAI（或兼容API）实现智能回复
- 适老化内容优化

## 启动方式
1. 配置环境变量 `OPENAI_API_KEY`
2. 进入 `cmd` 目录，运行 `go run main.go`
3. 前端通过 ws://localhost:8080/ws/chat 连接即可体验

## 目录结构
- cmd/main.go         // 程序入口
- api/handler.go      // WebSocket接口
- service/ai.go       // AI对接逻辑
- service/chat.go     // 适老化处理
- model/message.go    // 消息结构体 # chatbot
