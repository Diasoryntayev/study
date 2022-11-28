package main

import (
	"fmt"
	"study/patterns/structural/composite/pkg"
)

var (
	cpu = pkg.Cpu{
		Name:        "CP-1",
		Description: "Процессор-1",
	}
	cpu2 = pkg.Cpu{
		Name:        "CP-2",
		Description: "Процессор-2",
	}
	card = pkg.GraphicsCard{
		Name:        "Radeon-1",
		Description: "Видеокарта-1",
	}
	card2 = pkg.GraphicsCard{
		Name:        "Radeon-2",
		Description: "Видеокарта-2",
	}
	motherboard = pkg.Motherboard{
		Name:        "Gigabyte",
		Description: "Материнская плата",
		Components: []pkg.Component{
			cpu, cpu2, card, card2,
		},
	}
	pc = pkg.Pc{
		Name:        "PC",
		Description: "Personal comp",
		Components: []pkg.Component{
			motherboard,
		},
	}
)

func main() {
	pc.Search("Gigabyte")
	fmt.Println()
	pc.Search("CP-2")
}
