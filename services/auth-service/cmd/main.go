package main

import (
    "fmt"
	"net/http"
    "encoding/json"

    "login-service/database"
    "login-service/models/user"

	"github.com/golang-jwt/jwt/v5"
    "github.com/go-chi/chi/v5"
)

var secretKey = []byte("todolist")

/*
Users database contains only one table with name 'user_info'
It contains three columns:
id SERIAL - PRIMARY KEY
username VARCHAR(30) - UNIQUE, NOT NULL
password_hash VARCHAR(256) - NOT NULL
*/

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

func authenticateMiddleware(next http.Handler) http.Handler {
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

func isAlreadyAuth (next http.Handler) http.Handler {
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


func postLoginHandler (w http.ResponseWriter, r *http.Request) () {
        var req user.User
        err := json.NewDecoder(r.Body).Decode(&req)    
   
        if err != nil {
            json.NewEncoder(w).Encode(map[string]string{
                "Description": "Error in input json decoding", 
                "Info": fmt.Sprintf("%v", err), 
                "Error code": fmt.Sprintf("%v", http.StatusBadRequest), 
            })
            return
        }

        username := req.Username
        password := req.Password

        token := db.GetUserToken(username, password)
        
        if token == "" {
            http.Error(w, "Bad request", http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]int {
                "error": http.StatusBadRequest,
            })
            return
        }
        
        http.SetCookie(w, &http.Cookie{
            Name: "token",
            Value: token,
            Path: "/",
        })

        json.NewEncoder(w).Encode(map[string]string {"token": token})
}

func loginHandlerFunction (w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/login.html")
}

func main() {
    err := db.ConnectToDatabase("localhost", "5432", "postgres", "1234", "users")

    if err != nil {
        fmt.Printf("%v", err)
        return
    }
    defer db.CloseConnection()
    
    router := chi.NewRouter()

    router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    router.With(isAlreadyAuth).HandleFunc("/login", http.HandlerFunc(loginHandlerFunction))

    router.Route("/api/login", func (r chi.Router) {
        r.Post("/", postLoginHandler)
    })

    httpServer := &http.Server{
        Addr: ":8080",
        Handler: router,
    }

    httpServer.ListenAndServe()
}
