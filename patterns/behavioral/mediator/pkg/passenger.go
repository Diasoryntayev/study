package pkg

import "fmt"

type Passenger struct {
	Dispatcher
}

func (g *Passenger) PermitArrive() {
	fmt.Println("Пассажиры: Занимайте места...")
	g.Arrive()
}

func (g *Passenger) Go() {
	fmt.Println("Пассажиры: Отправление!")
	g.Dispatcher.NotifyAboutGo()
}

func (g *Passenger) Arrive() {
	if !g.CanArrive(g) {
		fmt.Println("Пассажиры: отправление, задерживается...")
		return
	}
	fmt.Println("Пассажиры занимайте места!")
}
