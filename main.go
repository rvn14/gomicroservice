package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var serverStartTime time.Time

func basicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    uptime := time.Since(serverStartTime).Truncate(time.Second).String()

    resp := struct {
        Uptime     string            `json:"uptime"`
        Status     string            `json:"status"`
        Components map[string]string `json:"components,omitempty"`
        Version    string            `json:"version,omitempty"`
    }{
        Uptime: uptime,
        Status: "ok",
        Components: map[string]string{
            // add real checks here, e.g. "db": checkDB(),
            "db": "unknown",
        },
        Version: "0.1.0",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func main() {
    serverStartTime = time.Now()
    router:= chi.NewRouter()
    router.Get("/", basicHandler)
    router.Get("/health", healthHandler)
    server:= http.Server{
        Addr: ":3000",
        Handler: router,
    }

    fmt.Println("Server starting on http://localhost:3000")
    err := server.ListenAndServe()
    if err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
    }
}