package main

import (
	"fmt"
	"net/http"

)

// Handler function for file upload
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form to retrieve the file
	r.ParseMultipartForm(10 << 20) // Limit the size of file uploads to 10 MB

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file:", err)
		http.Error(w, "File upload error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Process the file (you can save it or do other things here)
	fmt.Fprintf(w, "Uploaded file: %s\n", handler.Filename)
	fmt.Fprintf(w, "File size: %d\n", handler.Size)
	fmt.Fprintf(w, "MIME header: %v\n", handler.Header)
}

// Route setup function
func main() {
  Dbconnect()
  route()
	http.HandleFunc("/upload", uploadHandler)

	
}
