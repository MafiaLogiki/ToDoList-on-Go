package main

import (
    "github.com/MafiaLogiki/common/logger"
    "notification-service/internal/config"
    "notification-service/internal/consumer"
)

func main() {
    logger := logger.NewLogger()
    cfg := config.GetConfig(logger)
    consumer.StartConsuming(cfg, logger)
}
