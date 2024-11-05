
package handler
import (
    "net/http"
    "github.com/gorilla/mux" 
)

// CORS middleware to allow all origins and methods
// CORS middleware to allow all origins and handle Vercel environments
func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Allow all origins for testing; you can specify specific domains in production.
        w.Header().Set("Access-Control-Allow-Origin", "*")

        // Allow standard HTTP methods
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

        // Allow necessary headers
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

        // Support credentials if required (set to "false" if not needed)
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        // Allow preflight requests from Vercel’s environment
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Proceed to the next handler
        next.ServeHTTP(w, r)
    })
}

func setupRouter() *mux.Router {
    r := mux.NewRouter()
    r.Use(CORS) // Apply CORS middleware globally

    // Define API routes
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Welcome to the API root!"))
    }).Methods("GET")

    r.HandleFunc("/test", Createuserdata).Methods("POST", "OPTIONS")
    r.HandleFunc("/test1", sign).Methods("POST", "OPTIONS")
    r.HandleFunc("/test2", login).Methods("POST", "OPTIONS")
    r.HandleFunc("/Decode", Decode).Methods("POST", "OPTIONS")

    return r
}

