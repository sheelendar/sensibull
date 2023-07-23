package dao

// UnderlyingAsset contains Assets data coming from api.
type UnderlyingAsset struct {
	Success bool   `json:"success"`
	Payload []Data `json:"payload"`
}

// Data contains asset details.
type Data struct {
	Symbol         string `json:"symbol"`
	Underlying     string `json:"underlying"`
	Token          int    `json:"token"`
	InstrumentType string `json:"instrument_type"`
	Expiry         string `json:"expiry"`
	Strike         int    `json:"strike"`
}
