package consumer

import (
	"context"
	"notification-service/internal/config"

	"github.com/IBM/sarama"
	"github.com/MafiaLogiki/common/logger"
)

type consumerGroupHandler struct {
    l logger.Logger
}


func (h consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { 
    h.l.Info("Session started")
    return nil 
}

func (h consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { 
    h.l.Info("Session ended")
    return nil
}
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for msg := range claim.Messages() {
        h.l.Info("Received: %s\n", string(msg.Value))
        sess.MarkMessage(msg, "") 
    }

    return nil
}

func StartConsuming(cfg *config.Config, l logger.Logger) {
    consumerConfig := sarama.NewConfig()

    consumerConfig.Version = sarama.V3_5_0_0
    consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

    consumer, err := sarama.NewConsumerGroup([]string{"kafka:9092"}, "notification", consumerConfig)
    if err != nil {
        l.Fatal("Cant coneect to kafka: ", err)
    }

    handler := consumerGroupHandler{l: l}
    for {
        err := consumer.Consume(context.Background(), []string{"user-logged"}, handler)
        if err != nil {
            l.Fatal("Error from consumer: ", err) 
        }
    }
}
