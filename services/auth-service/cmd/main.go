package main

import (
    "fmt"
	"net/http"
    "encoding/json"

    "auth-service/internals/repository"

    "github.com/go-chi/chi/v5"
    "github.com/MafiaLogiki/common/domain"
    "github.com/MafiaLogiki/common/auth"
    "github.com/MafiaLogiki/common/logger"
)

var secretKey = []byte("todolist")

/*
Users database contains only one table with name 'user_info'
It contains three columns:
id SERIAL - PRIMARY KEY
username VARCHAR(30) - UNIQUE, NOT NULL
password_hash VARCHAR(256) - NOT NULL
*/


func postLoginHandler (w http.ResponseWriter, r *http.Request) () {
        var req domain.User
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
    router.Use(logger.LoggerMiddleware)
    router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    router.With(auth.IsAlreadyAuth).HandleFunc("/login", http.HandlerFunc(loginHandlerFunction))

    router.Route("/api/login", func (r chi.Router) {
        r.Post("/", postLoginHandler)
    })

    httpServer := &http.Server{
        Addr: ":8080",
        Handler: router,
    }

    httpServer.ListenAndServe()
}
