package abstractfactory

import (
	"fmt"
)

const (
	LuxuryCarType = "luxury_car"
	FamilyCarType = "family_car"
)

type CarFactory struct {
}

func (c *CarFactory) Build(v string) (Vehicle, error) {
	switch v {
		case LuxuryCarType:
			return new(LuxuryCar), nil
		case FamilyCarType:
			return new(FamilyCar), nil
		default:
			return nil, fmt.Errorf("invalid car type")
	}
}