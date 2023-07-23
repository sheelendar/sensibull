package main

import (
	"Go/src/gop/sensibull/consts"
	"Go/src/gop/sensibull/internal"
	"Go/src/gop/sensibull/logger"
	"Go/src/gop/sensibull/utils"
	"fmt"
	"net/http"
)

func main() {

	utils.InitRedis(consts.RedisHostAndPort, "", 0)
	internal.InitHttpClient()
	fmt.Println("Server listening on :", consts.HostAndPort)

	if err := http.ListenAndServe(consts.HostAndPort, nil); err != nil {
		fmt.Println(err)
		logger.SensibullError{Message: "Not able to start server port check on priority", ErrorCode: http.StatusInternalServerError}.Err()
	}
}
