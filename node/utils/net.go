package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

type NetInfo struct {
	BytesSent     uint64  `json:"bytesSent"`     // 字节发送数 B/s
	BytesRecv     uint64  `json:"bytesRecv"`     // 字节下载数 B/s
	PacketsSent   uint64  `json:"packetsSent"`   // 包上传数 个/s
	PacketsRecv   uint64  `json:"packetsRecv"`   // 包下载数 个/s
	UploadSpeed   float64 `json:"uploadSpeed"`   // 上传速度 B/s
	DownloadSpeed float64 `json:"downloadSpeed"` // 下载速度 B/s
}

// 指定采样时间间隔（1秒）
var interval = time.Second

// GetNodeNetInfo 获取网络 IO 信息
func GetNodeNetInfo() NetInfo {

	// 获取初始的网络 I/O 数据
	ioCounters1, err := net.IOCounters(false) // false 表示汇总所有接口的总数据
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	// 等待指定的时间间隔
	time.Sleep(interval)

	// 获取新的网络 I/O 数据
	ioCounters2, err := net.IOCounters(false)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	// 计算上传和下载的包数差值
	packetsSent := ioCounters2[0].PacketsSent - ioCounters1[0].PacketsSent
	packetsRecv := ioCounters2[0].PacketsRecv - ioCounters1[0].PacketsRecv

	// 计算上传和下载的字节数差值
	bytesSent := ioCounters2[0].BytesSent - ioCounters1[0].BytesSent
	bytesRecv := ioCounters2[0].BytesRecv - ioCounters1[0].BytesRecv

	// 速度 B/s
	uploadSpeed := float64(bytesSent)
	downloadSpeed := float64(bytesRecv)

	//// 打印上传和下载速度
	//fmt.Printf("Upload Speed: %.2f KB/s, Download Speed: %.2f KB/s\n", uploadSpeed / 1024, downloadSpeed / 1024)
	netInfo := NetInfo{
		BytesSent:     bytesSent,
		BytesRecv:     bytesRecv,
		PacketsSent:   packetsSent,
		PacketsRecv:   packetsRecv,
		UploadSpeed:   uploadSpeed,
		DownloadSpeed: downloadSpeed,
	}
	return netInfo
}
