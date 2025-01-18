package builder

import "testing"

func TestBuilder(t *testing.T) {
	var builder Builder
	//招工，建别墅。
	builder = NewHouseBuilder()
	//交给工程总监
	director := &Director{Builder: builder}
	building := NewBuilding()
	building = director.Direct()
	println(building.ToString())
	println()
	//替换施工方，建公寓。
	director.SetBuilder(NewApartmentBuilder())
	building = director.Direct()
	println(building.ToString())
}
