package types

type Epson struct {
	Name string
}

func (s Epson) ScanFile() {
	println("Epson scan file")
}
