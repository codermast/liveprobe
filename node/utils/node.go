package utils

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

type NodeInfo struct {
	NodeId  string     `json:"nodeId"`
	CPU     CPUInfo    `json:"cpu"`
	Memory  MemoryInfo `json:"memory"`
	Disk    DiskInfo   `json:"disk"`
	Network NetInfo    `json:"network"`
}

// GetNodeInfo 获取节点信息
func GetNodeInfo() NodeInfo {
	nodeInfo := NodeInfo{
		NodeId:  GetNodeId(),
		CPU:     GetNodeCPUInfo(),
		Memory:  GetNodeMemoryInfo(),
		Disk:    GetNodeDiskInfoInfo(),
		Network: GetNodeNetInfo(),
	}
	return nodeInfo
}

// SendNodeInfo 发送节点信息
func SendNodeInfo() {
	// TODO WebSocket 服务端地址
	serverURL := "ws://localhost:8080/ws"

	// 解析 URL
	u, err := url.Parse(serverURL)
	if err != nil {
		log.Fatal("URL parse error:", err)
	}

	// 创建 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("WebSocket connection error:", err)
	}
	defer conn.Close()

	// 将系统信息结构体转换为 JSON 格式
	jsonData, err := json.Marshal(GetNodeInfo())
	if err != nil {
		log.Fatal("JSON Marshal error:", err)
	}

	// 通过 WebSocket 发送 JSON 数据
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Fatal("Write message error:", err)
	}

	fmt.Println("Data sent successfully!")
}

// GetNodeId 获取节点 ID
func GetNodeId() string {
	nodeId := os.Getenv("NODE_ID")

	if nodeId == "" {
		// 生成 UUID
		newUUID := uuid.New()
		nodeId = newUUID.String()
		_ = os.Setenv("NODE_ID", nodeId)
	}

	return nodeId
}
