package pkg

type HpFactory struct{}

func (factory HpFactory) GetComputer() Computer {
	return HpComputer{
		Memory: 10,
		CPU:    6,
	}
}

func (factory HpFactory) GetMonitor() Monitor {
	return HpMonitor{
		Size: 40,
	}
}
