package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtToken = []byte("jwtToken")

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func CreateToken(name string) (string, error) {
	expTime := time.Now().Add(time.Hour * 1)

	claims := &Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(JwtToken)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenStr string) (*Claims, error) {
	jToken := func(token *jwt.Token) (interface{}, error) {
		return JwtToken, nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, jToken)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}
	
	claims := token.Claims.(*Claims)

	return claims, nil
}