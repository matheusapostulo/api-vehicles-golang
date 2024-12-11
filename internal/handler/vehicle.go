package handler

import (
	"app/internal"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Endpoint 1 - D4
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input VehicleJSON

		err := request.JSON(r, &input)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		vehicle := internal.Vehicle{
			Id: input.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           input.Brand,
				Model:           input.Model,
				Registration:    input.Registration,
				Color:           input.Color,
				FabricationYear: input.FabricationYear,
				Capacity:        input.Capacity,
				MaxSpeed:        input.MaxSpeed,
				FuelType:        input.FuelType,
				Transmission:    input.Transmission,
				Weight:          input.Weight,
				Dimensions: internal.Dimensions{
					Height: input.Height,
					Length: input.Length,
					Width:  input.Width,
				},
			},
		}

		err = h.sv.Create(vehicle)
		// Sobrar tempo instanciar error e comparar com Is
		if err != nil {
			response.JSON(w, http.StatusConflict, nil)
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    vehicle,
		})
	}
}

// Endpoint 2 - D2
func (h *VehicleDefault) GetVehiclesByColorYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")

		vehicles, err := h.sv.GetVehiclesByColorYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Endpoint - D3
func (h *VehicleDefault) GetVehiclesByBrandYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		startYear := chi.URLParam(r, "start_year")
		endYear := chi.URLParam(r, "end_year")

		vehicles, err := h.sv.GetVehiclesByBrandYears(brand, startYear, endYear)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Endpoint 4 - D3
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		averageSpeed, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageSpeed,
		})
	}
}

// Endpoint 5 -> D5
func (h *VehicleDefault) CreateVehicles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inputVehicles []VehicleJSON

		err := request.JSON(r, &inputVehicles)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		var vehiclesConvertedVehicle []internal.Vehicle

		for _, value := range inputVehicles {
			v := internal.Vehicle{
				Id: value.ID,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           value.Brand,
					Model:           value.Model,
					Registration:    value.Registration,
					Color:           value.Color,
					FabricationYear: value.FabricationYear,
					Capacity:        value.Capacity,
					MaxSpeed:        value.MaxSpeed,
					FuelType:        value.FuelType,
					Transmission:    value.Transmission,
					Weight:          value.Height,
					Dimensions: internal.Dimensions{
						Height: value.Height,
						Length: value.Length,
						Width:  value.Width,
					},
				},
			}
			vehiclesConvertedVehicle = append(vehiclesConvertedVehicle, v)
		}

		err = h.sv.CreateVehicles(vehiclesConvertedVehicle)
		if err != nil {
			response.JSON(w, http.StatusConflict, nil)
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Veiculos criados com sucesso",
			"data":    inputVehicles,
		})
	}
}

// Endpoint 6 -> D5
type RequestUpdateSpeed struct {
	NewSpeed float64 `json:"max_speed"`
}

func (h *VehicleDefault) UpdateVehicleSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		var input RequestUpdateSpeed
		err = request.JSON(r, &input)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		err = h.sv.UpdateVehicleSpeed(idInt, input.NewSpeed)

		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		response.JSON(w, http.StatusOK, nil)
	}
}

// Endpoint 7 -> D2
func (h *VehicleDefault) GetVehicleByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuelType := chi.URLParam(r, "type")
		vehicles, err := h.sv.GetVehicleByFuelType(fuelType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Endpoint 8 -> D1
func (h *VehicleDefault) DeleteVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, _ := strconv.Atoi(id)

		err := h.sv.DeleteVehicle(idInt)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

// Endpoint 9 -> D1
func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmissionType := chi.URLParam(r, "type")

		vehicles, err := h.sv.GetByTransmissionType(transmissionType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Endpoint 10 -> D5
type RequestUpdateFuelType struct {
	FuelType string `json:"fuel_type"`
}

func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, _ := strconv.Atoi(id)

		var input RequestUpdateFuelType
		err := request.JSON(r, &input)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		err = h.sv.UpdateFuelType(idInt, input.FuelType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		response.JSON(w, http.StatusOK, nil)

	}
}

// Endpoint 11 -> D3
func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		averageCapacity, err := h.sv.GetAverageCapacityByBrand(brand)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageCapacity,
		})

	}
}

// Endpoint 12 -> D5
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		length := r.URL.Query().Get("length")
		width := r.URL.Query().Get("width")

		lenghts := strings.Split(length, "-")
		widths := strings.Split(width, "-")

		var badFormattedData = "bad formatted data"
		minLengthFloat, err := strconv.ParseFloat(lenghts[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, badFormattedData)
			return
		}
		maxLengthFloat, err := strconv.ParseFloat(lenghts[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, badFormattedData)
			return
		}
		minWidthFloat, err := strconv.ParseFloat(widths[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, badFormattedData)
			return
		}
		maxWidthFloat, err := strconv.ParseFloat(widths[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, badFormattedData)
			return
		}

		vehicles, err := h.sv.GetByDimensions(minLengthFloat, maxLengthFloat, minWidthFloat, maxWidthFloat)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Endpoint 13 -> D3
func (h *VehicleDefault) GetByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minWeigth := r.URL.Query().Get("min")
		maxWeigth := r.URL.Query().Get("max")

		minWeigthFloat, err := strconv.ParseFloat(minWeigth, 64)
		if err != nil {
			return
		}
		maxWeigthFloat, err := strconv.ParseFloat(maxWeigth, 64)
		if err != nil {
			return
		}

		vehicles, err := h.sv.GetByWeight(minWeigthFloat, maxWeigthFloat)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}
