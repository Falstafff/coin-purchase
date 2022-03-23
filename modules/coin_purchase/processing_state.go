package coin_purchase

import (
	"fmt"
	"log"
)

type ProcessingState struct {
	StateName     string
	OrderExecutor *OrderExecutor
}

func (ps *ProcessingState) ProcessOrder() error {
	log.Println("order processing")
	ps.OrderExecutor.SetState(ps.OrderExecutor.FulfilledState)
	return ps.OrderExecutor.FinishOrder()
}

func (ps *ProcessingState) IsOrderCanBePlaced() (bool, error) {
	return false, nil
}

func (ps *ProcessingState) IsOrderFulfilled() (bool, error) {
	return false, nil
}

func (ps *ProcessingState) PlaceOrder() error {
	return fmt.Errorf("order cannot be placed from the %s state", ps.StateName)
}

func (ps *ProcessingState) FinishOrder() error {
	return fmt.Errorf("the order cannot be finished from the %s state", ps.StateName)
}
