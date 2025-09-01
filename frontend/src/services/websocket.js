/**
 * Creates a WebSocket connection to the telemetry server and sets up event handlers.
 *
 * @param {function(Object): void} onMessage - Callback invoked whenever a message is received from the server.
 * @returns {{ ws: WebSocket, sendSpeedDelta: function(number): void }}
 *          Returns the WebSocket instance and a helper to send speed adjustment events.
 */
export function createWebSocket(onMessage) {
    const ws = new WebSocket("ws://localhost:8080/ws");

    // WebSocket successfully opened
    ws.onopen = () => console.log("WebSocket connection established");

    // Handle incoming messages
    ws.onmessage = (event) => {
        try {
            const data = JSON.parse(event.data);
            onMessage(data);
        } catch (err) {
            console.error("Failed to parse WebSocket message:", err, event.data);
        }
    };

    // Handle errors
    ws.onerror = (err) => console.error("WebSocket error:", err);

    // Handle close events
    ws.onclose = (event) => {
        console.log(`WebSocket closed (code: ${event.code}, reason: ${event.reason})`);
    };

    /**
     * Sends a speed adjustment event to the backend.
     * @param {number} delta - Positive or negative speed change (e.g., +5, -5)
     */
    function sendSpeedDelta(delta) {
        if (ws.readyState === WebSocket.OPEN) {
            ws.send(JSON.stringify({ type: "speedDelta", value: delta }));
        } else {
            console.warn("WebSocket not open, cannot send speed delta");
        }
    }

    return { ws, sendSpeedDelta };
}
