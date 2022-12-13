package pkg

import "fmt"

type NoItemState struct {
	vendiingMachine *VendiingMachine
}

func (i *NoItemState) RequestTtem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	i.vendiingMachine.IncrementItemCount(count)
	i.vendiingMachine.setState(i.vendiingMachine.hasItem)
	return nil
}

func (i *NoItemState) InserMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
