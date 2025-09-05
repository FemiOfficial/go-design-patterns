package abstractfactory

type FamilyMotorbike struct {}

func (f *FamilyMotorbike) GetMotorbikeType() string {
	return "Family Motorbike"
}	

func (f *FamilyMotorbike) NumberOfWheels() int {
	return 2
}

func (f *FamilyMotorbike) NumberOfSeats() int {
	return 5
}