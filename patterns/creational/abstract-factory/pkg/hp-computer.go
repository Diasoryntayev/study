package pkg

import "fmt"

type HpComputer struct {
	Memory int
	CPU    int
}

func (pc HpComputer) PrintDetails() {
	fmt.Printf("[Hp] Pc CPU:[%d] Mem:[%d]\n", pc.CPU, pc.Memory)
}
