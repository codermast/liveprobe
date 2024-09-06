package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"time"
)

type HostInfo struct {
	BootDays uint64 `json:"bootDays"` // 运行天数
}

// GetNodeHostInfo 获取节点 Host 信息
func GetNodeHostInfo() HostInfo {
	hostInfo := HostInfo{
		BootDays: GetHostBootDays(),
	}

	return hostInfo
}

// 获取当前服务器的运行时间
func getHostBootTime() uint64 {
	bootTime, err := host.BootTime()

	if err != nil {
		fmt.Println(err)
	}

	// 当前时间戳
	curTime := uint64(time.Now().Unix())

	// 运行时间（秒）
	uptimeSeconds := curTime - bootTime

	return uptimeSeconds
}

// GetHostBootDays 天数
func GetHostBootDays() uint64 {

	bootTime := getHostBootTime()

	return bootTime / 60 / 60 / 24
}

// GetHostBootHours 小时数
func GetHostBootHours() uint64 {

	bootTime := getHostBootTime()

	return bootTime / 60 / 60
}

// GetHostBootMinutes 分钟数
func GetHostBootMinutes() uint64 {

	bootTime := getHostBootTime()

	return bootTime / 60
}

// GetHostBootSeconds 秒数
func GetHostBootSeconds() uint64 {

	bootTime := getHostBootTime()

	return bootTime
}
