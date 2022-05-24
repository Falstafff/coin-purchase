package services

import (
	"encoding/json"
	"github.com/Projects/coin-purchase/bot"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

func PutTradeEvent(tradeTickers []bot.TradeTicker) error {
	detail, err := json.Marshal(map[string][]bot.TradeTicker{
		"coins": tradeTickers,
	})

	if err != nil {
		return err
	}

	newSession, err := session.NewSession()

	if err != nil {
		return err
	}

	eb := eventbridge.New(newSession)

	entries := []*eventbridge.PutEventsRequestEntry{{
		Detail:     aws.String(string(detail)),
		DetailType: aws.String("BotTrade"),
		Source:     aws.String("bot.trade"),
	}}

	_, err = eb.PutEvents(&eventbridge.PutEventsInput{
		Entries: entries,
	})

	if err != nil {
		return err
	}

	return nil
}
