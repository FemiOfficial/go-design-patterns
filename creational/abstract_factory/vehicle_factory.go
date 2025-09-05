package abstractfactory

import (
	"fmt"
)

type VehicleFactory interface {
	Build(v string) (Vehicle, error)
}	

const (
	CarFactoryType = "car"
	MotorbikeFactoryType = "motorbike"
)
func  BuildFactory(f string) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("invalid vehicle factory type")
	}
}