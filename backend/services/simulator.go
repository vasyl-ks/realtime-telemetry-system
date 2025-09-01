package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"Realtime-Telemetry-System/config"
	"Realtime-Telemetry-System/models"

	"github.com/gorilla/websocket"
)

var (
	// BaseSpeed is the global vehicle speed, protected by SpeedMutex
	BaseSpeed  = 100.0
	SpeedMutex = sync.RWMutex{}

	// DBMutex protects database operations
	DBMutex = sync.Mutex{}

	// Clients holds all active WebSocket connections
	Clients = make(map[*websocket.Conn]bool)

	// ClientsMutex protects the Clients map
	ClientsMutex = &sync.RWMutex{}

	// SensorChannel broadcasts sensor data to all subscribers
	SensorChannel = make(chan models.SensorData, 200)
)

// SimulateSensors generates fake sensor readings every second and stores them in the DB.
// Each reading slightly varies based on the current BaseSpeed.
func SimulateSensors() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		SpeedMutex.RLock()
		currentSpeed := BaseSpeed
		SpeedMutex.RUnlock()

		// generate a new sensor reading
		data := models.SensorData{
			Timestamp:   time.Now(),
			Speed:       currentSpeed,
			Temperature: 19.45 + (currentSpeed / 100.0) + rand.Float64()*0.05,
			Pressure:    99.45 + (currentSpeed / 100.0) + rand.Float64()*0.05,
		}

		// insert into the database
		DBMutex.Lock()
		_, err := config.DB.Exec(
			"INSERT INTO sensor_data(timestamp, speed, temperature, pressure) VALUES (?, ?, ?, ?)",
			data.Timestamp, data.Speed, data.Temperature, data.Pressure,
		)
		DBMutex.Unlock()
		if err != nil {
			fmt.Println("DB insert error:", err)
		}

		// broadcast to WebSocket clients
		SensorChannel <- data
	}
}

// BroadcastData sends new sensor readings to all connected WebSocket clients.
// If a client connection fails, it is removed from the Clients map.
func BroadcastData() {
	for data := range SensorChannel {
		ClientsMutex.RLock()
		for conn := range Clients {
			if err := conn.WriteJSON(data); err != nil {
				// remove client safely on error
				ClientsMutex.RUnlock()
				ClientsMutex.Lock()
				_ = conn.Close()
				delete(Clients, conn)
				ClientsMutex.Unlock()
				ClientsMutex.RLock()
			}
		}
		ClientsMutex.RUnlock()

		// log the reading to the console
		fmt.Printf("[%s] Speed: %.2f km/h | Temp: %.2f°C | Pressure: %.2f Pa\n",
			data.Timestamp.Format("15:04:05"), data.Speed, data.Temperature, data.Pressure)
	}
}

// CleanupOldEntries keeps only the last 10 sensor entries in the database.
// Runs periodically every 100 seconds.
func CleanupOldEntries() {
	ticker := time.NewTicker(100 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		DBMutex.Lock()
		_, err := config.DB.Exec(`
			DELETE FROM sensor_data
			WHERE id <= (
				SELECT id FROM sensor_data
				ORDER BY timestamp DESC
				LIMIT 1 OFFSET 9
			)
		`)
		DBMutex.Unlock()
		if err != nil {
			fmt.Println("Cleanup error:", err)
		}
	}
}
