package main

import (
	"github.com/MafiaLogiki/common/logger"

	"message-service/internal/config"
	"message-service/internal/handlers"
	"message-service/internal/producer"
)

func main() {
    logger := logger.NewLogger()
    cfg := config.GetConfig(logger)

    producer := kafka.StartProducing(cfg, logger)
    handlers.CreateAndRunServer(cfg, producer)
}
