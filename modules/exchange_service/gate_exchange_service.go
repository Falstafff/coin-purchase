package exchange_service

type GateExchangeService struct {
}

func (ges *GateExchangeService) Purchase(pair string, lastPrice float32, amountUsd float32, margin float32) {

}

func (ges *GateExchangeService) PlaceConditionalSellOrder(pair string, lastPrice float32, amountUsd float32) {
}

func (ges *GateExchangeService) GetPriceBeforeNews(pair string) {}

func (ges *GateExchangeService) GetPair(pair string) {}
