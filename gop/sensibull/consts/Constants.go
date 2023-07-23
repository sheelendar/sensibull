package consts

const (
	// External api constants
	UnderlyingAssetURL  = "https://prototype.sbulltech.com/api/underlyings"
	DerivativesPriceURL = "https://prototype.sbulltech.com/api/derivatives"

	// Internal API constants
	UnderlyingPriceURL = "/underlying-prices"
	DerivativePriceURL = "/api/derivative-prices/{symbol}"

	//redis constants
	SymbolToTokenRedisKey = "SymbolToTokenRedisKey"
	DerivativeKey         = "derivative_%d"
	ALLUnderlyingAsset    = "allUnderlyingAsset"
	RedisHostAndPort      = "localhost:6379"

	Minute      = 60
	Second      = 1
	Hour        = Minute * 60
	FiveMinutes = Minute * 5

	//server constants

	HostAndPort = ":19093"

	Symbol = "symbol"
)
