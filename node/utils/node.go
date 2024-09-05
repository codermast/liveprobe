package utils

import (
	"github.com/google/uuid"
	"os"
)

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
