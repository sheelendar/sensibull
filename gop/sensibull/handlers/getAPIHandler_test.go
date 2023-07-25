package handlers

import (
	"Go/src/gop/sensibull/consts"
	"encoding/json"
	"testing"

	"Go/src/gop/sensibull/dao"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetUnderlyingPricesHandler(t *testing.T) {
	mockResponse := dao.UnderlyingAsset{}
	bytes, _ := json.Marshal(mockResponse)
	httpmock.RegisterResponder("GET", consts.UnderlyingAssetURL, httpmock.NewStringResponder(200, string(bytes)))

	testCases := []struct {
		name     string
		expected *dao.UnderlyingAsset
	}{
		{
			name:     "successCase",
			expected: &dao.UnderlyingAsset{},
		},
	}

	// Iterate through test cases and run the tests
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := GetUnderlyingPricesHandler()
			assert.Equal(t, test.expected, result)
		})

	}
}
