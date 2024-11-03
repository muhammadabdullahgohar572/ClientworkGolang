package handler

import (
    "net/http"
)

// CORS middleware to allow all origins and methods
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust as needed
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions { // Handle preflight requests
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
