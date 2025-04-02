package handlers

import (
    "net/http"
    "strconv"
    "encoding/json"

    "task-service/internal/repository"
    "task-service/internal/validators"
    
    "github.com/MafiaLogiki/common/domain"
    "github.com/MafiaLogiki/common/auth"
    "github.com/MafiaLogiki/common/logger"
    "github.com/go-chi/chi/v5"
)


func GetAllTasksForUserHandler(w http.ResponseWriter, r *http.Request) () {
    var tasks []domain.Task

    tokenString, _ := r.Cookie("token") // No error handling because authmiddleware

    id, err := auth.GetIdFromToken(tokenString.Value)
    if err != nil {
        http.Error(w, "Error", http.StatusBadRequest)
        return
    }

    tasks, err = db.GetAllTasksByUserId(id)
    if err != nil {
        http.Error(w, "Error", http.StatusBadRequest)
    }
    
    json.NewEncoder(w).Encode(tasks)
}

func GetTaskByIdHandler (w http.ResponseWriter, r *http.Request) () {
    taskId, _ := strconv.Atoi(chi.URLParam(r, "id")) 
    task, err := db.GetTaskById(taskId)
    if err != nil {
        json.NewEncoder(w).Encode(map[string]int {
            "Error:" : http.StatusInternalServerError,
        })
        return
    }
    json.NewEncoder(w).Encode(task)
}

func CreateAndRunServer (address string) error {
    
    router := chi.NewRouter()
    router.Use(logger.LoggerMiddleware)
    router.Use(auth.AuthenticateMiddleware)

    router.Route("/api/tasks", func (r chi.Router) {
        r.Get("/", GetAllTasksForUserHandler)
        r.With(validators.ValidateID).Get("/{id}", GetTaskByIdHandler)
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

