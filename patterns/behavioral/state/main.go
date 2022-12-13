package main

import (
	"fmt"
	"log"
	"study/patterns/behavioral/state/pkg"
)

func main() {
	vendingMachine := pkg.NewVendingMachine(1, 10)
	err := vendingMachine.RequestTtem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InserMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println()

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println()
	err = vendingMachine.RequestTtem()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = vendingMachine.InserMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}
}
