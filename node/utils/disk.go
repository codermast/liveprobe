package utils

import (
	"github.com/shirou/gopsutil/disk"
)

// 获取磁盘信息
func getDiskInfo() *disk.UsageStat {
	usage, err := disk.Usage("/")

	if err != nil {
		return nil
	}
	return usage
}

// GetDiskPercent 获取磁盘使用率
func GetDiskPercent() float64 {
	return getDiskInfo().UsedPercent
}

// GetDiskTotalSpace 获取磁盘总空间
func GetDiskTotalSpace() uint64 {
	return getDiskInfo().Total / (1024 * 1024 * 1024)
}

// GetDiskFreeSpace 获取磁盘可用空间
func GetDiskFreeSpace() uint64 {
	return getDiskInfo().Free / 1024 / 1024
}

// GetDiskUsedSpace 获取磁盘使用空间
func GetDiskUsedSpace() uint64 {
	return getDiskInfo().Used / (1024 * 1024 * 1024)
}

// GetDiskMountPath 获取磁盘挂载点
func GetDiskMountPath() string {
	return getDiskInfo().Path
}
