package main

import (
	"LiveProbe/node/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {

	// 1. 获取服务器端地址
	serverUrl := services.GetNodeServerUrl()

	if serverUrl == "" {
		fmt.Print("首次运行，请输入您的 WebSocket 服务端地址: ")
		// 从标准输入读取单个输入
		fmt.Scanln(&serverUrl)
		services.SetNodeServerUrl(serverUrl)
	}

	// 解析 URL
	u, err := url.Parse(serverUrl)
	if err != nil {
		log.Fatal("WebSocket 地址异常:", err)
	}

	// 2. 连接 WebSocket 服务端
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("WebSocket 连接异常:", err)
	} else {
		log.Println("WebSocket 连接成功！")
	}

	defer conn.Close()

	for {
		nodeInfo := services.GetNodeInfo()

		// 将系统信息结构体转换为 JSON 格式
		jsonData, err := json.Marshal(nodeInfo)
		if err != nil {
			log.Fatal("JSON Marshal error:", err)
		}

		// 通过 WebSocket 发送 JSON 数据
		err = conn.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			log.Fatal("Write message error:", err)
		} else {
			log.Println("数据发送成功！")
		}
	}
}
