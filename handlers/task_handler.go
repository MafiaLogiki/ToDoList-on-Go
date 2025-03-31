package TaskHandler

import (
    "net/http"
	"github.com/go-chi/chi/v5"
    "database/sql"
    "strconv"
    "encoding/json"

    "ToDoList/database"
    "ToDoList/models/login"
)

func getAllTasksHandler (database *sql.DB) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        tasks, err := db.GetAllTasks(database)
        if err != nil {
            json.NewEncoder(w).Encode(map[string]int {
                "Error:": http.StatusNotFound,
            })
            return
        }
        json.NewEncoder(w).Encode(tasks) 
    }
}

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

func postLoginHandler (database *sql.DB) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        var req login.Login
        err := json.NewDecoder(r.Body).Decode(&req)    
   
        if err != nil {
            json.NewEncoder(w).Encode(map[string]int{"Error": http.StatusBadRequest})
            return
        }

        username := req.Username
        password := req.Password

        token := db.GetUserToken(database, username, password)
        
        if token == "" {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]int {
                "Error in GetUserToken": http.StatusBadRequest,
            })
            return
        }
        
        http.SetCookie(w, &http.Cookie{
            Name: "token",
            Value: token,
            Path: "/",
        })

        json.NewEncoder(w).Encode(map[string]string {"token": token})
    }
}

func loginHandler (w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/login.html")
}

func taskHandler (w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/tasks.html")
}

func CreateAndRunServer (database *sql.DB, address string) error {
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
