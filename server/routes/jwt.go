package handler

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
	ID   uint `json:"ID"`
	Role uint `json:"role"`
}
