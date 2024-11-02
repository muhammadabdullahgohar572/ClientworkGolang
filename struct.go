package handler




import (
  "github.com/dgrijalva/jwt-go"
    "gorm.io/gorm"
)

type CreateUserData struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"Gender"`
	Comapny  string `json:"Company"`
	jwt.StandardClaims
}

