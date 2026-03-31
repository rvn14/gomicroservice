package main

import (
    "fmt"
    "net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
}

func main() {
    
    server:= http.Server{
        Addr: ":8080",
        Handler: http.HandlerFunc(basicHandler),
    }

    fmt.Println("Server starting on :8080")
    err := server.ListenAndServe()
    if err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
    }
}