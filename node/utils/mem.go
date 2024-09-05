package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func getMemoryInfo() *mem.VirtualMemoryStat {
	// 获取内存虚拟内存统计信息
	virtualMemory, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println("获取虚拟内存信息时出错:", err)
		return nil
	}

	return virtualMemory
}

func GetMemoryUsedPercent() float64 {
	if getMemoryInfo() == nil {
		return -1
	} else {
		return getMemoryInfo().UsedPercent
	}
}

func GetMemoryTotal() uint64 {
	if getMemoryInfo() == nil {
		return 0
	} else {
		return getMemoryInfo().Total / 1024 / 1024
	}
}

func GetMemoryAvailable() uint64 {
	if getMemoryInfo() == nil {
		return 0
	} else {
		return getMemoryInfo().Available / 1024 / 1024
	}
}

func GetMemoryUsed() uint64 {
	if getMemoryInfo() == nil {
		return 0
	} else {
		return getMemoryInfo().Used / 1024 / 1024
	}
}
