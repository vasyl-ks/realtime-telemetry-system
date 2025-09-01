import React, { useEffect, useState, useRef } from "react";
import { Line } from "react-chartjs-2";
import { createWebSocket } from "../services/websocket";
import SensorMetrics from "./SensorMetrics";
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from "chart.js";

// Register chart.js components globally
ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

/**
 * SensorChart component
 * Fetches historical sensor data and subscribes to live updates via WebSocket.
 * Displays velocity, temperature, and pressure charts along with metrics.
 */
export default function SensorChart() {
    const [dataPoints, setDataPoints] = useState([]); // sensor data history
    const historyLimit = 10; // keep only last N data points
    const wsRef = useRef(null); // WebSocket reference

    useEffect(() => {
        // Fetch historical sensor data on mount
        fetch("http://localhost:8080/api/sensors/history")
            .then(res => res.json())
            .then(historyData => setDataPoints(historyData.reverse())) // oldest -> newest
            .catch(err => console.error("Failed to fetch history:", err));

        // Open WebSocket for live updates
        const wsInterface = createWebSocket((data) => {
            // Keep only last `historyLimit` points
            setDataPoints(prev => [...prev.slice(-historyLimit), data]);
        });

        wsRef.current = wsInterface;

        // Cleanup on unmount
        return () => wsInterface.ws.close();
    }, []);

    /**
     * Handles speed adjustment
     * @param {number} delta - positive or negative change
     */
    const handleSpeedChange = (delta) => {
        wsRef.current?.sendSpeedDelta(delta);
    };

    // Chart labels (timestamps)
    const labels = dataPoints.map(d => new Date(d.timestamp).toLocaleTimeString());

    // Prepare datasets for charts
    const velocityData = {
        labels,
        datasets: [{
            label: "Velocity (km/h)",
            data: dataPoints.map(d => d.speed ?? 0),
            borderColor: "red",
            fill: false
        }],
    };
    const temperatureData = {
        labels,
        datasets: [{
            label: "Temperature (°C)",
            data: dataPoints.map(d => d.temperature ?? 0),
            borderColor: "blue",
            fill: false
        }],
    };
    const pressureData = {
        labels,
        datasets: [{
            label: "Pressure (Pa)",
            data: dataPoints.map(d => d.pressure ?? 0),
            borderColor: "green",
            fill: false
        }],
    };

    const options = {
        responsive: true,
        plugins: { legend: { position: "top" } },
    };

    return (
        <div>
            {/* Show metrics and speed controls */}
            <SensorMetrics
                data={dataPoints[dataPoints.length - 1]}
                onSpeedChange={handleSpeedChange}
            />

            {/* Render charts */}
            <Line data={velocityData} options={options} />
            <Line data={temperatureData} options={options} />
            <Line data={pressureData} options={options} />
        </div>
    );
}
