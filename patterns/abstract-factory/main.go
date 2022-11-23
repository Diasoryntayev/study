package main

import "study/patterns/abstract-factory/pkg"

var brands = []string{pkg.Asus, pkg.Hp, "dell"}

func main() {
	for _, brand := range brands {
		factory, err := pkg.GetFactory(brand)
		if err != nil {
			println(err.Error())
			continue
		}
		monitor := factory.GetMonitor()
		monitor.PrintDetails()
		computer := factory.GetComputer()
		computer.PrintDetails()
	}
}
