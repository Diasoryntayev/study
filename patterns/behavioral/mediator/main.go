package main

import (
	"study/patterns/behavioral/mediator/pkg"
	"time"
)

func main() {
	stationManager := pkg.NewStationManager()
	passengerBus := &pkg.Passenger{
		Dispatcher: stationManager,
	}
	corgo := &pkg.Cargo{
		Dispatcher: stationManager,
	}

	passengerBus.Arrive()
	time.Sleep(time.Second * 1)

	corgo.Arrive()
	time.Sleep(time.Second * 1)
	passengerBus.Go()
}
