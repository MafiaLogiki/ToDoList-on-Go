package main

import (
	"task-service/internal/config"
	"task-service/internal/handlers"
	"task-service/internal/repository"

	"github.com/MafiaLogiki/common/logger"
)


func main() {
    l := logger.NewLogger()
    cfg := config.GetConfig(l)

    err := db.ConnectToDatabase(cfg)
    
    if err != nil {
        l.Fatal(err)
        return
    }
    defer db.CloseConnection()

    handlers.CreateAndRunServer(cfg, l)
}
