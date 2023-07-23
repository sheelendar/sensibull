package internal

import (
	"Go/src/gop/sensibull/consts"
	"Go/src/gop/sensibull/handlers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

func InitHttpClient() {
	// Define the GET API route and its handler
	router.HandleFunc(consts.UnderlyingPriceURL, underlyingHandler).Methods("GET")
	router.HandleFunc(consts.DerivativePriceURL, derivativeHandler).Methods("GET")

	// Define the POST API route and its handler
	//router.HandleFunc("/post-data", derivativeHandler).Methods("POST")

	http.Handle("/", router)
}

func underlyingHandler(response http.ResponseWriter, req *http.Request) {
	res := handlers.GetUnderlyingPricesHandler()
	if res == nil {
		http.Error(response, "error occurred into underlying response", http.StatusInternalServerError)
		return
	}

	// Convert the underlying data to JSON format
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(response, "Failed to create JSON from underlyingHandler response", http.StatusInternalServerError)
		return
	}

	// Set the content type to application/json
	response.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	if _, err := response.Write(responseJSON); err != nil {
		fmt.Println(err)
	}
}

func derivativeHandler(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	symbol := vars[consts.Symbol]

	res, err := handlers.GetDerivativePricesHandler(symbol)
	if res == nil || err != nil {
		http.Error(response, "error occurred into derivative response", http.StatusInternalServerError)
		return
	}

	// Convert the Derivative data to JSON format
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(response, "Failed to create JSON from derivative response", http.StatusInternalServerError)
		return
	}

	// Set the content type to application/json
	response.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	response.Write(responseJSON)
}
