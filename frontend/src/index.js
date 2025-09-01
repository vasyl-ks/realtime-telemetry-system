import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";

/**
 * Entry point of the React application.
 *
 * This file mounts the root React component (`App`) into the DOM.
 * It uses React 18's `createRoot` API and wraps the app in `React.StrictMode`
 * for highlighting potential problems in development.
 */

// Create a root for the React app and render it
ReactDOM.createRoot(document.getElementById("root")).render(
    <React.StrictMode>
        <App /> {/* Root app component */}
    </React.StrictMode>
);
