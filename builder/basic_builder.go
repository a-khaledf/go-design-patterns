package builder

type BasicBuilder struct {
	building building
}

func NewBasicBuilder() *BasicBuilder {
	BasicBuilder := BasicBuilder{
		building: NewBuilding(),
	}

	return &BasicBuilder
}

func (b *BasicBuilder) SetRoof() Builder {
	b.building.Roof = false
	return b
}

func (b *BasicBuilder) SetGarage() Builder {
	b.building.Garage = false
	return b
}

func (b *BasicBuilder) SetFloors() Builder {
	b.building.Floors = 1

	return b
}

func (b *BasicBuilder) SetWindows() Builder {
	b.building.Windows = 4

	return b
}

func (b *BasicBuilder) GetBuilding() building {
	return b.building
}
