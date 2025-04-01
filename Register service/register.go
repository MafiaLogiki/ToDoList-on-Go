package main

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/go-chi/chi/v5"

    "register-service/database"
    "register-service/models"
    "register-service/token"
)

var secretKey = []byte("todolist")

func registerHandlerFunction (w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/register.html")
}

func createAndAddTokenToCookie(w http.ResponseWriter, id int) {
    token, err := token.CreateToken(id) 

    if err != nil {
        http.Error(w, "Error in creating token", http.StatusBadRequest)
    }

    http.SetCookie(w, &http.Cookie {
       Name: "token",
       Value: token,
       Path: "/",
   })
}

func postRegisterHandler(w http.ResponseWriter, r *http.Request) {
    var newUser user.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    
    if err != nil {
        http.Error(w, "Error", http.StatusBadRequest)
        return
    }
     
    id, err2 := db.CreateNewUser(newUser)
    
    if err2 != nil {
        fmt.Printf("%v\n", err2)
        http.Error(w, "Error in creating new user", http.StatusBadRequest)
        return
    }
    createAndAddTokenToCookie(w, id)
}

func main() {
    err := db.ConnectToDatabase("localhost", "5432", "postgres", "1234", "users")

    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    defer db.CloseConnection()

    router := chi.NewRouter()
    
    router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) 
    router.HandleFunc("/register", registerHandlerFunction)

    router.HandleFunc("/api/register", postRegisterHandler)
    
    httpServer := &http.Server {
        Addr: ":8081",
        Handler: router,
    }
    
    httpServer.ListenAndServe()
}
