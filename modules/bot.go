package modules

import (
	"fmt"
	"log"
)

type TradingBot struct {
	orderExecutor   OrderExecutor
	exchangeService *ExchangeService
}

func NewTradingBot(serviceType ExchangeServiceType, processingType ProcessingStrategyType) TradingBot {
	exchangeService := NewGateExchangeService(serviceType)
	processingStrategy := NewProcessingStrategy(processingType, &exchangeService)

	return TradingBot{
		orderExecutor:   NewOrderExecutor(processingStrategy),
		exchangeService: &exchangeService,
	}
}

func (tb TradingBot) Trade(base string) {
	ticker, err := tb.GetTicker(base)

	if err != nil {
		log.Printf("ticker request error: %s", err)
		return
	}

	tb.orderExecutor.Ticker = &ticker
	err = tb.orderExecutor.PlaceOrder()

	if err != nil {
		log.Printf("order placement error: %s", err)
		return
	}

	order := tb.orderExecutor.Order

	// todo save order to database
	log.Println(order)
}

func (tb TradingBot) GetTicker(base string) (CoinTicker, error) {
	ticker := CoinTicker{}

	// todo take out quote to config file
	pair, err := (*tb.exchangeService).GetPair(base, "USDT")

	if err != nil {
		log.Println(err)
		return ticker, fmt.Errorf("pair is not found")
	}

	log.Printf("got the next pair: %s", pair)

	lastPrice, err := (*tb.exchangeService).GetLastPrice(pair)

	if err != nil {
		log.Println(err)
		return ticker, fmt.Errorf("last price is not found")
	}

	log.Printf("got the next last price: %s", lastPrice)

	ticker.Pair = pair
	ticker.LastPrice = lastPrice

	return ticker, nil
}
