package builder

type ManufacturingDirector struct {
	builder BuildProcess
}

func (d *ManufacturingDirector) Construct() {
	// Implementation goes here
	d.builder.SetSeats().SetWheels().SetStructure()
}

func (d *ManufacturingDirector) SetBuilder(b BuildProcess) {
	// Implementation goes here
	d.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}
