package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufactoringComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufactoringComplex.SetBuilder(carBuilder)
	manufactoringComplex.Construct()

	// Car testing
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	// Bike testing
	bikeBuilder := &BikeBuilder{}
	manufactoringComplex.SetBuilder(bikeBuilder)
	manufactoringComplex.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Seats != 2 {
		t.Errorf("Seats on a bike must be 2 and they were %d\n", bike.Seats)
	}

	if bike.Structure != "Motorbike" {
		t.Errorf("Structure on a bike must be 'Bike' and was %s\n", bike.Structure)
	}
}
