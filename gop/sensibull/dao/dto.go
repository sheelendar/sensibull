package dao

// UnderlyingAsset contains Assets data coming from api.
type UnderlyingAsset struct {
	Success bool   `json:"success"`
	Payload []Data `json:"payload"`
}

// Data contains asset details.
type Data struct {
	Symbol         string  `json:"symbol"`
	Underlying     string  `json:"underlying"`
	Token          int     `json:"token"`
	InstrumentType string  `json:"instrument_type"`
	Expiry         string  `json:"expiry"`
	Strike         float64 `json:"strike"`
}

// SubscribeQuote is used to subscribe channel.
type SubscribeQuote struct {
	MsgCommand string `json:"msg_command"`
	DataType   string `json:"data_type"`
	Quote      []int  `json:"quote"`
}

// PingServerMessage is used to receive ping msg from server.
type PingServerMessage struct {
	DataType string `json:"data_type"`
	Payload  string `json:"payload"`
}

// QuoteServerMessage is used to receive quote msg from server.
type QuoteServerMessage struct {
	DataType string `json:"data_type"`
	Payload  Quote  `json:"payload"`
}

// Quote contains price and token for a derivative.
type Quote struct {
	Token int     `json:"token"`
	Price float64 `json:"price"`
}

// DBUnderlyingAssetObject contains Assets data to store in db.
type DBUnderlyingAssetObject struct {
	Success             bool         `json:"success"`
	DerivativeToDataMap map[int]Data `json:"payload"`
}

// SubscribedDetails contains subscribed details.
type SubscribedDetails struct {
	ShareToken   int  `json:"shareToken"`
	IsSubscribed bool `json:"isSubscribed"`
}
