package abstractfactory

import "testing"

func TestLuxuryCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	car, err := carFactory.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	if car.NumberOfWheels() != 4 {
		t.Errorf("Number of doors is not 4")
	}

	if car.NumberOfSeats() != 4 {
		t.Errorf("Number of wheels is not 4")
	}

	if car, ok := car.(Car); ok {
		if car.NumberOfDoors() != 4 {
			t.Errorf("Number of doors is not 4")
		}
	}
}	

func TestFamilyCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	car, err := carFactory.Build(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}
	
	if car.NumberOfWheels() != 4 {
		t.Errorf("Number of wheels is not 4")
	}
	
	if car.NumberOfSeats() != 5 {
		t.Errorf("Number of seats is not 5")
	}

	if car, ok := car.(Car); ok {
		if car.NumberOfDoors() != 5 {
			t.Errorf("Number of doors is not 5")
		}
	}
}

func TestSportMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbike, err := motorbikeFactory.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	if motorbike.NumberOfWheels() != 2 {
		t.Errorf("Number of wheels is not 2")
	}

	if motorbike.NumberOfSeats() != 1 {
		t.Errorf("Number of seats is not 1")
	}

	if motorbike, ok := motorbike.(Motorbike); ok {
		if motorbike.GetMotorbikeType() != "Sport Motorbike" {
			t.Errorf("Motorbike type is not Sport Motorbike")
		}
	}
}

func TestFamilyMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	
	motorbike, err := motorbikeFactory.Build(FamilyMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	
	if motorbike.NumberOfWheels() != 2 {
		t.Errorf("Number of wheels is not 2")
	}
	
	if motorbike.NumberOfSeats() != 5 {
		t.Errorf("Number of seats is not 5")
	}

	if motorbike, ok := motorbike.(Motorbike); ok {
		if motorbike.GetMotorbikeType() != "Family Motorbike" {
			t.Errorf("Motorbike type is not Family Motorbike")
		}
	}
}