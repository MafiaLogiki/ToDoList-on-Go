package handlers

import (
	"encoding/json"
    "encoding/gob"
    "bytes"
	"fmt"
	"net/http"

	"auth-service/internal/config"
	"auth-service/internal/repository"

	"github.com/go-chi/chi/v5"

	"github.com/MafiaLogiki/common/auth"
	"github.com/MafiaLogiki/common/domain"
	"github.com/MafiaLogiki/common/logger"
)

type handler struct {
    l logger.Logger
}

func NewHandler(logger logger.Logger) *handler {
    return &handler {
        l: logger,
    }
}

func (h *handler) Login(r *chi.Mux) {
    r.Post("/api/login", h.PostLoginHandler)
}

func (h *handler) PostLoginHandler (w http.ResponseWriter, r *http.Request) () {
        var req domain.User
        err := json.NewDecoder(r.Body).Decode(&req) 
   
        if err != nil {
            h.l.Error("Can't decode json file", err)
            http.Error(w, `{"status": "error", "message": "Invalid JSON"}`, http.StatusBadRequest)
            return
        }

        
        id := db.GetIdByUserData(req) 
        if id == 0 {
            h.l.Error("Wrong username or password") 
            http.Error(w, `{"status": "error", "message: "User doesn't exist"}`, http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]int {
                "error": http.StatusBadRequest,
            })
            return
        }
        
        var buffer bytes.Buffer
        err = gob.NewEncoder(&buffer).Encode(req)

        if err != nil {
            h.l.Error("Can't encode user.Domain into bytes")
            return
        }


        auth.CreateAndAddTokenToCookie(h.l, w, id)
}

func StartServer(cfg *config.Config, l logger.Logger) error {
    r := chi.NewRouter()
    h := NewHandler(l)

    r.Use(logger.LoggerMiddleware)
    // r.Use(middleware.AuthenticateMiddleware(l))

    h.Login(r)

    server := &http.Server {
        Addr: fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port),
        Handler: r,
    }

    l.Info(fmt.Sprintf("Server is running on %s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
    return server.ListenAndServe()
}
