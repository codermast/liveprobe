package services

import (
	"LiveProbe/server/models"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var nodeToWebChan = make(chan models.Message)

// WebSocket Upgrader 用来将 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 在生产环境中应根据需求校验请求的来源
		return true
	},
}

// nodeWSHandler 处理函数
func nodeWebSocketHandler(w http.ResponseWriter, r *http.Request) {
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
		err = conn.WriteMessage(messageType, []byte("Success"))
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}

		// 将消息发送给客户端
		nodeToWebChan <- models.Message{
			MessageType: messageType,
			Data:        message,
		}
	}
}

// webWSHandler 处理函数
func webWSHandler(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 升级为 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	// 简单的 WebSocket 消息处理逻辑
	for {

		message := <-nodeToWebChan

		// 将数据发送给客户端
		err = conn.WriteMessage(message.MessageType, message.Data)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
