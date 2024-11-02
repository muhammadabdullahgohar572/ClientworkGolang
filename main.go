package handler

import (
    "log"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    Dbconnect()  // Initialize your database connection
    router := setupRouter()  // Initialize your router
    router.ServeHTTP(w, r)  // Serve HTTP request with the router
}

func main() {
    // Local server startup for development purposes
    Dbconnect()
    router := setupRouter()
    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
