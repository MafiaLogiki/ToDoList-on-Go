package main

import (
    "auth-service/internal/repository"
    "auth-service/internal/handlers"
    "auth-service/internal/config"

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
    cfg := config.GetConfig(logger)

    err := db.ConnectToDatabase(cfg)

    if err != nil {
        logger.Fatal("Error in database connection")
        return
    }
    defer db.CloseConnection()
    logger.Info("Database connetcion is ok")

    err2 := handlers.StartServer(cfg, logger)
    
    if err2 != nil {
        logger.Fatal("Error in server starting")
    }    
}
