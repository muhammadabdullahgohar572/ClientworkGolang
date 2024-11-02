package main

import (
    "log"
    "net/http"
)

func main() {
    // Local server startup for development
    Dbconnect()
    router := setupRouter() // Use the router setup function from `route.go`

    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
    // Initialize the database connection and router
    Dbconnect()
    router := setupRouter()

    // Serve the request using the router
    router.ServeHTTP(w, r)
}
