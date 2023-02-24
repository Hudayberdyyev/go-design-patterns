package builder

type BikeBuilder struct {
	product VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.product.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.product.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.product.Structure = "Motorbike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.product
}
