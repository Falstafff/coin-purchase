package bot

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"github.com/shopspring/decimal"
)

type GateExchangeService struct {
	gateApiService GateApiService
}

func (ges *GateExchangeService) Purchase(ticker *CoinTicker, amountUsd string) (Order, error) {
	order := Order{}
	margin := ConfigInstance().BuyPriceMargin
	finalPrice := decimal.RequireFromString(ticker.LastPrice).Mul(decimal.NewFromFloat(margin))
	finalAmount := decimal.RequireFromString(amountUsd).Div(finalPrice)

	gateOrder, err := ges.gateApiService.CreateOrder(gateapi.Order{
		Type:         "limit",
		Account:      "spot",
		Side:         "buy",
		Amount:       finalAmount.String(),
		Price:        finalPrice.String(),
		CurrencyPair: ticker.Pair,
		TimeInForce:  "gtc",
		AutoBorrow:   false,
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

func (ges *GateExchangeService) GetAllPairs() ([]CurrencyPair, error) {
	pairs, err := ges.gateApiService.ListCurrencyPairs()

	if err != nil {
		return nil, err
	}

	currencyPairs := make([]CurrencyPair, len(pairs))

	for i, pair := range pairs {
		currencyPairs[i] = CurrencyPair{
			Base:  CoinBase(pair.Base),
			Quote: CoinQuote(pair.Quote),
			Id:    PairId(pair.Id),
		}
	}

	return currencyPairs, nil
}
