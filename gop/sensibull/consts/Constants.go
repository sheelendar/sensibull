package consts

const (
	// External api constants
	UnderlyingAssetURL  = "https://prototype.sbulltech.com/api/underlyings"
	DerivativesPriceURL = "https://prototype.sbulltech.com/api/derivatives"
	WebSocketServerURL  = "wss://prototype.sbulltech.com/api/ws"

	// Internal API constants
	UnderlyingPriceURL = "/underlying-prices"
	DerivativePriceURL = "/api/derivative-prices/{symbol}"

	//redis constants
	SymbolToTokenRedisKey = "SymbolToTokenRedisKey"
	DerivativeKey         = "derivative_%d"
	ALLUnderlyingAsset    = "allUnderlyingAsset"
	DTUTKM                = "dtutkm"
	RedisHostAndPort      = "redis:6379"

	//server constants
	HostAndPort   = ":19093"
	Symbol        = "symbol"
	DataTypeQuote = "quote"
	DataTypePing  = "ping"
	DataTypeError = "error"

	//jobs constants
	DefaultNumOfWorkers = 10
	DefaultNumOfJobs    = 10

	//msg
	ErrInDerivativeResponse        = "error occurred into derivative response"
	ErrInParsingDerivativeResponse = "Failed to create JSON from derivative response"
	ErrInUnderlyingResponse        = "error occurred into underlying response"
	ErrInParsingUnderlyingResponse = "Failed to create JSON from underlyingHandler response"
)
