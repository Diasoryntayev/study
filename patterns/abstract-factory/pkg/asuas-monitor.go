package pkg

import "fmt"

type AsusMonitor struct {
	Size int
}

func (pc AsusMonitor) PrindDetails() {
	fmt.Printf("[Asus] Monitor Size:[%d]\n", pc.Size)
}
