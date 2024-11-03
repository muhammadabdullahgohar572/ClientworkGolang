package handler

import (
    "net/http"
)

// CORS middleware to allow all origins and methods
func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Allow all origins; replace "*" with a specific origin if needed
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight OPTIONS request
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}


func setupRouter() *mux.Router {
    r := mux.NewRouter()
    r.Use(CORS) // Apply CORS middleware globally

    // Define your API routes
    r.HandleFunc("/test", Createuserdata).Methods("POST", "OPTIONS")

    r.HandleFunc("/test1", sign).Methods("POST", "OPTIONS")
    r.HandleFunc("/test2", login).Methods("POST", "OPTIONS")
    r.HandleFunc("/Decode", Decode).Methods("POST", "OPTIONS")

    return r
}
