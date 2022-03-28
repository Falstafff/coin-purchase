package modules

const (
	Gate = "gate"
)

type ExchangeServiceType string

type ExchangeService interface {
	Purchase(ticker *CoinTicker, amountUsd float32) (Order, error)
	PlaceConditionalSellOrder(ticker *CoinTicker, amountUsd float32) (Order, error)
	GetPriceBeforeNews(pair string) (string, error)
	GetLastPrice(pair string) (string, error)
	GetPair(base string, quote string) (string, error)
}
