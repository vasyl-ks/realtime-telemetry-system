import React from "react";
import SensorChart from "./components/SensorChart";
import "./App.css";

/**
 * Main application component.
 *
 * Renders the header and the main SensorChart component.
 * This is the root component for the Realtime Telemetry System frontend.
 *
 * @component
 * @returns {JSX.Element} The rendered app container.
 */
function App() {
    return (
        <div className="app-container">
            <h1>Realtime Telemetry System</h1>
            {/* Main sensor chart displaying live telemetry */}
            <SensorChart />
        </div>
    );
}

export default App;
