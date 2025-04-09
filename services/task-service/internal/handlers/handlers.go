package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    "bytes"

	"task-service/internal/config"
	"task-service/internal/repository"
	"task-service/internal/validators"

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

    
    result, errQuery := db.AddTaskToTable(task)
    if errQuery != nil {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }
    
    insertedId, _ := result.LastInsertId()
    data, _ := json.Marshal(insertedId)

    req, _ := http.NewRequest(http.MethodPost, "http://message-service:8084/api/message/create", bytes.NewBuffer([]byte(data)))
    cookies := r.Cookies()

    for _, cookie := range cookies {
       req.AddCookie(cookie) 
    }

    resp, errResp := http.DefaultClient.Do(req) 
    
    if errResp != nil {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }

    if resp.StatusCode != http.StatusOK {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string {
        "status": "success",
    })
}

func (h *handler) UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
    var newStatus string

    taskId, _ := strconv.Atoi(chi.URLParam(r, "id"))


    err := json.NewDecoder(r.Body).Decode(&newStatus)

    if err != nil {
        http.Error(w, `{"status": "error"}`, http.StatusBadRequest)
        return
    }
    
    err = db.UpdateTaskStatusById(taskId, newStatus)

    if err != nil {
        http.Error(w, `{"status": "error"}`, http.StatusInternalServerError)
        return
    }

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
       r.With(validators.ValidateID).Put("/{id}/status", h.UpdateTaskStatusHandler)
       r.Post("/create", h.CreateTaskHandler)
    })

    httpServer := &http.Server{
        Addr: fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port),
        Handler: router,
    }
    
    return httpServer.ListenAndServe()
}
