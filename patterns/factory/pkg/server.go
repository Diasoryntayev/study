package pkg

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (pc Server) GetType() string {
	return pc.Type
}

func (pc Server) PrindDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d]", pc.Type, pc.Core, pc.Memory)
}
