package main

import (
	"LiveProbe/node/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {

	// TODO WebSocket 服务端地址
	var wsUrl string
	fmt.Print("请输入你的WebSocket 服务端地址: ")
	// 从标准输入读取单个输入
	fmt.Scanln(&wsUrl)

	serverURL := wsUrl
	//serverURL := "ws://localhost:8080/ws"

	// 解析 URL
	u, err := url.Parse(serverURL)
	if err != nil {
		log.Fatal("WebSocket 地址异常:", err)
	}

	// 创建 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("WebSocket 连接异常:", err)
	} else {
		log.Println("WebSocket 连接成功！")
	}

	defer conn.Close()

	for {
		nodeInfo := utils.GetNodeInfo()

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
