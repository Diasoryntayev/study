package pkg

import "fmt"

const (
	ServerType       = "server"
	PersonalCompType = "personal"
	LaptopType       = "laptop"
)

type Computer interface {
	GetType() string // возвращает тип ввиде строки
	PrindDetails()   // выводит в консоль детальную инф о объекте
}

func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s Несуществующий тип объекта!\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalCompType:
		return NewPersonalComputer()
	case LaptopType:
		return NewLaptop()
	}
}
