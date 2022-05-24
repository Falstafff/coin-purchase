package bot

import (
	"fmt"
	"log"
)

type TradingBot struct {
	orderExecutor   *OrderExecutor
	exchangeService *ExchangeService
}

func NewTradingBot(serviceType ExchangeServiceType, processingType ProcessingStrategyType) *TradingBot {
	exchangeService := NewExchangeService(serviceType)
	processingStrategy := NewProcessingStrategy(processingType, &exchangeService)
	orderExecutor := NewOrderExecutor(processingStrategy)
	return &TradingBot{
		orderExecutor:   orderExecutor,
		exchangeService: &exchangeService,
	}
}

func (tb TradingBot) Trade(pair string) {
	ticker, err := tb.GetTicker(pair)

	if err != nil {
		log.Printf("ticker request error: %s", err)
		return
	}

	tb.orderExecutor.Ticker = ticker

	err = tb.orderExecutor.PlaceOrder()

	if err != nil {
		log.Printf("order placement error: %s", err)
		return
	}

	order := tb.orderExecutor.Order

	// todo save order to database
	log.Println(fmt.Sprintf("%#v", order))
}

func (tb TradingBot) GetTicker(pair string) (*CoinTicker, error) {
	ticker := &CoinTicker{}
	exchangeService := *tb.exchangeService

	log.Printf("got the next pair: %s", pair)

	lastPrice, err := exchangeService.GetLastPrice(pair)

	if err != nil {
		log.Println(err)
		return ticker, fmt.Errorf("last price is not found")
	}

	log.Printf("got the next last price: %s", lastPrice)

	ticker.Pair = pair
	ticker.LastPrice = lastPrice

	return ticker, nil
}
