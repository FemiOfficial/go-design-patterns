package abstractfactory

type LuxuryCar struct {}

func (l * LuxuryCar) NumberOfDoors() int {
	return 4
}

func (l * LuxuryCar) NumberOfWheels() int {
	return 4
}

	func (l * LuxuryCar) NumberOfSeats() int {
		return 4
}