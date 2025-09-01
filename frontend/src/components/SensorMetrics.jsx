import React from "react";

/**
 * SensorMetrics component
 * Displays the latest sensor readings (speed, temperature, pressure) in styled cards.
 * Provides buttons to adjust speed via callback.
 *
 * @param {Object} props
 * @param {Object} props.data - Latest sensor data
 * @param {number} props.data.speed - Current speed in km/h
 * @param {number} props.data.temperature - Current temperature in °C
 * @param {number} props.data.pressure - Current pressure in kPa
 * @param {Function} props.onSpeedChange - Callback to change speed (+/- delta)
 */
export default function SensorMetrics({ data, onSpeedChange }) {
    // If no data yet, render nothing
    if (!data) return null;

    // Card styling for consistent layout
    const cardStyle = {
        flex: 1,
        padding: "15px",
        margin: "10px",
        borderRadius: "10px",
        backgroundColor: "#f7f7f7",
        textAlign: "center",
        boxShadow: "0 2px 6px rgba(0,0,0,0.1)",
    };

    // Button styling for speed adjustment
    const buttonStyle = {
        margin: "0 5px",
        padding: "5px 10px",
        fontSize: "16px",
        cursor: "pointer",
        borderRadius: "5px",
        border: "1px solid #ccc",
        backgroundColor: "#fff",
    };

    return (
        <div style={{ display: "flex", justifyContent: "space-around", flexWrap: "wrap" }}>
            {/* Velocity card with increment/decrement buttons */}
            <div style={cardStyle}>
                <h3>Velocity</h3>
                <div style={{ display: "flex", alignItems: "center", justifyContent: "center" }}>
                    <button style={buttonStyle} onClick={() => onSpeedChange(-5)}>-5</button>
                    <p style={{ margin: "0 10px" }}>{data.speed.toFixed(2)} km/h</p>
                    <button style={buttonStyle} onClick={() => onSpeedChange(5)}>+5</button>
                </div>
            </div>

            {/* Temperature card */}
            <div style={cardStyle}>
                <h3>Temperature</h3>
                <p>{data.temperature.toFixed(2)} °C</p>
            </div>

            {/* Pressure card */}
            <div style={cardStyle}>
                <h3>Pressure</h3>
                <p>{data.pressure.toFixed(2)} Pa</p>
            </div>
        </div>
    );
}
