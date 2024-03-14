package handler

import (
	"app/internal"
	"app/tools"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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

// Constructor
// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// Inyection of the service
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

// GetByColorYear is a method that returns a handler for the route GET /vehicles/color/:color/year/:year
func (h *VehicleDefault) GetByColorYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get color from path
		colorStr := chi.URLParam(r, "color")

		// - get year from path
		yearStr := chi.URLParam(r, "year")
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "invalid year", http.StatusBadRequest)
			return
		}

		vehicles, err := h.sv.FindByColorYear(colorStr, year)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		if len(vehicles) == 0 {
			http.Error(w, "No se encontraron vehiculos con esos criterios", http.StatusNotFound)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}

// GetByBrandYearRange is a method that returns a handler for the route GET /vehicles/brand/:brand/year/:year
func (h *VehicleDefault) GetByBrandYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get brand from path
		brandStr := chi.URLParam(r, "brand")

		// - get start year from path
		startYearStr := chi.URLParam(r, "startYear")
		startYear, err := strconv.Atoi(startYearStr)
		if err != nil {
			http.Error(w, "invalid start year", http.StatusBadRequest)
			return
		}

		// - get end year from path
		endYearStr := chi.URLParam(r, "endYear")
		endYear, err := strconv.Atoi(endYearStr)
		if err != nil {
			http.Error(w, "invalid end year", http.StatusBadRequest)
			return
		}

		// PROCESS
		// - calling the service
		vehicles, err := h.sv.FindByBrandYearRange(brandStr, startYear, endYear)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		// - if no vehicles found
		if len(vehicles) == 0 {
			http.Error(w, "No se encontraron vehiculos con esos criterios", http.StatusNotFound)
			return
		}

		// RESPONSE
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}

// GetByBrandAverageSpeed is a method that returns a handler for the route GET /vehicles/brand/:brand/average-speed
func (h *VehicleDefault) GetByBrandAverageSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get brand from path
		brandStr := chi.URLParam(r, "brand")

		// PROCESS
		// - calling the service
		averageSpeed, err := h.sv.FindByBrandAverageSpeed(brandStr)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		// - if no vehicles found
		if averageSpeed == 0 {
			http.Error(w, "No se encontraron vehiculos de esa marca", http.StatusNotFound)
			return
		}

		// RESPONSE
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageSpeed,
		})
	}
}

// GetByFuelType is a method that returns a handler for the route GET /vehicles/fuel-type/:fuelType
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get fuel type from path
		fuelTypeStr := chi.URLParam(r, "type")

		// PROCESS
		// - calling the service
		vehicles, err := h.sv.FindByFuelType(fuelTypeStr)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		// - if no vehicles found
		if len(vehicles) == 0 {
			http.Error(w, "No se encontraron vehiculos con ese tipo de combustible", http.StatusNotFound)
			return
		}

		// RESPONSE
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}

// GetByTransmission is a method that returns a handler for the route
func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get transmission from path
		transmissionStr := chi.URLParam(r, "type")

		// PROCESS
		// - calling the service
		vehicles, err := h.sv.FindByTransmissionType(transmissionStr)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		// - if no vehicles found
		if len(vehicles) == 0 {
			http.Error(w, "No se encontraron vehiculos con ese tipo de transmision", http.StatusNotFound)
			return
		}

		// RESPONSE
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}

// GetByPassengers is a method that returns a handler for the route GET /vehicles/passengers/:passengers
func (h *VehicleDefault) GetByBrandAverageCapacity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get passengers from path
		brandStr := chi.URLParam(r, "brand")

		// PROCESS
		// - calling the service
		averageCapacity, err := h.sv.FindByBrandAverageCapacity(brandStr)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		// - if no vehicles found
		if averageCapacity == 0 {
			http.Error(w, "No se encontraron vehiculos de esa marca", http.StatusNotFound)
			return
		}

		// RESPONSE
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageCapacity,
		})
	}
}

func (h *VehicleDefault) GetByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get weight from path
		weightMinStr := r.URL.Query().Get("min")
		weightMaxStr := r.URL.Query().Get("max")

		weightMin, err := strconv.ParseFloat(weightMinStr, 64)
		if err != nil {
			http.Error(w, "invalid min weight", http.StatusBadRequest)
			return
		}
		weightMax, err := strconv.ParseFloat(weightMaxStr, 64)
		if err != nil {
			http.Error(w, "invalid max weight", http.StatusBadRequest)
			return
		}

		// PROCESS
		// - calling the service
		vehicles, err := h.sv.FindByWeightRange(weightMin, weightMax)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		// - if no vehicles found
		if len(vehicles) == 0 {
			http.Error(w, "No se encontraron vehiculos en ese rango de peso", http.StatusNotFound)
			return
		}

		// RESPONSE
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}

// GetByDimensions is a method that returns a handler for the route GET /vehicles/dimensions
func (h *VehicleDefault) GetByDimensionRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - get dimensions from query
		heightStr := r.URL.Query().Get("height")
		heightDimension := strings.Split(heightStr, "-")
		if len(heightDimension) != 2 {
			http.Error(w, "invalid height", http.StatusBadRequest)
			return
		}
		heightMin, err := strconv.ParseFloat(heightDimension[0], 64)
		if err != nil {
			http.Error(w, "invalid height", http.StatusBadRequest)
			return
		}
		heightMax, err := strconv.ParseFloat(heightDimension[1], 64)
		if err != nil {
			http.Error(w, "invalid height", http.StatusBadRequest)
			return
		}

		widthStr := r.URL.Query().Get("width")
		widthDimension := strings.Split(widthStr, "-")
		if len(widthDimension) != 2 {
			http.Error(w, "invalid width", http.StatusBadRequest)
			return
		}
		widthMin, err := strconv.ParseFloat(widthDimension[0], 64)
		if err != nil {
			http.Error(w, "invalid width", http.StatusBadRequest)
			return
		}
		widthMax, err := strconv.ParseFloat(widthDimension[1], 64)
		if err != nil {
			http.Error(w, "invalid width", http.StatusBadRequest)
			return
		}

		// PROCESS
		// - calling the service
		vehicles, err := h.sv.FindByDimensionRange(heightMin, widthMin, heightMax, widthMax)
		if err != nil {
			http.Error(w, "error getting vehicles", http.StatusInternalServerError)
			return
		}

		// - if no vehicles found
		if len(vehicles) == 0 {
			http.Error(w, "No se encontraron vehiculos con esas dimensiones", http.StatusNotFound)
			return
		}

		// RESPONSE
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}

/*
*	SAVE
 */

func (h *VehicleDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// -read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "error reading request"})
			return
		}

		// - parse to map (dynamic)
		bodyMap := make(map[string]any)
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "error parsing request"})
			return
		}

		// validate fields
		// - validate required fields
		if err := tools.CheckFieldExistance(bodyMap,
			"brand", "model", "registration", "color", "year", "passengers", "max_speed", "fuel_type", "transmission", "weight", "height", "length", "width"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": fmt.Sprintf("field %s is required", fieldError.Field)})
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "error validating request"})
			return
		}

		// - parse json to struct
		var body VehicleJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "error parsing request"})
			return
		}

		// Writing to the database
		file, err := os.OpenFile("../docs/db/vehicles_100.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "error opening file"})
			return
		}
		defer file.Close()

		// Serialize to json
		vehicleJSON, err := json.Marshal(body)
		fmt.Println("Serializing vehicle")
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "error serializing vehicle"})
			return
		}
		// Write to file
		if _, err := file.Write(vehicleJSON); err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "error writing to file"})
			return
		} else {
			// Si la escritura en el archivo fue exitosa, registra un mensaje de éxito
			log.Println("Successfully wrote to file")
		}

		// PROCESS
		// serialize to vehicle
		vehicle := internal.Vehicle{
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}
		/*
			resp := h.sv.Save(&vehicle)
			response.JSON(w, http.StatusCreated, map[string]any{
				"message": "success",
				"data":    resp,
			}
		*/
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    vehicle,
		})
	}
}

/*
func (h *VehicleDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// -read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error reading request"})
			return
		}

		// - parse to map (dynamic)
		bodyMap := make(map[string]interface{})
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error parsing request"})
			return
		}

		// validate fields
		// - validate required fields
		if err := tools.CheckFieldExistance(bodyMap,
			"brand", "model", "registration", "color", "year", "passengers", "max_speed", "fuel_type", "transmission", "weight", "height", "length", "width"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": fmt.Sprintf("field %s is required", fieldError.Field)})
			}
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error validating request"})
			return
		}

		// - parse json to struct
		var body VehicleJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error parsing request"})
			return
		}

		// Writing to the file
		file, err := os.OpenFile("../docs/db/vehicles_100.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error opening file"})
			return
		}
		defer file.Close()

		// Serialize to json
		vehicleJSON, err := json.Marshal(body)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error serializing vehicle"})
			return
		}

		// Write to file
		if _, err := file.Write(vehicleJSON); err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error writing to file"})
			return
		}

		// PROCESS
		// serialize to vehicle
		vehicle := internal.Vehicle{
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		// Log successful writing to file
		log.Println("Successfully wrote to file")

		// Respond with success message and data
		response.JSON(w, http.StatusCreated, map[string]interface{}{
			"message": "success",
			"data":    vehicle,
		})
	}
}

*/
