package auth

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("todolist")


func verifyToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface {}, error) {return secretKey, nil})

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return token, nil
}

func GetIdFromToken(token string) (int, error) {
    parsedToken, _ := jwt.Parse(token, func (token *jwt.Token) (interface {}, error) {return secretKey, nil})
    sub, _ := parsedToken.Claims.GetSubject()
    id, err := strconv.Atoi(sub)
    return id, err
}

//TODO: Define am i need this function here
func AuthenticateMiddleware(next http.Handler) http.Handler {
   return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
       cookie, err := r.Cookie("token")

       if err != nil {
           fmt.Printf("Error: %v\n", err)
           http.Redirect(w, r, "/login", http.StatusSeeOther)
           return
       }
        
       _, err2 := verifyToken(cookie.Value)

       if err2 != nil {
            fmt.Printf("Error: %v\n", err2)
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
       }

       next.ServeHTTP(w, r) 
   })
}

func IsAlreadyAuth (next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
       cookie, err := r.Cookie("token")

       if err != nil {
           next.ServeHTTP(w, r)
           return
       }
        
       _, err2 := verifyToken(cookie.Value)

       if err2 != nil { 
           next.ServeHTTP(w, r)
           return
       }

       http.Redirect(w, r, "/tasks", http.StatusSeeOther)
    })
}


func CreateToken (id int) (string, error) {
    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": id,
        "iss": "todo-app",
    })

    tokenString, err := claims.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
