package bot

type State interface {
	IsOrderCanBePlaced() (bool, error)
	IsOrderFulfilled() (bool, error)
	ProcessOrder() error
	PlaceOrder() error
	FinishOrder() error
}
