package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    _ "bytes"

	"task-service/internal/config"
	"task-service/internal/repository"

	"github.com/MafiaLogiki/common/auth"
	"github.com/MafiaLogiki/common/domain"
	"github.com/MafiaLogiki/common/logger"
	"github.com/MafiaLogiki/common/middleware"

	"github.com/go-chi/chi/v5"
)

type handler struct {
    l logger.Logger
}

func NewHandler (l logger.Logger) *handler {
    return &handler {
        l: l,
    }
}

func (h *handler) GetAllTasksForUserHandler(w http.ResponseWriter, r *http.Request) () {
    var tasks []domain.Task

    tokenString, _ := r.Cookie("token") // No error handling because authmiddleware

    id, err := auth.GetIdFromToken(tokenString.Value)
    if err != nil {
        h.l.Error("Internal error: ", err)
        http.Error(w, "Error", http.StatusBadRequest)
        return
    }

    tasks, err = db.GetAllTasksByUserId(id)
    if err != nil {
        fmt.Printf("%v", err)
        http.Error(w, "Error", http.StatusBadRequest)
        return
    }
    
    json.NewEncoder(w).Encode(tasks)
}

func (h *handler) GetTaskByIdHandler (w http.ResponseWriter, r *http.Request) () {
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

func (h *handler) CreateTaskHandler (w http.ResponseWriter, r *http.Request) {
    tokenString, _ := r.Cookie("token") // No error handling because authmiddleware

    id, err := auth.GetIdFromToken(tokenString.Value)
    
    if err != nil {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }
    
    var task domain.Task

    json.NewDecoder(r.Body).Decode(&task)
    task.UserId = id
    
    h.l.Info(fmt.Sprintf("Title: %v, Decription: %v, Status: %v, OwnerID: %v", task.Title, task.Description, task.Status, task.UserId))

    
    _, errQuery := db.AddTaskToTable(task)
    if errQuery != nil {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }
    
    // id, _ := result.LastInsertId()
    // data, _ := json.Marshal(id)

    /*
    resp, errResp := http.Post("message-service:8084", "application/json", bytes.NewBuffer(data))
    
    if errResp != nil {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }
    if resp.StatusCode != http.StatusOK {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }
    */
    json.NewEncoder(w).Encode(map[string]string {
        "status": "success",
    })
}

func CreateAndRunServer (cfg *config.Config, l logger.Logger) error {
    
    router := chi.NewRouter()
    h := NewHandler(l)

    router.Use(logger.LoggerMiddleware)
    router.Use(middleware.AuthenticateMiddleware(l))

    // router.Route("/api/tasks", func (r chi.Router) {
    //    r.Get("/", GetAllTasksForUserHandler)
    //    r.With(validators.ValidateID).Get("/{id}", GetTaskByIdHandler)
    // })

    router.Route("/api/tasks", func (r chi.Router) {
       r.Get("/", h.GetAllTasksForUserHandler)
       r.Post("/create", h.CreateTaskHandler)
    })
     
    //router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    //    http.Redirect(w, r, "/login", http.StatusSeeOther)
    //})

    httpServer := &http.Server{
        Addr: fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port),
        Handler: router,
    }
    
    return httpServer.ListenAndServe()
}
