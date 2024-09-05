package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

type MemoryInfo struct {
	Total       uint64  `json:"total"`       // 总内存大小
	Available   uint64  `json:"available"`   // 可用内存大小
	Used        uint64  `json:"used"`        // 已使用内存大小
	UsedPercent float64 `json:"usedPercent"` // 已使用内存百分比
}

// GetNodeMemoryInfo 获取节点内存信息
func GetNodeMemoryInfo() MemoryInfo {
	memInfo := MemoryInfo{
		Total:       getMemoryTotal(),
		Available:   getMemoryAvailable(),
		Used:        getMemoryUsed(),
		UsedPercent: getMemoryUsedPercent(),
	}

	return memInfo
}

// 获取内存信息
func getMemoryInfo() *mem.VirtualMemoryStat {
	// 获取内存虚拟内存统计信息
	virtualMemory, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println("获取虚拟内存信息时出错:", err)
		return nil
	}

	return virtualMemory
}

// 获取已经使用内存百分比
func getMemoryUsedPercent() float64 {
	if getMemoryInfo() == nil {
		return -1
	} else {
		return getMemoryInfo().UsedPercent
	}
}

// 获取内存总大小
func getMemoryTotal() uint64 {
	if getMemoryInfo() == nil {
		return 0
	} else {
		return getMemoryInfo().Total / 1024 / 1024
	}
}

// 获取内存可用大小
func getMemoryAvailable() uint64 {
	if getMemoryInfo() == nil {
		return 0
	} else {
		return getMemoryInfo().Available / 1024 / 1024
	}
}

// 获取内存已用大小
func getMemoryUsed() uint64 {
	if getMemoryInfo() == nil {
		return 0
	} else {
		return getMemoryInfo().Used / 1024 / 1024
	}
}
