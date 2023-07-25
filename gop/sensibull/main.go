package main

import (
	"fmt"
	"gop/sensibull/consts"
	"gop/sensibull/internal"
	"gop/sensibull/logger"
	"gop/sensibull/utils"

	"net/http"
)

func main() {
	redisCli := utils.InitRedis(consts.RedisHostAndPort, "", 0)
	internal.InitHttpClient()
	fmt.Println("Server listening on :", consts.HostAndPort)

	if err := http.ListenAndServe(consts.HostAndPort, nil); err != nil {
		fmt.Println(err)
		logger.SensibullError{Message: "Not able to start server port check on priority", ErrorCode: http.StatusInternalServerError}.Err()
	}

	defer func() {
		if redisCli != nil {
			redisCli.Close()
		}
	}()
}
