package main

import (
	"encoding/json"
	"github.com/Projects/coin-purchase/bot"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func Handler() {
	exchangesCoinMap, err := json.Marshal(bot.CreateExchangesCoinsMap())

	if err != nil {
		log.Fatalf("could not create exchange coin map json: %v", err)
		return
	}

	err = bot.PutJsonDataToBucket(bot.ExchangesCoinsMapFileName, exchangesCoinMap)

	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {
	lambda.Start(Handler)
}
