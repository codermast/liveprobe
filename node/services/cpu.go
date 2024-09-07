package services

import (
	"bufio"
	"bytes"
	"github.com/shirou/gopsutil/cpu"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CPUInfo struct {
	ModelName   string  `json:"modelName"`   // 型号
	Core        int32   `json:"core"`        // 核心数
	Mhz         float64 `json:"mhz"`         // 频率
	CacheSize   int32   `json:"cacheSize"`   // 缓存大小
	UsedPercent float64 `json:"usedPercent"` // 使用率
}

// 这里库提供的获取 CPU 信息的方法只能在 Windows 平台下成功运行，类 Unix 系统自己实现

// GetNodeCPUInfo 获取节点 CPU 信息
func GetNodeCPUInfo() CPUInfo {
	cpuInfo := CPUInfo{
		ModelName:   getCPUModelName(),
		Core:        getCPUCore(),
		Mhz:         getCPUMhz(),
		CacheSize:   getCPUCacheSize(),
		UsedPercent: getCpuPercent(),
	}

	return cpuInfo
}

// 获取 CPU 型号
func getCPUModelName() string {
	switch runtime.GOOS {
	case "windows":
		return getCPUModelNameForWindows()
	case "linux":
		return getCPUModelNameForLinux()
	case "darwin":
		return getCPUModelNameForMacOS()
	default:
		return "Unsupported OS"
	}
}
func getCPUModelNameForWindows() string {
	cpuInfo, err := cpu.Info()

	if err != nil {
		return "-1"
	}

	cpuModelName := cpuInfo[0].ModelName
	return cpuModelName
}
func getCPUModelNameForLinux() string {
	// 从 /proc/cpuinfo 文件中获取 CPU 型号
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "-1"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "model name") {
			// line 示例：model name  : Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "-1"
	}

	return "Unknown CPU Model"
}
func getCPUModelNameForMacOS() string {
	// getCPUModelNameForMacOS 使用 sysctl 命令获取 CPU 型号

	cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "-1"
	}

	cpuModelName := strings.TrimSpace(out.String())
	if cpuModelName == "" {
		return "Unknown CPU Model"
	}
	return cpuModelName
}

// 获取 CPU 核心数
func getCPUCore() int32 {
	switch runtime.GOOS {
	case "windows":
		return getCPUCoreForWindows()
	case "linux":
		return getCPUCoreForLinux()
	case "darwin":
		return getCPUCoreForMacOS()
	default:
		return 0
	}
}
func getCPUCoreForWindows() int32 {
	cpuInfo, err := cpu.Info()

	if err != nil {
		return 0
	}

	cpuCores := cpuInfo[0].Cores
	return cpuCores
}
func getCPUCoreForLinux() int32 {
	// 从 /proc/cpuinfo 获取 CPU 核心数
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return 0
	}
	defer file.Close()

	var cores int32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu cores") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				coreCount, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
				cores = int32(coreCount)
				break
			}
		}
	}
	return cores
}
func getCPUCoreForMacOS() int32 {
	// 使用 sysctl 获取 CPU 核心数
	cmd := exec.Command("sysctl", "-n", "hw.physicalcpu")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0
	}

	coreCount, _ := strconv.Atoi(strings.TrimSpace(out.String()))
	return int32(coreCount)
}

// 获取 CPU 主频
func getCPUMhz() float64 {
	switch runtime.GOOS {
	case "windows":
		return getCPUMhzForWindows()
	case "linux":
		return getCPUMhzForLinux()
	case "darwin":
		return getCPUMhzForMacOS()
	default:
		return -1
	}
}
func getCPUMhzForWindows() float64 {
	cpuInfo, err := cpu.Info()

	if err != nil {
		return -1
	}

	cpuMhz := cpuInfo[0].Mhz
	return cpuMhz
}
func getCPUMhzForLinux() float64 {
	// 从 /proc/cpuinfo 获取 CPU 主频
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu MHz") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				mhz, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
				return mhz
			}
		}
	}
	return -1
}
func getCPUMhzForMacOS() float64 {
	// getCPUMhzForMacOS 使用 sysctl 获取 CPU 主频

	cmd := exec.Command("sysctl", "-n", "hw.cpufrequency")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return -1
	}

	// 返回值是 Hz，我们将其转换为 MHz
	hz, _ := strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
	return hz / 1e6
}

// 获取 CPU 缓存大小
func getCPUCacheSize() int32 {
	switch runtime.GOOS {
	case "windows":
		return getCPUCacheSizeForWindows()
	case "linux":
		return getCPUCacheSizeForLinux()
	case "darwin":
		return getCPUCacheSizeForMacOS()
	default:
		return -1
	}
}
func getCPUCacheSizeForWindows() int32 {
	cpuInfo, err := cpu.Info()

	if err != nil {
		return -1
	}

	cpuCacheSize := cpuInfo[0].CacheSize
	return cpuCacheSize
}
func getCPUCacheSizeForLinux() int32 {
	// 从 /proc/cpuinfo 获取 CPU 缓存大小
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cache size") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				cacheSize, _ := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[1], " KB")))
				return int32(cacheSize)
			}
		}
	}
	return -1
}
func getCPUCacheSizeForMacOS() int32 {
	// getCPUCacheSizeForMacOS 使用 sysctl 获取 CPU 缓存大小

	cmd := exec.Command("sysctl", "-n", "hw.cachesize")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return -1
	}

	cacheSize, _ := strconv.Atoi(strings.TrimSpace(out.String()))
	return int32(cacheSize / 1024) // 返回 KB
}

// getCpuPercent 获取 CPU 使用率
func getCpuPercent() float64 {
	switch runtime.GOOS {
	case "windows":
		return getCpuPercentForWindows()
	case "linux":
		return getCpuPercentForLinux()
	case "darwin":
		return getCpuPercentForMacOS()
	default:
		return -1
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
func getCpuPercentForLinux() float64 {
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
func getCpuPercentForMacOS() float64 {
	return getCpuPercentForLinux()
}
