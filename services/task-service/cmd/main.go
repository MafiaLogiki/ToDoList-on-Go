package main

import (
    "task-service/internal/handlers"
    "task-service/internal/repository"

    "github.com/MafiaLogiki/common/logger"
)


func main() {
    logger := logger.NewLogger()

    err := db.ConnectToDatabase()
    if err != nil {
        return
    }
    defer db.CloseConnection()

    handlers.CreateAndRunServer(":8082", logger)
}
