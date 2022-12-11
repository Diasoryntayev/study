package pkg

type State interface {
	AddItem(int) error
	RequestTtem() error
	InserMoney(money int) error
	DispenseItem() error
}
