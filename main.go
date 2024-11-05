package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    Dbconnect()                  // Initialize the database connection
    router := setupRouter()      // Initialize the mux router from route.go
    http.Handle("/", router)     // Attach router to root

    log.Println("Starting server on port 8080...")
    err := http.ListenAndServe(":8080", nil) // Start server on port 8080
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

// Optional default handler function
func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Vercel!")
}
