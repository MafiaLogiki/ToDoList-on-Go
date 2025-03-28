package TaskHandler

import (
    "fmt"
    "net/http"
	"github.com/go-chi/chi/v5"
    "database/sql"
    "strconv"

    "ToDoList/database"
)

func getAllTasksHandler (database *sql.DB) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        tasks, err := db.GetAllTasks(database)
        if err != nil {
            http.Error(w, "Error while extracting data", http.StatusNotFound)
        } else {
            for i := 0; i < len(tasks); i++ {
                fmt.Printf("%v %v %v\n", tasks[i].Id, tasks[i].Title, tasks[i].Description)
           }
        }
    }
}

func validateID (next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")

        if _, err := strconv.Atoi(id); err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
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
            http.Error(w, "Error in getting task by id", http.StatusNotFound)
        }
        fmt.Printf("%v %v %v\n", task.Id, task.Title, task.Description)
    }
}

func CreateAndRunServer (database *sql.DB, address string) error {
    router := chi.NewRouter()
    
    router.HandleFunc("/tasks", getAllTasksHandler(database))
    router.With(validateID).HandleFunc("/tasks/{id}", getTaskByIdHandler(database))

    httpServer := &http.Server{
        Addr: address,
        Handler: router,
    }
    

    return httpServer.ListenAndServe()
}
