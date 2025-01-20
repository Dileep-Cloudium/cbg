package types

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}
