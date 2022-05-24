package bot

import (
	"log"
)

const (
	PurchaseOnlyType         = "purchase_only"
	PurchaseWithStopLossType = "purchase_with_stop_loss"
)

type ProcessingStrategyType string

type ProcessingStrategy interface {
	Execute(ticker *CoinTicker) (Order, error)
}

func NewProcessingStrategy(strategyType ProcessingStrategyType, service *ExchangeService) ProcessingStrategy {
	switch strategyType {
	case PurchaseOnlyType:
		return &PurchaseOnly{ExchangeService: service}
	case PurchaseWithStopLossType:
		return &PurchaseWithStopLoss{ExchangeService: service}
	default:
		return nil
	}
}

type PurchaseOnly struct {
	ExchangeService *ExchangeService
}

func (op *PurchaseOnly) Execute(ticker *CoinTicker) (Order, error) {
	log.Printf("executing %s strategy", PurchaseOnlyType)
	return (*op.ExchangeService).Purchase(ticker, ConfigInstance().PurchaseAmountInUsdt)
}

type PurchaseWithStopLoss struct {
	ExchangeService *ExchangeService
}

func (ps *PurchaseWithStopLoss) Execute(ticker *CoinTicker) (Order, error) {
	log.Printf("executing %s strategy", PurchaseWithStopLossType)
	return (*ps.ExchangeService).PlaceConditionalSellOrder(ticker, 100)
}
