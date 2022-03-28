package modules

import (
	"fmt"
	"log"
)

type NewState struct {
	StateName     string
	OrderExecutor *OrderExecutor
}

func (ns *NewState) ProcessOrder() error {
	return fmt.Errorf("order cannot be processed from the %s state", ns.StateName)
}

func (ns *NewState) IsOrderCanBePlaced() (bool, error) {
	return true, nil
}

func (ns *NewState) IsOrderFulfilled() (bool, error) {
	return false, nil
}

func (ns *NewState) PlaceOrder() error {
	log.Println("the order is ready to be placed")
	ns.OrderExecutor.SetState(ns.OrderExecutor.ProcessingState)
	return ns.OrderExecutor.ProcessOrder()
}

func (ns *NewState) FinishOrder() error {
	return fmt.Errorf("order cannot be finished from the %s state", ns.StateName)
}
