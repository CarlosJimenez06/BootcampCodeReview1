package service

import (
	"app/internal"
)

// Constructor
// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// Inyection of the repository
// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) FindByColorYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorYear(color, year)
	return
}

func (s *VehicleDefault) FindByBrandYearRange(brand string, startYear, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandYearRange(brand, startYear, endYear)
	return
}

func (s *VehicleDefault) FindByBrandAverageSpeed(brand string) (avg float64, err error) {
	avg, err = s.rp.FindByBrandAverageSpeed(brand)
	return
}

func (s *VehicleDefault) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByFuelType(fuelType)
	return
}

func (s *VehicleDefault) FindByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByTransmissionType(transmissionType)
	return
}

func (s *VehicleDefault) FindByBrandAverageCapacity(brand string) (avg float64, err error) {
	avg, err = s.rp.FindByBrandAverageCapacity(brand)
	return
}

func (s *VehicleDefault) FindByWeightRange(minWeight, maxWeight float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByWeightRange(minWeight, maxWeight)
	return
}

func (s *VehicleDefault) FindByDimensionRange(minHeight, minWidth, maxHeight, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByDimensionRange(minHeight, minWidth, maxHeight, maxWidth)
	return
}

// Save is a method that saves a vehicle
func (s *VehicleDefault) Save(v *internal.Vehicle) (err error) {
	err = s.rp.Save(v)
	return
}
