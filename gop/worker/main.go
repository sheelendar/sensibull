package main

import (
	"fmt"
	"github.com/sheelendar/src/sensibull/gop/consts"
	"github.com/sheelendar/src/sensibull/gop/utils"
)

func main() {
	redisCli := utils.InitRedis(consts.RedisHostAndPort, "", 0)
	fmt.Println("worker initiating :")
	webSConn := InitWebSocket()
	defer func() {
		if webSConn != nil {
			webSConn.Close()
		}
		if redisCli != nil {
			redisCli.Close()
		}
	}()
}
