package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"task-service/internal/config"
	"task-service/internal/repository"
	_ "task-service/internal/validators"

	"github.com/MafiaLogiki/common/auth"
	"github.com/MafiaLogiki/common/domain"
	"github.com/MafiaLogiki/common/logger"
	"github.com/MafiaLogiki/common/middleware"

	"github.com/go-chi/chi/v5"
)


func GetAllTasksForUserHandler(w http.ResponseWriter, r *http.Request) () {
    var tasks []domain.Task

    tokenString, _ := r.Cookie("token") // No error handling because authmiddleware
    fmt.Print("AAAAAAAAAAAAAAAAAAAAA\n\n\n\n");

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

func CreateAndRunServer (cfg *config.Config, l logger.Logger) error {
    
    router := chi.NewRouter()

    router.Use(logger.LoggerMiddleware)
    router.Use(middleware.AuthenticateMiddleware(l))

    // router.Route("/api/tasks", func (r chi.Router) {
    //    r.Get("/", GetAllTasksForUserHandler)
    //    r.With(validators.ValidateID).Get("/{id}", GetTaskByIdHandler)
    // })

    router.Get("/api/tasks", GetAllTasksForUserHandler);
     
    //router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    //    http.Redirect(w, r, "/login", http.StatusSeeOther)
    //})

    httpServer := &http.Server{
        Addr: fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port),
        Handler: router,
    }
    
    return httpServer.ListenAndServe()
}
