package main

import (
    "auth-service/internals/repository"
    "auth-service/internals/handlers"

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
        logger.Fatal("Error in database connection")
        return
    }
    defer db.CloseConnection()
    logger.Info("Database connetcion is ok")

    err2 := handlers.StartServer(":8080", logger)
    
    if err2 != nil {
        logger.Fatal("Error in server starting")
    }    
}
