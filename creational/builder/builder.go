package creational

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder (b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() VehicleProduct {
	m.builder.SetSeats()
	m.builder.SetWheels()
	m.builder.SetStructure()
	return m.builder.GetVehicle()
}

type CarBuilder struct {
	vehicle VehicleProduct
}

type BusBuilder struct {
	vehicle VehicleProduct
}

func(b *BusBuilder) SetWheels() BuildProcess {
	b.vehicle.Wheels = 8
	return b
}

func(b *BusBuilder) SetSeats() BuildProcess {
	b.vehicle.Seats = 30
	return b
}

func(b *BusBuilder) SetStructure() BuildProcess {
	b.vehicle.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.vehicle
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.vehicle.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.vehicle.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.vehicle.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicle
}
