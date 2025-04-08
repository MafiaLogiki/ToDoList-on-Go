package consumer

import (
    "github.com/IBM/sarama"
    "github.com/MafiaLogiki/common/logger"
)

type consumerGroupHandler struct {
    l logger.Logger
}



func StartConsuming(cfg *sarama.Config, l logger.Logger) error {
    
}
