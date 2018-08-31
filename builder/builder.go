package builder

type building struct {
	Roof    bool
	Garage  bool
	Floors  int
	Windows int
}

type Builder interface {
	SetRoof() Builder
	SetGarage() Builder
	SetFloors() Builder
	SetWindows() Builder
	GetBuilding() building
}

func NewBuilding() building {
	building := building{}

	return building
}

type Build struct {
	builder Builder
}

func NewBuild() *Build {
	build := Build{}

	return &build
}

func (b *Build) SetBuilder(builder Builder) {
	b.builder = builder
}
