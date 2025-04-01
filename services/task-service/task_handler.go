package main

import (
    "net/http"
    "strconv"
    "encoding/json"

    "task-service/internal/repository"

	"github.com/go-chi/chi/v5"
    "github.com/MafiaLogiki/common/auth"
    "github.com/MafiaLogiki/common/domain"
)

func validateID (next http.Handler) http.Handler {
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


func getTaskByIdHandler (w http.ResponseWriter, r *http.Request) () {
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

func getAllTasksForUserHandler(w http.ResponseWriter, r *http.Request) () {
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

func CreateAndRunServer (address string) error {
    
    router := chi.NewRouter()
    
    router.Route("/api/tasks", func (r chi.Router) {
        r.With(auth.AuthenticateMiddleware).Get("/", getAllTasksForUserHandler)
        r.With(auth.AuthenticateMiddleware).With(validateID).Get("/{id}", getTaskByIdHandler)
    })

    router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
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
    CreateAndRunServer(":8082")
}
