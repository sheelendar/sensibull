package main

import (
	"gop/sensibull/consts"
	"gop/sensibull/utils"
)

func main() {
	redisCli := utils.InitRedis(consts.RedisHostAndPort, "", 0)
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
