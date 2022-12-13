package pkg

import "fmt"

type ItemRequestedState struct {
	vendiingMachine *VendiingMachine
}

func (i *ItemRequestedState) RequestTtem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestedState) InserMoney(money int) error {
	if money < i.vendiingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. PPlease insert %d", i.vendiingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendiingMachine.setState(i.vendiingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
