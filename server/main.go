package main

import (
	"LiveProbe/server/services"
	"embed"
)

// 将静态资源直接打包
//
//go:embed web
var webContent embed.FS

func main() {

	services.ServerStart(webContent)

}
