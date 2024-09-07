package models

type NodeInfo struct {
	NodeId  string     `json:"nodeId"`
	CPU     CPUInfo    `json:"cpu"`
	Memory  MemoryInfo `json:"memory"`
	Swap    SwapInfo   `json:"swap"`
	Disk    DiskInfo   `json:"disk"`
	Network NetInfo    `json:"network"`
}

type CPUInfo struct {
	ModelName string  `json:"modelName"` // 型号
	Core      int32   `json:"core"`      // 核心数
	Mhz       float64 `json:"mhz"`       // 频率
	CacheSize int32   `json:"cacheSize"` // 缓存大小
	Percent   float64 `json:"percent"`   // 使用率
}

type MemoryInfo struct {
	Total       uint64  `json:"total"`       // 总内存大小
	Available   uint64  `json:"available"`   // 可用内存大小
	Used        uint64  `json:"used"`        // 已使用内存大小
	UsedPercent float64 `json:"usedPercent"` // 已使用内存百分比
}

type SwapInfo struct {
	Total       uint64  `json:"total"`       // 总内存大小
	Available   uint64  `json:"available"`   // 可用内存大小
	Used        uint64  `json:"used"`        // 已使用内存大小
	UsedPercent float64 `json:"usedPercent"` // 已使用内存百分比
}

type DiskInfo struct {
	Total       uint64  `json:"total"`       // 总内存大小
	Available   uint64  `json:"available"`   // 可用内存大小
	Used        uint64  `json:"used"`        // 已使用内存大小
	UsedPercent float64 `json:"usedPercent"` // 已使用内存百分比
	MountPath   string  `json:"mountPath"`   // 磁盘挂载点
}

type NetInfo struct {
	BytesSent     uint64  `json:"bytesSent"`     // 字节发送数 B/s
	BytesRecv     uint64  `json:"bytesRecv"`     // 字节下载数 B/s
	PacketsSent   uint64  `json:"packetsSent"`   // 包上传数 个/s
	PacketsRecv   uint64  `json:"packetsRecv"`   // 包下载数 个/s
	UploadSpeed   float64 `json:"uploadSpeed"`   // 上传速度 B/s
	DownloadSpeed float64 `json:"downloadSpeed"` // 下载速度 B/s
}
