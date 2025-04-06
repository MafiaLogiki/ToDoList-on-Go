package middleware

import (
    "net/http"
    "github.com/MafiaLogiki/common/auth"
    "github.com/MafiaLogiki/common/logger"
)

func AuthenticateMiddleware(l logger.Logger) func (http.Handler) http.Handler {
    return func (next http.Handler) http.Handler {
        return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
            cookie, err := r.Cookie("token")
            
            l.Info("Middleware is started");
            if err != nil {
                l.Info("Middleware status: No cookie named token")
                http.Redirect(w, r, "http://nginx:8080/login", http.StatusUnauthorized)
                return
            }
        
            _, err2 := auth.VerifyToken(l, cookie.Value)

            if err2 != nil {
                l.Info("Middleware status: cant verify cookie");
                http.Redirect(w, r, "http://nginx:8080/login", http.StatusUnauthorized)
                return
            }

            next.ServeHTTP(w, r) 
        })
    }
}
