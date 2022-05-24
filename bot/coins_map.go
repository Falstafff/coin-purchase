package bot

import (
	"log"
)

type CoinsMap map[CoinBase]map[CoinQuote]PairId

type ExchangesCoinsMap map[ExchangeServiceType]CoinsMap

func CreateCoinsMap(pairs []CurrencyPair) CoinsMap {
	result := make(CoinsMap)

	for _, pair := range pairs {
		if result[pair.Base] == nil {
			result[pair.Base] = make(map[CoinQuote]PairId)
		}

		result[pair.Base][pair.Quote] = pair.Id
	}

	return result
}

func CreateExchangesCoinsMap() ExchangesCoinsMap {
	exchangesCoinsMap := make(ExchangesCoinsMap)
	exchangeServices := GetAllExchangeServices()

	for exchangeServiceType, exchangeService := range exchangeServices {
		pairs, err := exchangeService.GetAllPairs()

		if err != nil {
			log.Println(err)
			continue
		}

		exchangesCoinsMap[exchangeServiceType] = CreateCoinsMap(pairs)
	}

	return exchangesCoinsMap
}
