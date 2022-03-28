package modules

type OrderExecutor struct {
	NewState        State
	ProcessingState State
	FulfilledState  State
	CurrentState    State

	Ticker *CoinTicker
	Order  Order
}

func NewOrderExecutor(processingStrategy ProcessingStrategy) OrderExecutor {
	orderExecutor := &OrderExecutor{}

	newState := &NewState{
		StateName:     "new",
		OrderExecutor: orderExecutor,
	}
	processingState := &ProcessingState{
		StateName:          "processing",
		OrderExecutor:      orderExecutor,
		ProcessingStrategy: processingStrategy,
	}
	fulfilledState := &FulfilledState{
		StateName:     "fulfilled",
		OrderExecutor: orderExecutor,
	}

	orderExecutor.SetState(newState)
	orderExecutor.NewState = newState
	orderExecutor.ProcessingState = processingState
	orderExecutor.FulfilledState = fulfilledState

	return *orderExecutor
}

func (o *OrderExecutor) IsOrderCanBePlaced() (bool, error) {
	return o.CurrentState.IsOrderCanBePlaced()
}

func (o *OrderExecutor) IsOrderFulfilled() (bool, error) {
	return o.CurrentState.IsOrderFulfilled()
}

func (o *OrderExecutor) PlaceOrder() error {
	return o.CurrentState.PlaceOrder()
}

func (o *OrderExecutor) FinishOrder() error {
	return o.CurrentState.FinishOrder()
}

func (o *OrderExecutor) ProcessOrder() error {
	return o.CurrentState.ProcessOrder()
}

func (o *OrderExecutor) SetState(s State) {
	o.CurrentState = s
}

func (o *OrderExecutor) SetOrder(order Order) {
	o.Order = order
}
