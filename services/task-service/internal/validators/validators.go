package validators

import (
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/go-chi/chi/v5"
)

func ValidateID (next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")

        if _, err := strconv.Atoi(id); err != nil {
            json.NewEncoder(w).Encode(map[string]int {
                "Invalid ID": http.StatusBadRequest,
            })
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
