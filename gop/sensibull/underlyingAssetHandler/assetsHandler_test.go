package UnderlyingAssetHandler

import (
	"Go/src/gop/consts"
	"Go/src/gop/dao"
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetUnderlyingPricesHandler(t *testing.T) {
	mockResponse := dao.UnderlyingAsset{}
	bytes, _ := json.Marshal(mockResponse)
	httpmock.RegisterResponder("GET", consts.UnderlyingAssetURL, httpmock.NewStringResponder(200, string(bytes)))

	testCases := []struct {
		name        string
		expected    *dao.UnderlyingAsset
		expectedErr error
	}{
		{
			name:     "successCase",
			expected: &dao.UnderlyingAsset{},
		},
	}

	// Iterate through test cases and run the tests
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := executeAPIURL(consts.UnderlyingAssetURL)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expected, err)
		})

	}
}
