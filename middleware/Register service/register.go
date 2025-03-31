package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"

    "register-service/database"
)


func main() {
    err := db.ConnectToDatabase("localhost", "5432", "postgres", "1234", "users")

    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    defer db.CloseConnection()

    router := chi.NewRouter()

    httpServer := http.Server {
        Addr: "8081",
        Handler: router,
    }
    
    router.HandleFunc("/register", 
    httpServer.ListenAndServe()
}
