package main

import (
	"register-service/internal/config"
	"register-service/internal/handlers"
	"register-service/internal/repository"

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
