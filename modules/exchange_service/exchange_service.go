package exchange_service

type ExchangeService interface {
	Purchase(pair string, lastPrice float32, amountUsd float32, margin float32)
	PlaceConditionalSellOrder(pair string, lastPrice float32, amountUsd float32)
	GetPriceBeforeNews(pair string)
	GetPair(pair string)
}
