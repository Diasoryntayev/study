package main

import (
	"fmt"
	"study/patterns/structural/bridge/pkg/types"
)

var (
	// scanners
	hpScanner    = types.Hp{}
	epsonScanner = types.Epson{Name: "Epson scanner"}
	// computers
	linuxPc = types.LinuxPc{}
	winPc   = types.WinPc{}
	macPc   = types.MacPc{}
)

func main() {
	linuxPc.AddScanner(hpScanner)
	linuxPc.Scan()

	winPc.AddScanner(hpScanner)
	winPc.Scan()

	macPc.AddScanner(hpScanner)
	macPc.Scan()
	fmt.Println()

	linuxPc.AddScanner(epsonScanner)
	linuxPc.Scan()

	winPc.AddScanner(epsonScanner)
	winPc.Scan()

	macPc.AddScanner(epsonScanner)
	macPc.Scan()
}
