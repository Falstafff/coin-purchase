package modules

import (
	"fmt"
	"log"
)

type FulfilledState struct {
	StateName     string
	OrderExecutor *OrderExecutor
}

func (fs *FulfilledState) ProcessOrder() error {
	return fmt.Errorf("the order cannot be processed from the %s state", fs.StateName)
}

func (fs *FulfilledState) IsOrderCanBePlaced() (bool, error) {
	return false, nil
}

func (fs *FulfilledState) IsOrderFulfilled() (bool, error) {
	return true, nil
}

func (fs *FulfilledState) PlaceOrder() error {
	return fmt.Errorf("order cannot be placed from the %s state", fs.StateName)
}

func (fs *FulfilledState) FinishOrder() error {
	log.Println("the order has been placed")
	return nil
}
