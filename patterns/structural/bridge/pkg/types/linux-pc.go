package types

import "study/patterns/structural/bridge/pkg/base"

type LinuxPc struct {
	scanner base.Scanner
}

func (pc *LinuxPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *LinuxPc) Scan() {
	println("Scan pc linux system")
	pc.scanner.ScanFile()
}
