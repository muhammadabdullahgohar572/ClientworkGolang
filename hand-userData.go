package main

import (
	"encoding/json"

	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


var jwtkey =[]byte("Abdullah")
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




func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  var MatchuserData CreateUserData;
  var Extinguser CreateUserData

  json.NewDecoder(r.Body).Decode(&MatchuserData);

  if err :=databas.Where("email = ? ",MatchuserData.Email).First(&Extinguser).Error;err !=nil{
	http.Error(w,"User Email not Found",http.StatusUnauthorized)
	return
  }

  if !CompareHashAndPassword(MatchuserData.Password,Extinguser.Password) {
	http.Error(w, "invalid password", http.StatusUnauthorized)
	return
  }
  
  TokenExpire :=time.Now().Add(24*time.Hour)
  cleams :=&CreateUserData{
	Name:     Extinguser.Name,
    Email:    Extinguser.Email,
    Password: Extinguser.Password,
    Gender:   Extinguser.Gender,
    Comapny: Extinguser.Comapny, 
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: TokenExpire.Unix(),
	},
  }

 

 

 token := jwt.NewWithClaims(jwt.SigningMethodHS256,cleams)
 tokenString,err := token.SignedString(jwtkey)
 
 if err !=nil {
	http.Error(w, "Error signing token", http.StatusInternalServerError)
   return
 }
 
 json.NewEncoder(w).Encode(map[string]string{
	 "token": tokenString,
 })

}



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