# Realtime Telemetry System

**Realtime Telemetry System** is a full-stack, real-time monitoring application built with Go, React, WebSockets, and SQLite.
It simulates live sensor data — velocity, temperature, and pressure — and visualizes it on an interactive dashboard.
This project demonstrates real-time communication, data persistence, and responsive UI design, all packaged in a Dockerized environment for easy deployment.

---

## Features

* Real-time sensor simulation (velocity, temperature, pressure)
* Live updates via WebSocket
* Historical data storage (last 10 readings) in SQLite
* Interactive speed controls on the frontend
* Dockerized setup for quick and reproducible deployment
* Responsive charts using React + Chart.js
* Unit testing coverage for core backend logic, ensuring reliability of sensor data updates and business rules
* Automatic database cleanup: every 100 seconds the system removes the 90 oldest entries generated during that interval, ensuring the database remains lightweight and efficient over time

---

## Learnings & Skills Gained

* Implemented WebSocket communication for real-time updates
* Built a REST API in Go for historical and latest sensor data
* Managed data persistence with SQLite and structured models
* Developed dynamic charts using React and Chart.js
* Applied component-based design and state management in React
* Containerized full-stack application with Docker and Docker Compose
* Implemented unit tests in Go to validate backend functionality, ensure correctness, and maintain reliability throughout development
* Integrated frontend and backend with robust error handling

---

## Getting Started

### Prerequisites

* Docker
* Go 1.20+
* Node.js 18+ and npm

### Installation

1. Clone the repository:

```bash
git clone https://github.com/vasyl-ks/realtime-telemetry-system
cd realtime-telemetry
```

2. Build and run the application:

```bash
docker-compose up --build
```

3. Access the frontend:
   [http://localhost:3000](http://localhost:3000)

### Backend API Endpoints

* Latest reading: [http://localhost:8080/api/sensors/latest](http://localhost:8080/api/sensors/latest)
* Last 10 readings: [http://localhost:8080/api/sensors/history](http://localhost:8080/api/sensors/history)

---

## Usage

* **Live Dashboard:** Observe velocity, temperature, and pressure in real-time charts
* **Interactive Control:** Adjust simulation speed with +5 or -5 buttons
* **API Access:** Fetch historical or latest readings for testing or integration

---

## Project Structure

```text
Realtime-Telemetry-System
│   docker-compose.yml
│   LICENSE
│   README
├── backend
│   │   Dockerfile
│   │   go.mod
│   │   go.sum
│   │   main.go
│   │   main_test.go
│   ├── config
│   │       db.go
│   ├── data
│   │       sensors.db
│   ├── handlers
│   │       http.go
│   │       ws.go
│   ├── models
│   │       sensor.go
│   └── services
│           simulator.go
└── frontend
    │   Dockerfile
    └── src
        │   App.css
        │   App.jsx
        │   index.css
        │   index.js
        ├── components
        │       SensorChart.jsx
        │       SensorMetrics.jsx
        └── services
                websocket.js
```

---

## Examples

### Adjusting Speed

| Action            | Result                    |
| ----------------- | ------------------------- |
| Click +5 Velocity | Speed increases by 5 km/h |
| Click -5 Velocity | Speed decreases by 5 km/h |

### Viewing Sensor Data

* Velocity, temperature, and pressure charts update live every second
* Historical data shows the last 10 readings

---

## Architecture Overview

The frontend React application communicates with the Go backend through WebSockets for live updates.
The backend also exposes HTTP endpoints to fetch the latest and historical sensor data.
All readings are stored in a SQLite database, which ensures persistent storage of sensor history.

---

## Possible Improvements

Other developers are encouraged to enhance this project. Ideas include:

* Extend historical storage to retain more data with pagination
* Add user authentication and multi-user support
* Include additional sensor types and custom simulation settings
* Enhance the frontend with more interactive charts and visualization options
* Deploy to the cloud with automated CI/CD pipelines for easier accessibility

---

## License

MIT License – free to use, modify, and share.
