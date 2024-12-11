package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(v Vehicle) (err error)
	GetVehiclesByColorYear(color, year string) (v map[int]Vehicle, err error)
	GetVehiclesByBrandYears(brand, startYear, endYear string) (v map[int]Vehicle, err error)
	GetAverageSpeedByBrand(brand string) (speed float64, err error)
	CreateVehicles(vehicles []Vehicle) (err error)
	UpdateVehicleSpeed(id int, newSpeed float64) (err error)
	GetVehicleByFuelType(fuelType string) (v map[int]Vehicle, err error)
	DeleteVehicle(id int) (err error)
	GetByTransmissionType(transmissionType string) (v map[int]Vehicle, err error)
	UpdateFuelType(id int, fuelType string) (err error)
	GetAverageCapacityByBrand(brand string) (averageCapacity int, err error)
	GetByDimensions(minLengthFloat, maxLengthFloat, minWidthFloat, maxWidthFloat float64) (v map[int]Vehicle, err error)
	GetByWeight(minWeigthFloat, maxWeigthFloat float64) (v map[int]Vehicle, err error)
}
