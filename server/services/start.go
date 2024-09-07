package services

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

// ServerStart 服务启动器
func ServerStart(webContent fs.FS) {
	// 嵌入式的 Web 静态资源，这里的 web 是和根目录下的 web 目录名称一致，才能正确映射
	webFs, err := fs.Sub(webContent, "web")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(webFs))

	// 注册 Node 节点 WebSocket 处理器
	http.HandleFunc("/node", nodeWebSocketHandler)
	// 注册 Web 服务 WebSocket 处理器
	http.HandleFunc("/web", webWSHandler)
	// 注册 Web 服务
	http.Handle("/", fs)

	// 启动 Node 节点 WebSocket 服务
	port := ":8080"
	fmt.Println("Node server started at ws://localhost" + port + "/node")
	fmt.Println("Web server started at ws://localhost" + port + "/web")
	fmt.Println("Dashboard server started at http://localhost" + port + "/")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("Error starting server:", err)
	}
}
