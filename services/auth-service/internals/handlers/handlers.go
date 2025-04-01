package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/go-chi/chi/v5"
    "github.com/MafiaLogiki/common/domain"
    "auth-service/internals/repository"
)

type handler struct {}

func (h *handler) Login(r *chi.Mux) {
    r.Post("/login", h.postLoginHandler)
}

func (h *handler) postLoginHandler (w http.ResponseWriter, r *http.Request) () {
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

func StartServer(addr string) {

}
