package pkg

import "fmt"

type HasMoneyState struct {
	vendiingMachine *VendiingMachine
}

func (i *HasMoneyState) RequestTtem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *HasMoneyState) InserMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendiingMachine.itemCount = i.vendiingMachine.itemCount - 1
	if i.vendiingMachine.itemCount == 0 {
		i.vendiingMachine.setState(i.vendiingMachine.noItem)
	} else {
		i.vendiingMachine.setState(i.vendiingMachine.hasItem)
	}
	return nil
}
