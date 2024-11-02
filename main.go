package handler

import (
    
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    Dbconnect()
    router := setupRouter()
    router.ServeHTTP(w, r)
}
