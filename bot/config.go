package bot

import (
	"os"
	"sync"
)

var lock = &sync.Mutex{}

type config struct {
	GateApiKey    string
	GateSecretKey string

	PurchaseAmountInUsdt string
	DefaultQuote         string
	BuyPriceMargin       float64
}

var configInstance *config

func ConfigInstance() *config {
	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		configInstance = &config{
			GateApiKey:           os.Getenv("GATE_API_KEY"),
			GateSecretKey:        os.Getenv("GATE_SECRET_KEY"),
			PurchaseAmountInUsdt: os.Getenv("PURCHASE_AMOUNT_IN_USDT"),
			DefaultQuote:         "USDT",
			BuyPriceMargin:       1.02,
		}
	}

	return configInstance
}
