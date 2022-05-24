package bot

const (
	Gate = "gate"
)

type ExchangeServiceType string

type PairId string

type CoinBase string

type CoinQuote string

type CurrencyPair struct {
	Id    PairId
	Base  CoinBase
	Quote CoinQuote
}

type ExchangeService interface {
	Purchase(ticker *CoinTicker, amountUsd string) (Order, error)
	PlaceConditionalSellOrder(ticker *CoinTicker, amountUsd float32) (Order, error)
	GetPriceBeforeNews(pair string) (string, error)
	GetLastPrice(pair string) (string, error)
	GetPair(base string, quote string) (string, error)
	GetAllPairs() ([]CurrencyPair, error)
}

func NewExchangeService(serviceType ExchangeServiceType) ExchangeService {
	switch serviceType {
	case Gate:
		return &GateExchangeService{NewGateApiService()}
	default:
		return nil
	}
}

type ExchangeServicesMap map[ExchangeServiceType]ExchangeService

func GetAllExchangeServices() ExchangeServicesMap {
	exchangeServicesMap := make(ExchangeServicesMap)
	exchangeServicesMap[Gate] = NewExchangeService(Gate)
	return exchangeServicesMap
}
