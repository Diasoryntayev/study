package pkg

type VendiingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendiingMachine {
	v := &VendiingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &HasItemState{
		vendiingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendiingMachine: v,
	}

	hasMoneyState := &HasMoneyState{
		vendiingMachine: v,
	}

	noItemState := &NoItemState{
		vendiingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendiingMachine) RequestTtem() error {
	return v.currentState.RequestTtem()
}

func (v *VendiingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendiingMachine) InserMoney(money int) error {
	return v.currentState.InserMoney(money)
}

func (v *VendiingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendiingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendiingMachine) IncrementItemCount(count int) {
	v.itemCount = v.itemCount + count
}
