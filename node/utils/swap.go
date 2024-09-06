package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

type SwapInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"usedPercent"`
}

// GetNodeSwapInfo 获取swap信息
func GetNodeSwapInfo() SwapInfo {
	swap, err := mem.SwapMemory()

	if err != nil {
		fmt.Println("Error:", err)
		return SwapInfo{}
	}

	return SwapInfo{
		Total:       swap.Total,
		Used:        swap.Used,
		Available:   swap.Free,
		UsedPercent: swap.UsedPercent,
	}
}
