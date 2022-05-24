package bot

type TradeTicker struct {
	Pair     PairId              `json:"pair"`
	Platform ExchangeServiceType `json:"platform"`
}

func GetTradeTickers(symbols []string, exchangesCoinsMap ExchangesCoinsMap) []TradeTicker {
	var tradeTickers []TradeTicker
	coinQuote := CoinQuote(ConfigInstance().DefaultQuote)

	for _, symbol := range symbols {
		coinBase := CoinBase(symbol)

		for exchangeService, coinMap := range exchangesCoinsMap {
			if pairId, hasCoinId := coinMap[coinBase][coinQuote]; hasCoinId {
				tradeTickers = append(tradeTickers, TradeTicker{Pair: pairId, Platform: exchangeService})
				break
			}
		}
	}

	return tradeTickers
}
