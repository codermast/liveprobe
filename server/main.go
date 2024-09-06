package main

import (
	"LiveProbe/server/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		services.StartWebSocketServer()
	}()

	wg.Wait()
}
