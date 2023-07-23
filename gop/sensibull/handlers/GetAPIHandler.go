// handlers/GetAPIHandler.go
package handlers

import (
	"Go/src/gop/sensibull/consts"
	"Go/src/gop/sensibull/dao"
	UnderlyingAssetHandler "Go/src/gop/sensibull/underlyingAssetHandler"
	"Go/src/gop/sensibull/utils"
	"fmt"
	"time"
)

func GetUnderlyingPricesHandler() *dao.UnderlyingAsset {
	underlyingRes := dao.UnderlyingAsset{}
	// check underlying asset into db first
	err := utils.GetObjectFromRedis(consts.ALLUnderlyingAsset, &underlyingRes)
	if err == nil {
		fmt.Println("getting underlying asset response from redis")
		return &underlyingRes
	}
	fmt.Println("underlying asset response not find in redis")

	// make api call to get latest underlying asset details.
	response := UnderlyingAssetHandler.GetAllUnderlying()
	if response != nil {
		go func() {
			if err := utils.SaveObjectsInRedis(consts.ALLUnderlyingAsset, response, time.Minute); err != nil {
				fmt.Println(err)
			}
		}()
	}
	return response
}

func GetDerivativePricesHandler(symbol string) (*dao.UnderlyingAsset, error) {
	// Check if the data is present in the cache
	cacheMap, err := utils.GetSymbolToTokenMapFromCache(consts.SymbolToTokenRedisKey)
	if err != nil {
		return nil, nil
	}
	token, ok := cacheMap[symbol]
	if !ok {
		return nil, fmt.Errorf("invalid symbol")
	}

	derivativeKey := fmt.Sprintf(consts.DerivativeKey, token)
	underlyingRes := dao.UnderlyingAsset{}

	err = utils.GetObjectFromRedis(derivativeKey, &underlyingRes)
	if err == nil {
		fmt.Println("getting derivative response from redis")
		return &underlyingRes, nil
	}

	fmt.Println("error while getting derivative response from redis")
	response := UnderlyingAssetHandler.GetDerivativeDetailsByToken(token)
	if response != nil {
		err = utils.SaveObjectsInRedis(derivativeKey, response, time.Minute)
		if err != nil {
			fmt.Println("error while saving derivative response")
		}
	}
	return response, nil
}
