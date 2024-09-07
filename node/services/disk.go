package services

import (
	"github.com/shirou/gopsutil/disk"
)

type DiskInfo struct {
	Total       uint64  `json:"total"`       // 总内存大小
	Available   uint64  `json:"available"`   // 可用内存大小
	Used        uint64  `json:"used"`        // 已使用内存大小
	UsedPercent float64 `json:"usedPercent"` // 已使用内存百分比
	MountPath   string  `json:"mountPath"`   // 磁盘挂载点
}

// GetNodeDiskInfoInfo 获取节点磁盘信息
func GetNodeDiskInfoInfo() DiskInfo {
	diskInfo := DiskInfo{
		Total:       getDiskTotalSpace(),
		Available:   getDiskAvailableSpace(),
		Used:        getDiskUsedSpace(),
		UsedPercent: getDiskUsedPercent(),
		MountPath:   getDiskMountPath(),
	}

	return diskInfo
}

// 获取磁盘信息
func getDiskInfo() *disk.UsageStat {
	usage, err := disk.Usage("/")

	if err != nil {
		return nil
	}
	return usage
}

// getDiskPercent 获取磁盘使用率
func getDiskUsedPercent() float64 {
	return getDiskInfo().UsedPercent
}

// getDiskTotalSpace 获取磁盘总空间
func getDiskTotalSpace() uint64 {
	return getDiskInfo().Total / (1024 * 1024 * 1024)
}

// getDiskAvailableSpace 获取磁盘可用空间
func getDiskAvailableSpace() uint64 {
	return getDiskInfo().Free / 1024 / 1024
}

// getDiskUsedSpace 获取磁盘使用空间
func getDiskUsedSpace() uint64 {
	return getDiskInfo().Used / (1024 * 1024 * 1024)
}

// getDiskMountPath 获取磁盘挂载点
func getDiskMountPath() string {
	return getDiskInfo().Path
}
