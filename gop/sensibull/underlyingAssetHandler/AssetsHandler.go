package UnderlyingAssetHandler

import (
	"encoding/json"
	"fmt"
	"gop/sensibull/consts"
	"gop/sensibull/dao"
	"gop/sensibull/logger"
	"gop/sensibull/utils"
	"net/http"
	"time"
)

// executeAPIURL call get url and return UnderlyingAsset response.
func executeAPIURL(URL string) (*dao.UnderlyingAsset, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("error while requesting ungerlying asset prices")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error into response from underlying asset")
	}

	var data dao.UnderlyingAsset
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("not able to parse underlying asset price response" + err.Error())
	}
	logger.SensibullError{Message: fmt.Sprintf("Data: %v", data)}.Err()
	return &data, nil
}

// GetAllUnderlying rerun price and detail of all underlying share.
func GetAllUnderlying() *dao.UnderlyingAsset {
	res, err := executeAPIURL(consts.UnderlyingAssetURL)
	if err != nil {
		logger.SensibullError{Message: err.Error(), ErrorCode: http.StatusInternalServerError}.Err()
		return nil
	}
	return res
}

// ProcessResponse process UnderlyingAsset response and same symbol mapping into cache.
func ProcessResponse(res *dao.UnderlyingAsset) (map[string]int, error) {
	if res == nil || res.Payload == nil {
		return nil, fmt.Errorf("underlyingAsset is nil")
	}
	symbolToTokenMap, _ := utils.GetSymbolToTokenMapFromCache(consts.SymbolToTokenRedisKey)
	if symbolToTokenMap == nil {
		symbolToTokenMap = map[string]int{}
	}
	for i := 0; i < len(res.Payload); i++ {
		if _, ok := symbolToTokenMap[res.Payload[i].Symbol]; !ok {
			symbolToTokenMap[res.Payload[i].Symbol] = res.Payload[i].Token
		}
	}
	err := utils.SaveObjectInRedis(consts.SymbolToTokenRedisKey, symbolToTokenMap, time.Minute*15)
	if err != nil {
		logger.SensibullError{Message: err.Error()}.Err()
		return nil, fmt.Errorf("not able to store SymbolToTokenMap into db")
	}
	return symbolToTokenMap, err
}

// InitRefreshRequiredCacheAndGetALLTokenMap build cache and refresh derivative cache in 1 minutes and underlying share cache in 15 minutes.
func InitRefreshRequiredCacheAndGetALLTokenMap() map[int]dao.SubscribedDetails {
	symbolToTokenMap := GetSymbolToTokenMap()
	derivativeTokenToUnderlyingTokenMap := make(map[int]dao.SubscribedDetails)
	for _, token := range symbolToTokenMap {
		res := GetDerivativeDetailsByToken(token)
		if res == nil || res.Payload == nil {
			continue
		}
		_, err := ProcessResponse(res)
		if err != nil {
			logger.SensibullError{Message: err.Error()}.Err()
		}
		err = utils.SaveDerivativeResponseInDB(token, res)
		if err != nil {
			logger.SensibullError{Message: err.Error()}.Err()
		}
		saveIntoMap(derivativeTokenToUnderlyingTokenMap, token, res.Payload)
	}
	derivativeTokenToUnderlyingTokenMap = SaveDerivativeTokenMapping(derivativeTokenToUnderlyingTokenMap)
	return derivativeTokenToUnderlyingTokenMap
}

func SaveDerivativeTokenMapping(derivativeTokenToUnderlyingTokenMap map[int]dao.SubscribedDetails) map[int]dao.SubscribedDetails {
	var tokenMap map[int]dao.SubscribedDetails
	err := utils.GetObjectFromRedis(consts.DTUTKM, tokenMap)
	if err != nil || len(tokenMap) == 0 {
		tokenMap = derivativeTokenToUnderlyingTokenMap
	} else {
		for k, v := range derivativeTokenToUnderlyingTokenMap {
			if _, ok := tokenMap[k]; !ok {
				tokenMap[k] = v
			}
		}
	}
	err = utils.SaveObjectInRedis(consts.DTUTKM, tokenMap, time.Hour)
	if err != nil {
		logger.SensibullError{Message: "error while saving derivativeTokenToUnderlyingTokenMap in db"}.Err()
	}
	return tokenMap
}

// saveIntoMap save mapping data into tokenMap.
func saveIntoMap(tokenMap map[int]dao.SubscribedDetails, token int, payload []dao.Data) {
	if payload == nil {
		return
	}
	for i := 0; i < len(payload); i++ {
		tokenMap[payload[i].Token] = dao.SubscribedDetails{ShareToken: token}
	}
}

// GetSymbolToTokenMap return symbol to token map from cache or api.
func GetSymbolToTokenMap() map[string]int {
	symbolToTokenMap, err := utils.GetSymbolToTokenMapFromCache(consts.SymbolToTokenRedisKey)
	if err == nil || symbolToTokenMap != nil {
		logger.SensibullError{Message: "getting symbolToTokenMap from cache"}.Info()
		return symbolToTokenMap
	}
	logger.SensibullError{Message: "symbolToTokenMap not find into cache"}.Info()
	res := GetAllUnderlying()
	if res == nil {
		logger.SensibullError{Message: "error while getting underling prices from api"}.Err()
		return nil
	}
	symbolToTokenMap, err = ProcessResponse(res)
	if err != nil || symbolToTokenMap == nil {
		logger.SensibullError{Message: "error while getting underling prices from api"}.Err()
		fmt.Println()
		return nil
	}
	return symbolToTokenMap
}

// GetDerivativeDetailsByToken return derivative prices by token from api.
func GetDerivativeDetailsByToken(token int) *dao.UnderlyingAsset {
	res, err := executeAPIURL(fmt.Sprintf("%s/%d", consts.DerivativesPriceURL, token))
	if err != nil {
		logger.SensibullError{Message: err.Error(), ErrorCode: http.StatusInternalServerError}.Err()
		return nil
	}
	return res
}
