package pkg

import "fmt"

type AsusComputer struct {
	Memory int
	CPU    int
}

func (pc AsusComputer) PrindDetails() {
	fmt.Printf("[Asus] Pc CPU:[%d] Mem:[%d]\n", pc.CPU, pc.Memory)
}
