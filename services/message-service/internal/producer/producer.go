package kafka

import (
    "github.com/IBM/sarama"
    "message-service/internal/config"
    "github.com/MafiaLogiki/common/logger"
)

func StartProducing(cfg *config.Config, l logger.Logger) sarama.AsyncProducer {
    producerConfig := sarama.NewConfig()

    producerConfig.Version = sarama.V3_5_0_0 
    producerConfig.Producer.Return.Successes = true
    producerConfig.Producer.Retry.Max = 3
    producerConfig.Producer.RequiredAcks = 1
    producerConfig.ClientID = "auth-service"

    producer, err := sarama.NewAsyncProducer([]string{"kafka:9092"}, producerConfig)
    if err != nil {
        l.Fatal("Can't connect to kafka: ", err) 
    }

    go func(logger.Logger) {
        for err := range producer.Errors() {
            l.Warn("Error in sended message: ", err)
        }
    }(l)

    return producer
}


