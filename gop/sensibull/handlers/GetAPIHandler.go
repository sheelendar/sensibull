// handlers/getAPIHandler.go
package handlers

import (
	"fmt"
	"gop/sensibull/consts"
	"gop/sensibull/dao"
	"gop/sensibull/logger"
	UnderlyingAssetHandler "gop/sensibull/underlyingAssetHandler"
	"gop/sensibull/utils"
	"time"
)

// GetUnderlyingPricesHandler return underlying detail from api or cache.
func GetUnderlyingPricesHandler() *dao.UnderlyingAsset {
	underlyingRes := dao.UnderlyingAsset{}
	// check underlying asset into db first
	err := utils.GetObjectFromRedis(consts.ALLUnderlyingAsset, &underlyingRes)
	if err == nil {
		logger.SensibullError{Message: "getting underlying asset response from redis"}.Info()
		return &underlyingRes
	}
	logger.SensibullError{Message: "underlying asset response not find in redis"}.Info()

	// make api call to get latest underlying asset details.
	response := UnderlyingAssetHandler.GetAllUnderlying()
	if response != nil {
		go func() {
			if err := utils.SaveObjectInRedis(consts.ALLUnderlyingAsset, response, time.Minute*15); err != nil {
				fmt.Println(err)
			}
		}()
	}
	_, err = UnderlyingAssetHandler.ProcessResponse(response)
	if err != nil {
		logger.SensibullError{Message: "error in processing underlying response"}.Err()
	}
	return response
}

// GetDerivativePricesHandler return derivatives details from cache or api.
func GetDerivativePricesHandler(symbol string) (*dao.UnderlyingAsset, error) {
	// check if the data is present in the cache
	cacheMap, err := utils.GetSymbolToTokenMapFromCache(consts.SymbolToTokenRedisKey)
	if err != nil || cacheMap == nil {
		cacheMap = UnderlyingAssetHandler.GetSymbolToTokenMap()
		if cacheMap == nil {
			return nil, fmt.Errorf("token not found for given syble in system")
		}
	}
	token, ok := cacheMap[symbol]
	if !ok {
		return nil, fmt.Errorf("invalid symbol")
	}

	underlyingRes, err := utils.GetDerivativeResponseFromDB(token)
	if err == nil {
		logger.SensibullError{Message: "getting derivative response from redis"}.Info()
		return underlyingRes, nil
	}

	logger.SensibullError{Message: "error while getting derivative response from redis: " + err.Error()}.Err()
	underlyingRes = UnderlyingAssetHandler.GetDerivativeDetailsByToken(token)
	if underlyingRes != nil && underlyingRes.Success {
		go func() {
			err = utils.SaveDerivativeResponseInDB(token, underlyingRes)
			if err != nil {
				logger.SensibullError{Message: "error while saving derivative response"}.Err()
			}
		}()
	}
	return underlyingRes, nil
}
