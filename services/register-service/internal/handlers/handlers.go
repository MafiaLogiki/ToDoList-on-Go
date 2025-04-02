package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MafiaLogiki/common/auth"
	"github.com/MafiaLogiki/common/domain"
	"github.com/MafiaLogiki/common/logger"

	"register-service/internal/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
    l logger.Logger
}

func CreateHandler(l logger.Logger) handler {
    return handler {
        l: l,
    }
}

func (h *handler) postRegisterHandler(w http.ResponseWriter, r *http.Request) {
    var newUser domain.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    
    if err != nil {
        h.l.Error("Error in json decoding")
        http.Error(w, "Error", http.StatusBadRequest)
        return
    }
     
    id, err2 := db.CreateNewUser(newUser)
    
    if err2 != nil {
        h.l.Error("Error in creating new user ", err)
        http.Error(w, "Error in creating new user", http.StatusBadRequest)
        return
    }
    auth.CreateAndAddTokenToCookie(h.l, w, id)
}

func CreateAndRunServer(addr string, l logger.Logger) {
    
    router := chi.NewRouter()
    
    h := CreateHandler(l)
    router.HandleFunc("/api/register", h.postRegisterHandler)
    
    httpServer := &http.Server {
        Addr: ":8081",
        Handler: router,
    }
    
    httpServer.ListenAndServe()
}
