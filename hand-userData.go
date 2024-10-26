package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashAndPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {

		return false // Password does not match the stored hash.

	}
	return err == nil
}

func Createuserdata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var CreateuserdataNew CreateUserData

	if err := json.NewDecoder(r.Body).Decode(&CreateuserdataNew); err != nil {
		http.Error(w, "1", http.StatusBadRequest)
		return
	}

	HashPassword, err := HashPassword(CreateuserdataNew.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	CreateuserdataNew.Password = HashPassword

	if err := databas.Create(&CreateuserdataNew).Error; err != nil {
		http.Error(w, "Databasea", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(CreateuserdataNew); err != nil {
		http.Error(w, "e", http.StatusBadRequest)
	}

}

func sign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var MatchuserData CreateUserData
	var Extinguser CreateUserData;
	json.NewDecoder(r.Body).Decode(MatchuserData);
	if err :=databas.Where("email =?",MatchuserData.Email).First(&Extinguser).Error;err !=nil {
       http.Error(w,"user Email not Found",http.StatusUnauthorized)
	   return		
	}
	
	if !CompareHashAndPassword(MatchuserData.Password,Extinguser.Password) {
		http.Error(w,"involid password",http.StatusUnauthorized)
	}
     


}

// var jwtKey = []byte("abdullah")
// func login(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var userdatac CreateUserData
// 	var dbUser CreateUserData

// 	// Decode incoming request for email and password
// 	json.NewDecoder(r.Body).Decode(&userdatac)

// 	// Find user by email in the database
// 	if err := Database.Where("email = ?", userdatac.Email).First(&dbUser).Error; err != nil {
// 		http.Error(w, "User not found", http.StatusUnauthorized)
// 		return
// 	}

// 	// Check password hash
// 	if !CheckPasswordHash(userdatac.Password, dbUser.Password) {
// 		http.Error(w, "Invalid password", http.StatusUnauthorized)
// 		return
// 	}

// 	// Create JWT token with all user data
// 	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
// 	claims := &CreateUserData{
// 		UserName: dbUser.UserName,
// 		Email:    dbUser.Email,
// 		Password: dbUser.Password, // Include only if necessary (not recommended)
// 		Age:      dbUser.Age,
// 		Gender:   dbUser.Gender,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	// Generate token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		http.Error(w, "Error generating token", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the token
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"token": tokenString,
// 	})
// }


// func decodeToken(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Get the token from the URL query parameter
// 	tokenStr := r.URL.Query().Get("token")

// 	// Parse the token and validate its signature
// 	claims := &CreateUserData{}
// 	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil // Replace jwtKey with your secret key
// 	})

// 	// Check if there was an error in decoding or if the token is invalid
// 	if err != nil || !token.Valid {
// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Return user data based on the token
// 	userData := map[string]interface{}{
// 		"UserName": claims.UserName,
// 		"Email":    claims.Email,
// 		"Password": claims.Password, // Include only if necessary
// 		"Age":      claims.Age,
// 		"Gender":   claims.Gender,
// 	}

// 	// Return the extracted user data as a JSON response
// 	json.NewEncoder(w).Encode(userData)
// }