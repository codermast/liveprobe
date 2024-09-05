package main

type CPUInfo struct {
	ModelName string  `json:"modelName"` // 型号
	Core      uint64  `json:"core"`      // 核心数
	Mhz       float64 `json:"mhz"`       // 频率
	CacheSize uint64  `json:"cacheSize"` // 缓存大小
	Percent   float64 `json:"percent"`   // 使用率
}

type NodeInfo struct {
	NodeId  string  `json:"nodeId"`
	CPU     CPUInfo `json:"cpu"`
	Memory  string  `json:"memory"`
	Disk    string  `json:"disk"`
	Network string  `json:"network"`
}

func main() {

}
