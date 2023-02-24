package builder

type CarBuilder struct {
	product VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.product.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.product.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.product.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.product
}
