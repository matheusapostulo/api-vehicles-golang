package repository

import (
	"app/internal"
	"errors"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

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

func (r *VehicleMap) FindOne(id int) (v internal.Vehicle, err error) {
	v, ok := r.db[id]
	if !ok {
		return internal.Vehicle{}, errors.New("not found")
	}

	return v, nil
}

func (r *VehicleMap) Create(v internal.Vehicle) (err error) {
	_, ok := r.db[v.Id]
	if ok {
		return errors.New("identificador do veículo já existente")
	}
	r.db[v.Id] = v
	return nil
}

func (r *VehicleMap) Update(id int, v internal.Vehicle) (err error) {
	r.db[id] = v
	return nil
}

func (r *VehicleMap) Delete(id int) (err error) {
	_, ok := r.db[id]
	if !ok {
		return errors.New("not found")
	}

	delete(r.db, id)
	return
}
