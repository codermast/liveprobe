package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// WebSocket Upgrader 用来将 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 在生产环境中应根据需求校验请求的来源
		return true
	},
}

// WebSocket 处理函数
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 升级为 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	// 简单的 WebSocket 消息处理逻辑
	for {
		// 读取客户端消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received message: %s\n", message)

		// Echo 回消息给客户端
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	// 注册 WebSocket 处理器
	http.HandleFunc("/ws", wsHandler)

	// 启动 WebSocket 服务
	port := ":8080"
	fmt.Println("WebSocket server started at ws://localhost" + port + "/ws")
	log.Fatal(http.ListenAndServe(port, nil))
}
