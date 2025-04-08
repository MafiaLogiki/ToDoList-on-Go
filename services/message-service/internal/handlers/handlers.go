package handlers

import (
	"fmt"
	"message-service/internal/config"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/MafiaLogiki/common/logger"
	"github.com/go-chi/chi/v5"
)

type handler struct {
    log logger.Logger
    prod sarama.AsyncProducer
}

func NewHandler (log logger.Logger, prod sarama.AsyncProducer) *handler {
   return &handler {
        log: log,
        prod: prod,
   }
}

func (h *handler) taskCreatedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Recieved")
}

func CreateAndRunServer(cfg *config.Config, l logger.Logger, producer sarama.AsyncProducer) error {
    router := chi.NewRouter()
    h := NewHandler(l, producer)

    router.Use(logger.LoggerMiddleware)

    router.Post("/api/message/create", h.taskCreatedHandler)

    server := &http.Server {
        Addr: fmt.Sprintf("%v:%v", cfg.Listen.BindIp, cfg.Listen.Port),
        Handler: router,
    }

    return server.ListenAndServe()
}
