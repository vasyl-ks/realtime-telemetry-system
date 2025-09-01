package handlers

import (
	"fmt"
	"log"
	"net/http"

	"Realtime-Telemetry-System/config"
	"Realtime-Telemetry-System/models"
	"Realtime-Telemetry-System/services"

	"github.com/gorilla/websocket"
)

// Upgrader is used to upgrade HTTP connections to WebSocket connections.
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// References to speed control variables in services.
var SpeedMutex = &services.SpeedMutex
var BaseSpeedPtr = &services.BaseSpeed

// WSHandler upgrades the HTTP connection to a WebSocket, manages the client,
// sends the last 10 sensor readings, and listens for incoming messages.
func WSHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	// Register client
	services.ClientsMutex.Lock()
	services.Clients[conn] = true
	services.ClientsMutex.Unlock()

	// Send the last 10 sensor entries to the newly connected client
	sendLastEntries(conn)

	// Listen for incoming messages from the client
	for {
		var msg map[string]interface{}
		if err := conn.ReadJSON(&msg); err != nil {
			// Remove client on error/disconnect
			services.ClientsMutex.Lock()
			delete(services.Clients, conn)
			services.ClientsMutex.Unlock()
			_ = conn.Close()
			break
		}

		// Handle speed adjustment messages
		if msg["type"] == "speedDelta" {
			if val, ok := msg["value"].(float64); ok {
				adjustBaseSpeed(val)
			}
		}
	}
}

// sendLastEntries retrieves the last 10 sensor readings and sends them to the client
func sendLastEntries(conn *websocket.Conn) {
	DBMutex.Lock()
	rows, err := config.DB.Query(`
		SELECT timestamp, speed, temperature, pressure
		FROM sensor_data
		ORDER BY timestamp DESC
		LIMIT 10`)
	DBMutex.Unlock()

	if err != nil {
		log.Printf("DB query error while sending last entries: %v", err)
		return
	}
	defer rows.Close()

	var history []models.SensorData
	for rows.Next() {
		var data models.SensorData
		if err := rows.Scan(&data.Timestamp, &data.Speed, &data.Temperature, &data.Pressure); err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		history = append(history, data)
	}

	// Send in chronological order (oldest first)
	for i := len(history) - 1; i >= 0; i-- {
		if err := conn.WriteJSON(history[i]); err != nil {
			log.Printf("Failed to send JSON to client: %v", err)
		}
	}
}

// adjustBaseSpeed safely updates the global base speed with limits
func adjustBaseSpeed(delta float64) {
	SpeedMutex.Lock()
	defer SpeedMutex.Unlock()

	*BaseSpeedPtr += delta
	if *BaseSpeedPtr > 125 {
		*BaseSpeedPtr = 125
	} else if *BaseSpeedPtr < 75 {
		*BaseSpeedPtr = 75
	}

	fmt.Printf("Speed adjusted by %.0f, new baseSpeed=%.2f\n", delta, *BaseSpeedPtr)
}
