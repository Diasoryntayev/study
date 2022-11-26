package main

import "study/patterns/behavioral/observer/pkg"

func main() {
	sub := &pkg.Subscriber{Name: "Sub"}
	sub2 := &pkg.Subscriber{Name: "Sub-2"}
	sub3 := &pkg.Subscriber{Name: "Sub-3"}
	sub4 := &pkg.Subscriber{Name: "Sub-4"}

	chanel := pkg.Publisher{
		Name:      "Publisher channel",
		Consumers: make(map[string]pkg.Consumer),
	}
	chanel.Subscribe(sub)
	chanel.Subscribe(sub2)
	chanel.Subscribe(sub3)
	chanel.Subscribe(sub4)
	chanel.Notify()
	println("UnSubscribe Sub-3")
	chanel.UnSubscribe(sub3)
	chanel.Notify()
}
