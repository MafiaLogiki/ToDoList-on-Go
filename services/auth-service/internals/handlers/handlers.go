package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"

    "auth-service/internals/repository"
    
    "github.com/go-chi/chi/v5"
    
    "github.com/MafiaLogiki/common/domain"
    "github.com/MafiaLogiki/common/logger"
    "github.com/MafiaLogiki/common/auth"
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
    r.With(auth.AuthenticateMiddleware).HandleFunc("/*", func (w http.ResponseWriter, r *http.Request){})
}

func (h *handler) PostLoginHandler (w http.ResponseWriter, r *http.Request) () {
        var req domain.User
        err := json.NewDecoder(r.Body).Decode(&req) 
   
        if err != nil {
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

func StartServer(addr string, l logger.Logger) error {
    r := chi.NewRouter()
    h := NewHandler(l)

    r.Use(logger.LoggerMiddleware)
    r.Use(auth.AuthenticateMiddleware)

    h.Login(r)

    server := &http.Server {
        Addr: addr,
        Handler: r,
    }

    return server.ListenAndServe()
}
