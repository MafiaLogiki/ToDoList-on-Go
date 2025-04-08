package handlers

import (
    "net/http"
    "github.com/MafiaLogiki/common/logger"
    "message-service/internal/config"
    "github.com/go-chi/chi/v5"
    "github.com/IBM/sarama"
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

func (h *handler) newUserHandler(w http.ResponseWriter, r *http.Request) {
    
}

func (h *handler) loginHandler(w http.ResponseWriter, r *http.Request) {
    
}

func (h *handler) newTaskHandler(w http.ResponseWriter, r *http.Request) {
    
}

func (h *handler) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
    
}


func CreateAndRunServer(cfg *config.Config, producer sarama.AsyncProducer) {
    router := chi.NewRouter()


}
