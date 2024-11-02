package main

import (
    "net/http"
    
)

// Your main function
func main() {
    Dbconnect()
    route()
}

// Vercel requires an exported HTTP handler function
func Handler(w http.ResponseWriter, r *http.Request) {
    // Initialize your database and routes here as needed
    Dbconnect()
    route()
    
    // Simple response to confirm the handler is working
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from Vercel!"))
}
