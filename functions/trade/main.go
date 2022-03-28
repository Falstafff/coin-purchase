package main

import (
	"context"
	"github.com/Projects/coin-purchase/modules"
)

func Handler(ctx context.Context) {

}

func main() {
	base := "BTC"

	tradingBot := modules.NewTradingBot(modules.Gate, modules.PurchaseOnlyType)

	tradingBot.Trade(base)

	//bot := bot.NewTradingBot()
	//
	//bot

	//lambda.Start(Handler)

	//start := time.Now()
	//
	//gateExchangeService := exchange_service.NewGateExchangeService()
	//result, _ := gateExchangeService.GetPair("BTC", "USDT")
	//
	//fmt.Println(result)
	//elapsed := time.Since(start)
	//log.Printf("Binomial took %s", elapsed)
}
