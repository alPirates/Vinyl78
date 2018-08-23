package routes

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// SignedString - signed string
const (
	SignedString = "secret"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtUserClaims struct {
	jwt.StandardClaims
	UUID string `json:"UUID"`
	Role uint   `json:"role"`
}
