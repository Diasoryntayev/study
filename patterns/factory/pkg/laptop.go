package pkg

import "fmt"

type Laptop struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewLaptop() Computer {
	return Laptop{
		Type:    LaptopType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (pc Laptop) GetType() string {
	return pc.Type
}

func (pc Laptop) PrindDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor:[%v]", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
