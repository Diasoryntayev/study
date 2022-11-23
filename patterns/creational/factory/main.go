package main

import "study/patterns/creational/factory/pkg"

var types = []string{pkg.PersonalCompType, pkg.LaptopType, pkg.ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrindDetails()
	}
}
