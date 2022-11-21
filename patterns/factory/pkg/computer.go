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
		fmt.Printf("%s Несуществнющий тип объекта!\n", typeName)
		return nil
	case ServerType:
		return nil
	case PersonalCompType:
		return nil
	case LaptopType:
		return nil
	}
}
