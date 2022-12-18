package types

import "study/patterns/structural/bridge/pkg/base"

type MacPc struct {
	scanner base.Scanner
}

func (pc *MacPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}

func (pc *MacPc) Scan() {
	println("Scan pc Mac system")
	pc.scanner.ScanFile()
}
