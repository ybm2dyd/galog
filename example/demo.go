package main

import (
	"galog"
	"time"
)

func main() {
	galog.Init("./", galog.WhenDay, 1)
	playerlogin := new(galog.Playerlogin)
	playerlogin.EventTime = time.Now().Format("2006-01-02 15:04:05")
	playerlogin.Log()
	galog.Clean()
}
