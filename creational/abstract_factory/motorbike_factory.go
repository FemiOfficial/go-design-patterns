package abstractfactory

import "fmt"

const (
	SportMotorbikeType = "sport_motorbike"
	FamilyMotorbikeType = "family_motorbike"
)

type MotorbikeFactory struct {
	// VehicleFactory
}

func (m *MotorbikeFactory) Build(v string) (Vehicle, error) {
	switch v {
		case SportMotorbikeType:
			return new(SportMotorbike), nil
		case FamilyMotorbikeType:
			return new(FamilyMotorbike), nil
		default:
			return nil, fmt.Errorf("invalid motorbike type")
	}
}