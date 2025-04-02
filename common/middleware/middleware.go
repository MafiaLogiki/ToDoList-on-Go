package middleware

import (
    "net/http"
    "fmt"
    "github.com/MafiaLogiki/common/auth"
    "github.com/MafiaLogiki/common/logger"
)

func AuthenticateMiddleware(l logger.Logger) func (http.Handler) http.Handler {
    return func (next http.Handler) http.Handler {
        return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
            cookie, err := r.Cookie("token")

            if err != nil {
           
                http.Redirect(w, r, "/login", http.StatusUnauthorized)
                return
            }
        
            _, err2 := auth.VerifyToken(l, cookie.Value)

            if err2 != nil {
                fmt.Printf("Error: %v\n", err2)
                http.Redirect(w, r, "/login", http.StatusUnauthorized)
                return
            }

            next.ServeHTTP(w, r) 
        })
    }
}


