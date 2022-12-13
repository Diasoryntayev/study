package pkg

import "fmt"

type HasItemState struct {
	vendiingMachine *VendiingMachine
}

func (i *HasItemState) RequestTtem() error {
	if i.vendiingMachine.itemCount == 0 {
		i.vendiingMachine.setState(i.vendiingMachine.noItem)
		return fmt.Errorf("No item present")
	}

	fmt.Printf("Item requested\n")
	i.vendiingMachine.setState(i.vendiingMachine.itemRequested)
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendiingMachine.IncrementItemCount(count)
	return nil
}

func (i *HasItemState) InserMoney(money int) error {
	return fmt.Errorf("Please select item firs")
}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item firs")
}
