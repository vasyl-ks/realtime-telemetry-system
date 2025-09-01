package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"Realtime-Telemetry-System/config"
	"Realtime-Telemetry-System/models"
)

// DBMutex protects database queries from concurrent access.
var DBMutex sync.Mutex

// LatestSensorHandler returns the most recent sensor reading as JSON.
// Responds with HTTP 204 if no data is available.
func LatestSensorHandler(w http.ResponseWriter, r *http.Request) {
	row := config.DB.QueryRow(
		"SELECT timestamp, speed, temperature, pressure FROM sensor_data ORDER BY timestamp DESC LIMIT 1",
	)

	var data models.SensorData
	if err := row.Scan(&data.Timestamp, &data.Speed, &data.Temperature, &data.Pressure); err != nil {
		http.Error(w, "No data available", http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode latest sensor data: %v", err)
	}
}

// HistorySensorHandler returns the last N sensor readings in chronological order as JSON.
func HistorySensorHandler(w http.ResponseWriter, r *http.Request) {
	const historyLimit = 10

	// Lock DB access for query
	DBMutex.Lock()
	rows, err := config.DB.Query(`
		SELECT timestamp, speed, temperature, pressure
		FROM sensor_data
		ORDER BY timestamp DESC
		LIMIT ?`, historyLimit)
	DBMutex.Unlock()

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Printf("DB query error: %v", err)
		return
	}
	defer rows.Close()

	// Collect rows into a slice
	var history []models.SensorData
	for rows.Next() {
		var data models.SensorData
		if err := rows.Scan(&data.Timestamp, &data.Speed, &data.Temperature, &data.Pressure); err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		history = append(history, data)
	}

	// Reverse to chronological order (oldest -> newest)
	for i, j := 0, len(history)-1; i < j; i, j = i+1, j-1 {
		history[i], history[j] = history[j], history[i]
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(history); err != nil {
		log.Printf("Failed to encode history data: %v", err)
	}
}
