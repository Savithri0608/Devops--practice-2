package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"
)

// Simple struct for health response
type Health struct {
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
    Host      string `json:"host,omitempty"`
}

func main() {
    mux := http.NewServeMux()

    // Root endpoint
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Hello from Go Web App — version 1.0!\n"))
    })

    // Health check endpoint
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        h := Health{
            Status:    "ok",
            Timestamp: time.Now().UTC().Format(time.RFC3339),
            Host:      getHost(),
        }
        json.NewEncoder(w).Encode(h)
    })

    port := getPort()
    log.Printf("🚀 Server running on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, mux))
}

func getPort() string {
    if p := os.Getenv("PORT"); p != "" {
        return p
    }
    return "8080"
}

func getHost() string {
    hn, err := os.Hostname()
    if err != nil {
        return "unknown"
    }
    return hn
}
