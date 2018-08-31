package builder

import (
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	basicBuilding := building{
		Roof:    false,
		Garage:  false,
		Floors:  1,
		Windows: 4,
	}

	build := NewBuild()
	BasicBuilder := NewBasicBuilder()
	build.SetBuilder(BasicBuilder)

	building := build.builder.SetRoof().SetGarage().SetFloors().SetWindows().GetBuilding()

	if !reflect.DeepEqual(building, basicBuilding) {
		t.Errorf("Basic Biilding is not working correctly")
	}
}
