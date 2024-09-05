package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// GetCpuPercent 获取 CPU 使用率
func GetCpuPercent() float64 {
	osName := runtime.GOOS

	if osName == "windows" {
		return getCpuPercentForWindows()
	} else {
		return getCpuPercentForUnix()
	}
}

func getCpuPercentForWindows() float64 {
	// 采样时间设置为 1 秒
	percent, err := cpu.Percent(1*time.Second, false)

	if err != nil {
		return -1
	}

	return percent[0]
}

func getCpuPercentForUnix() float64 {
	out, err := exec.Command("top", "-l", "1", "-stats", "cpu").Output()
	if err != nil {
		return -1
	}

	output := string(out)
	lines := strings.Split(output, "\n")

	var cpuUsage float64
	for _, line := range lines {
		if strings.HasPrefix(line, "CPU usage:") {
			parts := strings.Fields(line)
			if len(parts) > 2 {
				usageStr := strings.Trim(parts[2], "%")
				cpuUsage, err = strconv.ParseFloat(usageStr, 64)
				if err != nil {
					return -1
				}
				break
			}
		}
	}

	return cpuUsage
}
