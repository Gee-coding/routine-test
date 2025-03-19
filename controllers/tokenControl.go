package controllers 

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

func CreateToken(username string, password string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["name"] = username
    claims["admin"] = true
    claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}