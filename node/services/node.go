package services

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"path/filepath"
)

type NodeInfo struct {
	NodeId  string     `json:"nodeId"`
	Host    HostInfo   `json:"host"`
	CPU     CPUInfo    `json:"cpu"`
	Memory  MemoryInfo `json:"memory"`
	Swap    SwapInfo   `json:"swap"`
	Disk    DiskInfo   `json:"disk"`
	Network NetInfo    `json:"network"`
}

// GetNodeInfo 获取节点信息
func GetNodeInfo() NodeInfo {
	nodeInfo := NodeInfo{
		NodeId:  GetNodeId(),
		Host:    GetNodeHostInfo(),
		CPU:     GetNodeCPUInfo(),
		Memory:  GetNodeMemoryInfo(),
		Swap:    GetNodeSwapInfo(),
		Disk:    GetNodeDiskInfoInfo(),
		Network: GetNodeNetInfo(),
	}
	return nodeInfo
}

// GetNodeId 获取节点 ID
func GetNodeId() string {

	configInfo, err := getConfigInfo()

	if err != nil {
		fmt.Println("Error getting config info:", err)
		return ""
	}

	nodeId := configInfo.NodeId

	if nodeId == "" {
		// 生成 UUID
		newUUID := uuid.New()
		nodeId = newUUID.String()
		configInfo.NodeId = nodeId
		err := setConfigInfo(configInfo)
		if err != nil {
			return ""
		}
	}

	return nodeId
}

// GetNodeServerUrl 获取节点 Server URL
func GetNodeServerUrl() string {
	configInfo, err := getConfigInfo()
	if err != nil {
		fmt.Println("Error getting config info:", err)
		return ""
	}
	return configInfo.ServerUrl
}

// SetNodeServerUrl 设置节点 Server URL
func SetNodeServerUrl(serverUrl string) error {
	configInfo, err := getConfigInfo()
	if err != nil {
		fmt.Println("Error getting config info:", err)
		return err
	}
	configInfo.ServerUrl = serverUrl
	return setConfigInfo(configInfo)
}

type ConfigInfo struct {
	NodeId    string `json:"nodeId"`
	ServerUrl string `json:"serverUrl"`
}

// 配置文件名称
const configFileName = "nodeInfo.txt"

func getConfigFilePath() (string, error) {
	// 获取用户目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return "", err
	}
	configPath := filepath.Join(homeDir, ".liveProbe", configFileName)

	// 创建配置文件所在目录（如果不存在）
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return "", err
	}

	// 判断文件是否存在
	_, err = os.Stat(configPath)

	if os.IsNotExist(err) {
		err = ioutil.WriteFile(configPath, []byte("{\"nodeId\":\"\",\"serverUrl\":\"\"}"), 0644)
		if err != nil {
			return "", err
		}
	}

	return configPath, nil
}

func getConfigInfo() (ConfigInfo, error) {

	configPath, _ := getConfigFilePath()

	// 读取配置文件
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return ConfigInfo{}, err
	}
	var config ConfigInfo
	if err := json.Unmarshal(configData, &config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return ConfigInfo{}, err
	}
	return config, nil
}

func setConfigInfo(configInfo ConfigInfo) error {
	configPath, _ := getConfigFilePath()

	configData, err := json.Marshal(configInfo)
	if err != nil {
		fmt.Println("Error marshalling config info:", err)
		return err
	}

	if err := ioutil.WriteFile(configPath, configData, 0644); err != nil {
		fmt.Println("Error writing config info:", err)
		return err
	}
	return nil
}
