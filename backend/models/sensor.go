package models

import "time"

// SensorData represents a single sensor reading from the telemetry system.
// It includes the timestamp of the reading, the current speed of the vehicle,
// temperature, and pressure values.
type SensorData struct {
	Timestamp   time.Time `json:"timestamp"`   // The time when the reading was recorded
	Speed       float64   `json:"speed"`       // Current speed in km/h
	Temperature float64   `json:"temperature"` // Temperature in °C
	Pressure    float64   `json:"pressure"`    // Pressure in Pa
}
