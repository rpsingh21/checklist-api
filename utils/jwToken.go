package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWToken object of service
type JWToken struct {
	expirationDuration time.Duration
	secretKey          []byte
}

// Claims Custome object
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewJWToken Create new service object
func NewJWToken(expirationDuration time.Duration, secretKey string) *JWToken {
	return &JWToken{
		expirationDuration: expirationDuration,
		secretKey:          []byte(secretKey),
	}
}

// CreateToken genrate new toekn for spacfic user
func (jw *JWToken) CreateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jw.secretKey)
}

// func check(token)  {

// }
