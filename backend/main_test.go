package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"Realtime-Telemetry-System/config"
	"Realtime-Telemetry-System/handlers"
	"Realtime-Telemetry-System/services"
)

// main initializes the database, starts background tasks, and serves HTTP endpoints.
func main() {
	// Seed random generator for sensor simulation
	rand.Seed(time.Now().UnixNano())

	// Setup SQLite database and tables
	config.SetupDB()

	// Start background goroutines
	go services.CleanupOldEntries() // periodically clean old sensor data
	go services.SimulateSensors()   // generate sensor readings every second
	go services.BroadcastData()     // broadcast readings to WebSocket clients

	// HTTP REST endpoints
	http.HandleFunc("/api/sensors/latest", handlers.LatestSensorHandler)
	http.HandleFunc("/api/sensors/history", handlers.HistorySensorHandler)

	// WebSocket endpoint
	http.HandleFunc("/ws", handlers.WSHandler)

	fmt.Println("Server running at http://localhost:8080")

	// Start HTTP server (blocking call)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// panic here is reasonable since the server failed to start
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
