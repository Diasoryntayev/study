package main

import (
	"fmt"
	"study/patterns/structural/decarator/pkg"
)

var (
	base = pkg.BasePc{}
	home = pkg.HomePc{
		Cpu:          4,
		GraphicsCard: 1,
		Wrapper:      base,
	}

	server = pkg.ServerPc{
		Cpu:     16,
		Memory:  256,
		Wrapper: base,
	}
)

func main() {
	fmt.Printf("Base Price[%.2f]\n", base.GetPrice())
	fmt.Printf("Home Cpu:[%d] Cards:[%d] Price:[%.2f]\n", home.Cpu, home.GraphicsCard, home.GetPrice())
	fmt.Printf("Server Cpu:[%d] Memory:[%d] Price:[%.2f]\n", server.Cpu, server.Memory, server.GetPrice())
}
