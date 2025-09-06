package config

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite" // SQLite driver
)

// DB is the global database handle accessible throughout the project
var DB *sql.DB

// SetupDB initializes the SQLite database and creates the sensor_data table if it doesn't exist.
// It ensures the ./data directory exists and opens a persistent connection to the database.
func SetupDB() {
	// Ensure the data directory exists
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		if err := os.Mkdir("./data", 0755); err != nil {

			log.Fatalf("Failed to create data directory: %v", err)
		}
	}

	// Open SQLite database connection
	var err error
	DB, err = sql.Open("sqlite", "./data/sensors.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Ensure the database connection is alive
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create sensor_data table if it doesn't exist
	const createTableSQL = `
	CREATE TABLE IF NOT EXISTS sensor_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME,
		speed REAL,
		temperature REAL,
		pressure REAL
	);`

	if _, err := DB.Exec(createTableSQL); err != nil {
		log.Fatalf("Failed to create sensor_data table: %v", err)
	}

	log.Println("Database initialized successfully")
}
