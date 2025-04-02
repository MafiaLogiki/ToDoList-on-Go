package main

import (
    "net/http"

    "task-service/internal/handlers"
    "task-service/internal/repository"
    "task-service/internal/validators"

	"github.com/go-chi/chi/v5"
    "github.com/MafiaLogiki/common/auth"
)


func CreateAndRunServer (address string) error {
    
    router := chi.NewRouter()
    router.Use(auth.AuthenticateMiddleware)

    router.Route("/api/tasks", func (r chi.Router) {
        r.Get("/", handlers.GetAllTasksForUserHandler)
        r.With(validators.ValidateID).Get("/{id}", handlers.GetTaskByIdHandler)
    })
    
    // router.With(auth.AuthenticateMiddleware).HandleFunc("/tasks", http.HandlerFunc(taskHandler))
    
    router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    })

    httpServer := &http.Server{
        Addr: address,
        Handler: router,
    }
    
    return httpServer.ListenAndServe()
}

func main() {
    // logger := logger.NewLogger()

    err := db.ConnectToDatabase()
    if err != nil {
        return
    }
    defer db.CloseConnection()

    CreateAndRunServer(":8082")
}
