package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	// FindByColorYear is a method that returns a map of vehicles by color and year
	FindByColorYear(color string, year int) (v map[int]Vehicle, err error)

	// FindByBrandYearRange is a method that returns a map of vehicles by brand and year range
	FindByBrandYearRange(brand string, startYear, endYear int) (v map[int]Vehicle, err error)

	// FindByBrandAverageSpeed is a method that returns the average speed of vehicles by brand
	FindByBrandAverageSpeed(brand string) (avg float64, err error)

	// FindByFuelType is a method that returns a map of vehicles by fuel type
	FindByFuelType(fuelType string) (v map[int]Vehicle, err error)

	// FindByTransmissionType is a method that returns a map of vehicles by transmission type
	FindByTransmissionType(transmissionType string) (v map[int]Vehicle, err error)

	// FindByBrandAverageCapacity is a method that returns the average capacity of vehicles by brand
	FindByBrandAverageCapacity(brand string) (avg float64, err error)

	// FindByWeightRange is a method that returns a map of vehicles by weight range
	FindByWeightRange(startWeight, endWeight float64) (v map[int]Vehicle, err error)

	// FindByDimensionRange is a method that returns a map of vehicles by dimension range
	FindByDimensionRange(minHeight, minWidth, maxHeight, maxWidth float64) (v map[int]Vehicle, err error)

	// Save is a method that creates a new vehicle
	Save(v *Vehicle) (err error)
}
