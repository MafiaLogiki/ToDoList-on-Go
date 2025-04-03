package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"

    "auth-service/internal/repository"
    "auth-service/internal/config"
    
    "github.com/go-chi/chi/v5"
    
    "github.com/MafiaLogiki/common/domain"
    "github.com/MafiaLogiki/common/logger"
    "github.com/MafiaLogiki/common/middleware"
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
    r.With(middleware.AuthenticateMiddleware(h.l)).HandleFunc("/*", func (w http.ResponseWriter, r *http.Request){})
}

func (h *handler) PostLoginHandler (w http.ResponseWriter, r *http.Request) () {
        var req domain.User
        err := json.NewDecoder(r.Body).Decode(&req) 
   
        if err != nil {
            h.l.Error("Can't decode json file")
            json.NewEncoder(w).Encode(map[string]string{
                "Description": "Error in input json decoding", 
                "Info": fmt.Sprintf("%v", err), 
                "Error code": fmt.Sprintf("%v", http.StatusBadRequest), 
            })
            return
        }

        username := req.Username
        password := req.Password

        token := db.GetUserToken(username, password)
        
        if token == "" {
            h.l.Error("Wrong username or password") 
            http.Error(w, "Bad request", http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]int {
                "error": http.StatusBadRequest,
            })
            return
        }
        
        http.SetCookie(w, &http.Cookie{
            Name: "token",
            Value: token,
            Path: "/",
        })

        json.NewEncoder(w).Encode(map[string]string {"token": token})
}

func StartServer(cfg *config.Config, l logger.Logger) error {
    r := chi.NewRouter()
    h := NewHandler(l)

    r.Use(logger.LoggerMiddleware)
    r.Use(middleware.AuthenticateMiddleware(l))

    h.Login(r)

    server := &http.Server {
        Addr: fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port),
        Handler: r,
    }

    l.Info(fmt.Sprintf("Server is running on %s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
    return server.ListenAndServe()
}
