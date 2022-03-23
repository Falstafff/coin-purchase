package exchange_service

import (
	"context"
	"fmt"
	"github.com/gateio/gateapi-go/v6"
)

type GateApiService struct {
	apiClient *gateapi.APIClient
}

func NewApiService() GateApiService {
	apiClient := gateapi.NewAPIClient(gateapi.NewConfiguration())
	GateApiService := GateApiService{apiClient}
	return GateApiService
}

func (gs *GateApiService) ListCurrencyPairs() ([]gateapi.CurrencyPair, error) {
	ctx := createStandardContext()
	result, _, err := gs.apiClient.SpotApi.ListCurrencyPairs(ctx)
	return result, handleGateError(err)
}

func (gs *GateApiService) CreateOrder(order gateapi.Order) (gateapi.Order, error) {
	ctx := createAuthContext()
	result, _, err := gs.apiClient.SpotApi.CreateOrder(ctx, order)
	return result, handleGateError(err)
}

func (gs *GateApiService) CreateConditionalOrder(order gateapi.SpotPriceTriggeredOrder) (gateapi.TriggerOrderResponse, error) {
	ctx := createAuthContext()
	result, _, err := gs.apiClient.SpotApi.CreateSpotPriceTriggeredOrder(ctx, order)
	return result, handleGateError(err)
}

func (gs *GateApiService) ListCandlesticks(pair string) ([][]string, error) {
	cxt := createStandardContext()
	result, _, err := gs.apiClient.SpotApi.ListCandlesticks(cxt, pair, nil)
	return result, handleGateError(err)
}

func (gs *GateApiService) ListTrades(pair string) ([]gateapi.Trade, error) {
	cxt := createStandardContext()
	result, _, err := gs.apiClient.SpotApi.ListTrades(cxt, pair, nil)
	return result, handleGateError(err)
}

func createStandardContext() context.Context {
	return context.Background()
}

func createAuthContext() context.Context {
	return context.WithValue(context.Background(),
		gateapi.ContextGateAPIV4,
		gateapi.GateAPIV4{
			Key:    "YOUR_API_KEY",
			Secret: "YOUR_API_SECRET",
		},
	)
}

func handleGateError(err error) error {
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			return fmt.Errorf("gate api error: %s\n", e.Error())
		} else {
			return fmt.Errorf("generic error: %s\n", err.Error())
		}
	}

	return nil
}
