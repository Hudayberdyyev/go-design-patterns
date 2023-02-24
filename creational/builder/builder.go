package builder

type manufacturingDirector struct {
	builder BuildProcess
}

func (d *manufacturingDirector) Construct() {
	// Implementation goes here
	d.builder.SetSeats().SetWheels().SetStructure()
}

func (d *manufacturingDirector) SetBuilder(b BuildProcess) {
	// Implementation goes here
	d.builder = b
}

var directorInstance *manufacturingDirector

func GetDirectorInstance() *manufacturingDirector {
	if directorInstance == nil {
		directorInstance = new(manufacturingDirector)
	}
	return directorInstance
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
