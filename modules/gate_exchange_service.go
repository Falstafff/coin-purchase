package modules

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"github.com/shopspring/decimal"
)

type GateExchangeService struct {
	gateApiService GateApiService
}

func NewGateExchangeService(serviceType ExchangeServiceType) ExchangeService {
	switch serviceType {
	case Gate:
		return &GateExchangeService{NewGateApiService()}
	default:
		return nil
	}
}

func (ges *GateExchangeService) Purchase(ticker *CoinTicker, amountUsd float32) (Order, error) {
	order := Order{}
	// todo take out to config
	margin := 1.02
	finalPrice := decimal.RequireFromString(ticker.LastPrice).Mul(decimal.NewFromFloat(margin))
	finalAmount := decimal.NewFromFloat32(amountUsd).Div(finalPrice)

	gateOrder, err := ges.gateApiService.CreateOrder(gateapi.Order{
		Account:      finalAmount.String(),
		Price:        finalPrice.String(),
		Side:         "buy",
		CurrencyPair: ticker.Pair,
	})

	if err != nil {
		return order, err
	}

	order.Id = gateOrder.Id
	order.Amount = gateOrder.Amount
	order.Price = gateOrder.Price
	order.CreatedAt = gateOrder.CreateTime
	order.Pair = gateOrder.CurrencyPair

	return order, nil
}

func (ges *GateExchangeService) PlaceConditionalSellOrder(ticker *CoinTicker, amountUsd float32) (Order, error) {
	order := Order{}
	return order, nil
}

func (ges *GateExchangeService) GetPriceBeforeNews(pair string) (string, error) {
	options := &gateapi.ListCandlesticksOpts{Limit: optional.NewInt32(4), Interval: optional.NewString("1m")}
	candlesticks, err := ges.gateApiService.ListCandlesticks(pair, options)

	if err != nil {
		return "", err
	}

	twoMinutesBeforeNewsCandle := candlesticks[1]
	initClosePrice := twoMinutesBeforeNewsCandle[2]

	return initClosePrice, nil
}

func (ges *GateExchangeService) GetPair(base string, quote string) (string, error) {
	allPairs, _ := ges.gateApiService.ListCurrencyPairs()

	ourPair, err := func() (gateapi.CurrencyPair, error) {
		for _, pair := range allPairs {
			if pair.Base == base && pair.Quote == quote {
				return pair, nil
			}
		}
		return gateapi.CurrencyPair{}, fmt.Errorf("pair not found")
	}()

	if err != nil {
		return "", err
	}

	return ourPair.Id, nil
}

func (ges *GateExchangeService) GetLastPrice(pair string) (string, error) {
	tickers, err := ges.gateApiService.ListTickers(pair)

	if err != nil {
		return "", err
	}

	return tickers[0].Last, nil
}
