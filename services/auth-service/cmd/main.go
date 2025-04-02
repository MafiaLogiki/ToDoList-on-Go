package main

import (
    "fmt"
    _ "encoding/json"

    "auth-service/internals/repository"
    "auth-service/internals/handlers"

    _ "github.com/MafiaLogiki/common/domain"
    _ "github.com/MafiaLogiki/common/auth"
    "github.com/MafiaLogiki/common/logger"
)

/*
Users database contains only one table with name 'user_info'
It contains three columns:
id SERIAL - PRIMARY KEY
username VARCHAR(30) - UNIQUE, NOT NULL
password_hash VARCHAR(256) - NOT NULL
*/



func main() { 
    logger := logger.NewLogger()
    err := db.ConnectToDatabase()

    if err != nil {
        fmt.Printf("%v", err)
        return
    }
    defer db.CloseConnection()
    
    handlers.StartServer(":8080", logger)
}
