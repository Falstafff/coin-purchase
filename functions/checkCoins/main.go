package main

import (
	"github.com/Projects/coin-purchase/bot"
	"github.com/Projects/coin-purchase/internal/services"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"strings"
)

type ImportantNewsEvent struct {
	Detail struct {
		News []struct {
			CoinSymbols string `json:"coinSymbol"`
			Platform    string `json:"platform"`
		} `json:"news_checker"`
	} `json:"detail"`
}

func Handler(event ImportantNewsEvent) {
	news := event.Detail.News

	if len(news) == 0 {
		log.Println("important news_checker is empty")
		return
	}

	exchangesCoinsMap, err := bot.GetExchangeCoinsMap()

	if err != nil {
		log.Println(err)
		return
	}

	symbols := strings.Split(news[0].CoinSymbols, ",")

	if len(symbols) == 0 {
		log.Println("symbols not found")
		return
	}

	tradeTickers := bot.GetTradeTickers(symbols, exchangesCoinsMap)

	log.Printf("trade tickers was found: %v", tradeTickers)

	err = services.PutTradeEvent(tradeTickers)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("trade event sent")
}

func main() {
	lambda.Start(Handler)
}
