package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte

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
	var MatchuserData CreateUserData
	var Extinguser CreateUserData

	json.NewDecoder(r.Body).Decode(&MatchuserData)

	if err := databas.Where("email = ? ", MatchuserData.Email).First(&Extinguser).Error; err != nil {
		http.Error(w, "User Email not Found", http.StatusUnauthorized)
		return
	}

	if !CompareHashAndPassword(MatchuserData.Password, Extinguser.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	TokenExpire := time.Now().Add(24 * time.Hour)
	cleams := &CreateUserData{
		Name:     Extinguser.Name,
		Email:    Extinguser.Email,
		Password: Extinguser.Password,
		Gender:   Extinguser.Gender,
		Comapny:  Extinguser.Comapny,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: TokenExpire.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cleams)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Error signing token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

}

func sign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Extinguser CreateUserData
	var MatchuserData CreateUserData

	json.NewDecoder(r.Body).Decode(&MatchuserData)

	if err := databas.Where("email = ?", MatchuserData.Email).First(&Extinguser).Error; err != nil {

		http.Error(w, "email not found", http.StatusUnauthorized)
		return
	}

	if !CompareHashAndPassword(MatchuserData.Password, Extinguser.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	TokenExpire := time.Now().Add(24 * time.Hour)

	Cleaims := &CreateUserData{
		Name:     Extinguser.Name,
		Email:    Extinguser.Email,
		Password: Extinguser.Password,
		Gender:   Extinguser.Gender,
		Comapny:  Extinguser.Comapny,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: TokenExpire.Unix(),
		},
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Cleaims)

	tokenstring, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error signing token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenstring,
	})

}

func Decode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenStr := r.URL.Query().Get("token")
	if jwtKey == nil {
		jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	}
	cleims := &CreateUserData{}

	token, err := jwt.ParseWithClaims(tokenStr, cleims, func(*jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userData := map[string]interface{}{
		"UserName": cleims.Name,
		"Email":    cleims.Email,
		"Password": cleims.Password, // Include only if necessary
		"Age":      cleims.Comapny,
		"Gender":   cleims.Gender,
	}

	if err := json.NewEncoder(w).Encode(userData); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// http://localhost:8080/Decode?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MCwiQ3JlYXRlZEF0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJVcGRhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsIkRlbGV0ZWRBdCI6bnVsbCwibmFtZSI6ImFiZHVsbGFoOCIsImVtYWlsIjoiYWJkdWxsYWg5IiwicGFzc3dvcmQiOiIkMmEkMTQkUFF2M3BYS0s2bmtkVXo0ZDAuTURndVB0clRjejB3bVpZS1hVSlA5OGVwQ05CRGJXV0JYTjYiLCJHZW5kZXIiOiJNYWxlIiwiQ29tcGFueSI6IlRlY2ggSW5ub3ZhdG9ycyIsImV4cCI6MTczMDU0NTAwMn0.uWzVfBoeMEFt3RCtGV779sPmtlbrW2WWkMlbhAA_nZk
