export interface NodeInfo {
	nodeId: string;
	host: HostInfo;
	cpu: CPUInfo;
	memory: MemoryInfo;
	swap: SwapInfo;
	disk: DiskInfo;
	network: NetInfo;
}


interface HostInfo {
	bootDays : number;  // 运行天数
}

interface CPUInfo {
	modelName: string; // 型号
	core: number; // 核心数
	mhz: number; // 频率
	cacheSize: number; // 缓存大小
	usedPercent: number; // 使用率
}

interface MemoryInfo {
	total: number; // 总内存大小
	available: number; // 可用内存大小
	used: number; // 已使用内存大小
	usedPercent: number; // 已使用内存百分比
}

interface SwapInfo {
	total: number;
	available: number;
	used: number;
	usedPercent: number;
}

interface DiskInfo {
	total: number; // 总内存大小
	available: number; // 可用内存大小
	used: number; // 已使用内存大小
	usedPercent: number; // 已使用内存百分比
	mountPath: string; // 磁盘挂载点
}

interface NetInfo {
	bytesSent: number; // 字节发送数 B/s
	bytesRecv: number; // 字节下载数 B/s
	packetsSent: number; // 包上传数 个/s
	packetsRecv: number; // 包下载数 个/s
	uploadSpeed: number; // 上传速度 B/s
	downloadSpeed: number; // 下载速度 B/s
}