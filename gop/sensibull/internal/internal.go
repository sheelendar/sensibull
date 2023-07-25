package internal

import (
	"encoding/json"
	"fmt"
	"gop/sensibull/consts"
	"gop/sensibull/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

// InitHttpClient init internal url for upcoming request.
func InitHttpClient() {
	router.HandleFunc(consts.UnderlyingPriceURL, underlyingHandler).Methods("GET")
	router.HandleFunc(consts.DerivativePriceURL, derivativeHandler).Methods("GET")
	http.Handle("/", router)
}

// underlyingHandler call GetUnderlyingPricesHandler return all underlying response.
func underlyingHandler(response http.ResponseWriter, req *http.Request) {
	res := handlers.GetUnderlyingPricesHandler()
	if res == nil {
		http.Error(response, consts.ErrInUnderlyingResponse, http.StatusInternalServerError)
		return
	}

	// Convert the underlying data to JSON format
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(response, consts.ErrInParsingUnderlyingResponse, http.StatusInternalServerError)
		return
	}

	// Set the content type to application/json
	response.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	if _, err := response.Write(responseJSON); err != nil {
		fmt.Println(err)
	}
}

// derivativeHandler call GetDerivativePricesHandler and return derivative response.
func derivativeHandler(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	symbol := vars[consts.Symbol]

	res, err := handlers.GetDerivativePricesHandler(symbol)
	if res == nil || err != nil {
		http.Error(response, consts.ErrInDerivativeResponse, http.StatusInternalServerError)
		return
	}

	// Convert the Derivative data to JSON format
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(response, consts.ErrInParsingDerivativeResponse, http.StatusInternalServerError)
		return
	}

	// Set the content type to application/json
	response.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	response.Write(responseJSON)
}
