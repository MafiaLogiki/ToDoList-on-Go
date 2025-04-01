package main

import (
    "net/http"
    "database/sql"
    "strconv"
    "encoding/json"

    "task-service/database"

	"github.com/go-chi/chi/v5"
    "github.com/MafiaLogiki/common/auth"
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


func getTaskByIdHandler (database *sql.DB) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        taskId, _ := strconv.Atoi(chi.URLParam(r, "id")) 
        task, err := db.GetTaskById(database, taskId)
        if err != nil {
            json.NewEncoder(w).Encode(map[string]int {
                "Error:" : http.StatusInternalServerError,
            })
            return
        }
        json.NewEncoder(w).Encode(task)
    }
}

func CreateAndRunServer (address string) error {
    
    router := chi.NewRouter()
    
    router.Route("/api/tasks", func (r chi.Router) {
        r.With(auth.AuthenticateMiddleware).Get("/", getAllTasksHandler(database))
        r.With(auth.AuthenticateMiddleware).With(validateID).Get("/{id}", getTaskByIdHandler(database))
    })

    router.Route("/api/login", func (r chi.Router) {
        r.Post("/", postLoginHandler(database))
    })

    router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    router.With(auth.IsAlreadyAuth).HandleFunc("/login", http.HandlerFunc(loginHandler))
    router.With(auth.AuthenticateMiddleware).HandleFunc("/tasks", http.HandlerFunc(taskHandler))
    
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

}
