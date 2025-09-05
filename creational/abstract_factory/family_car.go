package abstractfactory

type FamilyCar struct {}

func (f * FamilyCar) NumberOfDoors() int {
	return 5
}

func (f * FamilyCar) NumberOfWheels() int {
	return 4
}

func (f * FamilyCar) NumberOfSeats() int {
	return 5
}