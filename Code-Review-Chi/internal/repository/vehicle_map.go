package repository

import "app/internal"

// Constructor
// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// Inyection of the db
// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) FindByColorYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) FindByBrandYearRange(brand string, startYear, endYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) FindByBrandAverageSpeed(brand string) (avg float64, err error) {
	var totalSpeed float64
	var totalVehicles float64

	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			totalSpeed += value.MaxSpeed
			totalVehicles++
		}
	}
	if totalVehicles == 0 {
		avg = 0
		return
	}

	avg = totalSpeed / totalVehicles

	return
}

func (r *VehicleMap) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.FuelType == fuelType {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Transmission == transmissionType {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindByBrandAverageCapacity(brand string) (avg float64, err error) {
	var totalCapacity float64
	var totalVehicles float64

	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			totalCapacity += float64(value.Capacity)
			totalVehicles++
		}
	}
	if totalVehicles == 0 {
		avg = 0
		return
	}

	avg = totalCapacity / totalVehicles

	return
}

func (r *VehicleMap) FindByWeightRange(minWeight, maxWeight float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Weight >= minWeight && value.Weight <= maxWeight {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindByDimensionRange(minHeight, minWidth, maxHeight, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Height >= minHeight && value.Width >= minWidth && value.Height <= maxHeight && value.Width <= maxWidth {
			v[key] = value
		}
	}
	return
}

// Save is a method that saves a vehicle
func (r *VehicleMap) Save(v *internal.Vehicle) (err error) {
	r.db[v.Id] = *v
	return
}
