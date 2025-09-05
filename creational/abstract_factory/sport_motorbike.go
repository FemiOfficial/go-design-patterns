package abstractfactory

type SportMotorbike struct {}

func (s *SportMotorbike) GetMotorbikeType() string {
	return "Sport Motorbike"
}	

func (s *SportMotorbike) NumberOfWheels() int {
	return 2
}

func (s *SportMotorbike) NumberOfSeats() int {
	return 1
}