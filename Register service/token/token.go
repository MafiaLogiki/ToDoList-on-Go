package token

import (
    "fmt"

    "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("todolist")

func CreateToken (id int) (string, error) {
    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user-id": id,
        "iss": "todo-app",
    })

    tokenString, err := claims.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface {}, error) {return secretKey, nil})

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return token, nil
}

