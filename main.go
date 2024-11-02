package handler

import (
    "log"
    "net/http"
)

// Entry point for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
    Dbconnect()               // Initialize the database connection
    router := setupRouter()    // Setup the router
    router.ServeHTTP(w, r)     // Serve the HTTP request with the router
}

// Optional main function for local development
func main() {
    Dbconnect()
    router := setupRouter()
    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
