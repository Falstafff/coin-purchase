package main

import (
	"context"
	"fmt"
	"github.com/Projects/coin-purchase/modules/coin_purchase"
	"github.com/Projects/coin-purchase/modules/exchange_service"
)

func Handler(ctx context.Context) {
	orderExecutor := coin_purchase.NewOrderExecutor()

	err := orderExecutor.PlaceOrder()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//lambda.Start(Handler)

	//orderExecutor := coin_purchase.NewOrderExecutor()
	//
	//err := orderExecutor.PlaceOrder()
	//if err != nil {
	//	fmt.Println(err)
	//}

	gateService := exchange_service.NewGateService()
	result, _ := gateService.ListCandlesticks("BTC_USDT")

	fmt.Println(result)
}
