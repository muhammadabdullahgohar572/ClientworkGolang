package main

import (
    "example.com/practiceno1/handler" // Import the handler package from the module path
    "fmt"
    "net/http"
)

func main() {
    Dbconnect() // Initialize the database connection
    router := setupRouter() // Initialize the mux router
    http.Handle("/", router) // Attach router to root
    http.ListenAndServe(":8080", nil) // Start server on port 8080
}

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Vercel!")
}
