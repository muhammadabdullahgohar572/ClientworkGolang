package main

import (
    "net/http"
)

func main() {
    Dbconnect() // Initialize the database connection
    router := setupRouter() // Initialize the mux router
    http.Handle("/", router) // Attach router to root
    http.ListenAndServe(":8080", nil) // Start server on port 8080
}
