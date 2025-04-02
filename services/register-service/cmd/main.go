package main

import (
    "fmt"
    "register-service/internal/repository"
    "register-service/internal/handlers"

    "github.com/MafiaLogiki/common/logger"
)


func main() {
    l := logger.NewLogger()
    err := db.ConnectToDatabase("localhost", "5432", "postgres", "1234", "users")

    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    defer db.CloseConnection()

    handlers.CreateAndRunServer(":8080", l)
}
