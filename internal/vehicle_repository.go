package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	FindOne(id int) (v Vehicle, err error)
	Create(v Vehicle) (err error)
	Update(id int, v Vehicle) (err error)
	Delete(id int) (err error)
}
