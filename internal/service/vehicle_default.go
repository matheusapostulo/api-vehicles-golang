package service

import (
	"app/internal"
	"errors"
	"strconv"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

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

func (s *VehicleDefault) Create(v internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleDefault) GetVehiclesByColorYear(color, year string) (v map[int]internal.Vehicle, err error) {
	vehicles, err := s.rp.FindAll()
	if err != nil {
		return nil, err
	}

	filteredVehicles := make(map[int]internal.Vehicle)
	for key, value := range vehicles {
		intData, _ := strconv.Atoi(year)
		if value.VehicleAttributes.Color == color && value.VehicleAttributes.FabricationYear == intData {
			filteredVehicles[key] = value
		}
	}

	if len(filteredVehicles) == 0 {
		return nil, errors.New("not found")
	}

	return filteredVehicles, nil
}

func (s *VehicleDefault) GetVehiclesByBrandYears(brand, startYear, endYear string) (v map[int]internal.Vehicle, err error) {
	allVehicles, _ := s.rp.FindAll()

	filteredVehicles := make(map[int]internal.Vehicle)

	for key, value := range allVehicles {
		startYearInt, _ := strconv.Atoi(startYear)
		endYearInt, _ := strconv.Atoi(endYear)

		if value.Brand == brand && value.FabricationYear >= startYearInt && value.FabricationYear <= endYearInt {
			filteredVehicles[key] = value
		}
	}

	if len(filteredVehicles) == 0 {
		return nil, errors.New("not found")
	}

	return filteredVehicles, nil
}

func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (speed float64, err error) {
	allVehicles, _ := s.rp.FindAll()

	var totalSpeed float64
	var qtd int

	for _, value := range allVehicles {
		if value.Brand == brand {
			totalSpeed += value.MaxSpeed
			qtd += 1
		}
	}

	if qtd == 0 {
		return 0.0, errors.New("not found")
	}

	return totalSpeed / float64(qtd), nil
}

func (s *VehicleDefault) CreateVehicles(vehicles []internal.Vehicle) (err error) {
	for _, v := range vehicles {
		err = s.rp.Create(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *VehicleDefault) UpdateVehicleSpeed(id int, newSpeed float64) (err error) {
	v, err := s.rp.FindOne(id)
	if err != nil {
		return err
	}

	v.MaxSpeed = newSpeed

	err = s.rp.Update(id, v)
	return
}

func (s *VehicleDefault) GetVehicleByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	allVehicles, err := s.FindAll()
	if err != nil {
		return nil, err
	}

	filteredVehicles := make(map[int]internal.Vehicle)

	for key, value := range allVehicles {
		println(fuelType)
		if value.FuelType == fuelType {
			filteredVehicles[key] = value
		}
	}

	if len(filteredVehicles) == 0 {
		return nil, errors.New("not found")
	}

	return filteredVehicles, nil
}

func (s *VehicleDefault) DeleteVehicle(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *VehicleDefault) GetByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {
	vehicles, _ := s.FindAll()

	filteredVehicles := make(map[int]internal.Vehicle)

	for key, value := range vehicles {
		if value.Transmission == transmissionType {
			filteredVehicles[key] = value
		}
	}

	if len(filteredVehicles) == 0 {
		return nil, errors.New("not found")
	}

	return filteredVehicles, nil
}

func (s *VehicleDefault) UpdateFuelType(id int, fuelType string) (err error) {
	v, err := s.rp.FindOne(id)
	if err != nil {
		return err
	}

	v.FuelType = fuelType
	s.rp.Update(id, v)

	return
}

func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (average int, err error) {
	allVehicles, _ := s.FindAll()

	var totalCapacity int
	var qtd int

	for _, value := range allVehicles {
		if value.Brand == brand {
			qtd += 1
			totalCapacity += value.Capacity
		}
	}

	if qtd == 0 {
		return 0, errors.New("not found")
	}

	return totalCapacity / qtd, nil
}

func (s *VehicleDefault) GetByDimensions(minLengthFloat, maxLengthFloat, minWidthFloat, maxWidthFloat float64) (v map[int]internal.Vehicle, err error) {
	allVehicles, _ := s.FindAll()

	filteredVehicles := make(map[int]internal.Vehicle)

	for key, value := range allVehicles {
		if value.Height >= minLengthFloat && value.Height <= maxLengthFloat && value.Width >= minWidthFloat && value.Width <= maxWidthFloat {
			filteredVehicles[key] = value
		}
	}

	if len(filteredVehicles) == 0 {
		return nil, errors.New("not found")
	}

	return filteredVehicles, nil
}

func (s *VehicleDefault) GetByWeight(min, max float64) (v map[int]internal.Vehicle, err error) {
	allVehicles, _ := s.rp.FindAll()

	filteredVehicles := make(map[int]internal.Vehicle)

	for key, value := range allVehicles {
		if value.Weight >= min && value.Weight <= max {
			filteredVehicles[key] = value
		}
	}

	if len(filteredVehicles) == 0 {
		return nil, errors.New("not found")
	}

	return filteredVehicles, nil
}
