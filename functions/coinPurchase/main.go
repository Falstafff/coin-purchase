package main

import (
	"fmt"
	"github.com/Projects/coin-purchase/bot"
	"github.com/aws/aws-lambda-go/lambda"
)

type CoinPurchaseEvent struct {
	Pair     string `json:"pair"`
	Platform string `json:"platform"`
}

func Handler(event CoinPurchaseEvent) {
	fmt.Println(fmt.Sprintf("%#v", event))
	bot.NewTradingBot(bot.Gate, bot.PurchaseOnlyType).Trade(event.Pair)
}

func main() {
	lambda.Start(Handler)
}
