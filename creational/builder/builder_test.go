package creational

import "testing"

func TestBuilderPattern(tt * testing.T) {
	director := ManufacturingDirector{}
	carBuilder := CarBuilder{}
	director.SetBuilder(&carBuilder)
	car := director.Construct()
	if car.Wheels != 4 {
		tt.Errorf("Wheels: %d; want 4", car.Wheels)
	}
	if car.Seats != 5 {
		tt.Errorf("Seats: %d; want 5", car.Seats)
	}
	if car.Structure != "Car" {
		tt.Errorf("Structure: %s; want Car", car.Structure)
	}

	busBuilder := BusBuilder{}
	director.SetBuilder(&busBuilder)
	bus := director.Construct()

	if bus.Wheels != 8 {
		tt.Errorf("Wheels: %d; want 8", bus.Wheels)
	}
	if bus.Seats != 30 {
		tt.Errorf("Seats: %d; want 30", bus.Seats)
	}
	if bus.Structure != "Bus" {
		tt.Errorf("Structure: %s; want Bus", bus.Structure)
	}
}
