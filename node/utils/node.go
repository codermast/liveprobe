package utils

import (
	"github.com/google/uuid"
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
