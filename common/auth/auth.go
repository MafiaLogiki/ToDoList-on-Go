package auth

import (
    "fmt"
    "net/http"

    "github.com/golang-jwt/jwt/v5"
    "github.com/MafiaLogiki/common/logger"
)

var secretKey = []byte("todolist")

func VerifyToken(l logger.Logger, tokenString string) (*jwt.Token, error) {
    l.Info("Starting verifying token")
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface {}, error) {return secretKey, nil})


    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    l.Info("Token verifying done success")
    return token, nil
}

func GetIdFromToken(token string) (int, error) {
    parsedToken, _ := jwt.Parse(token, func (token *jwt.Token) (interface {}, error) {return secretKey, nil})
    claims, _ := parsedToken.Claims.(jwt.MapClaims)
    return int(claims["sub"].(float64)), nil
}


func CreateToken (l logger.Logger, id int) (string, error) {
    l.Info("Creating token with id ", id)
    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": id,
        "iss": "todo-app",
    })

    tokenString, err := claims.SignedString(secretKey)
    if err != nil {
        l.Error("Error in token creating: ", err)
        return "", err
    }

    l.Info("Creating token done successfully")
    return tokenString, nil
}

func CreateAndAddTokenToCookie(l logger.Logger, w http.ResponseWriter, id int) {
    token, err := CreateToken(l, id)

    if err != nil {
        http.Error(w, "Error while creating token", http.StatusBadRequest)
        return
    }

    http.SetCookie(w, &http.Cookie {
       Name: "token",
       Value: token,
       Path: "/",
   })
}
