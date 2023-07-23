package UnderlyingAssetHandler

import (
	"Go/src/gop/sensibull/consts"
	"Go/src/gop/sensibull/dao"
	"Go/src/gop/sensibull/logger"
	"Go/src/gop/sensibull/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func getUnderlyingDataFromAPI(URL string) (*dao.UnderlyingAsset, error) {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, fmt.Errorf("error while requesting ungerlying asset prices")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("API returned a non-200 status code:", resp.Status)
		return nil, fmt.Errorf("error into response from underlying asset")
	}

	var data dao.UnderlyingAsset
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("not able to parse underlying asset price response")
	}
	fmt.Println("ID:", data)
	return &data, nil
}

func GetAllUnderlying() *dao.UnderlyingAsset {

	res, err := getUnderlyingDataFromAPI(consts.UnderlyingAssetURL)
	if err != nil {
		logger.SensibullError{Message: err.Error(), ErrorCode: http.StatusInternalServerError}.Err()
		return nil
	}
	err = processResponse(res)
	if err != nil {
		logger.SensibullError{Message: err.Error(), ErrorCode: http.StatusInternalServerError}.Err()
		return nil
	}
	return res
}

func processResponse(res *dao.UnderlyingAsset) error {
	if res == nil || res.Payload == nil {
		return fmt.Errorf("underlyingAsset is nil")
	}
	symbolToTokenMap := make(map[string]int)
	for i := 0; i < len(res.Payload); i++ {
		symbolToTokenMap[res.Payload[i].Symbol] = res.Payload[i].Token
	}
	err := utils.SaveObjectsInRedis(consts.SymbolToTokenRedisKey, symbolToTokenMap, time.Minute)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("not able to store SymbolToTokenMap into db")
	}
	return nil
}

func GetDerivativeDetailsByToken(token int) *dao.UnderlyingAsset {
	res, err := getUnderlyingDataFromAPI(fmt.Sprintf("%s/%d", consts.DerivativesPriceURL, token))
	if err != nil {
		logger.SensibullError{Message: err.Error(), ErrorCode: http.StatusInternalServerError}.Err()
		return nil
	}
	return res
}
