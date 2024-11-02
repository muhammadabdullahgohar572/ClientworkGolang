package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CORS middleware to allow cross-origin requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func route() {
	r := mux.NewRouter()
	r.Use(CORS)

	// Define route with OPTIONS handler for preflight
	r.HandleFunc("/test", Createuserdata).Methods("POST", "OPTIONS")
	r.HandleFunc("/test1", sign).Methods("POST", "OPTIONS")
	r.HandleFunc("/test2", login).Methods("POST", "OPTIONS")
	r.HandleFunc("/Decode", Decode).Methods("POST", "OPTIONS")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
