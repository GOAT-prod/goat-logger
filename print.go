package goatlogger

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	moscowTimeLocation, _ = time.LoadLocation("Europe/Moscow")
)

func printLog(level LogLevel, app, tag, msg string) {
	info := logInfo{
		application: app,
		level:       level,
		time:        time.Now().In(moscowTimeLocation).Format(time.RFC3339),
		tag:         tag,
		msg:         msg,
	}

	infoBytes, _ := json.Marshal(info)
	fmt.Println(string(infoBytes))
}
